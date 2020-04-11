package main

import (
	"flag"
	"fmt"
	"math"
	"net"
	"os"
	"os/signal"
	"strconv"

	"github.com/golang/glog"
	"github.com/sbezverk/jalapeno-gateway/pkg/bgpclient"
	arango "github.com/sbezverk/jalapeno-gateway/pkg/dbclient/arangoclient"
	mock "github.com/sbezverk/jalapeno-gateway/pkg/dbclient/dbmockclient"
	"github.com/sbezverk/jalapeno-gateway/pkg/gateway"
	"github.com/sbezverk/jalapeno-gateway/pkg/srvclient"
)

const (
	// DefaultGatewayPort defines default port Gateway's gRPC server listens on
	// this port is a container port, not the port used for Jalapeno Gateway kubernetes Service.
	defaultGatewayPort = "40040"
)

var (
	dbAddr      string
	bgpAddr     string
	gatewayPort string
	mockdata    bool
)

func init() {
	flag.StringVar(&dbAddr, "database-address", "", "{dns name}:port or X.X.X.X:port of the graph database, for example: arangodb.jalapeno:8529")
	flag.StringVar(&bgpAddr, "gobgp-address", "", "{dns name}:port or X.X.X.X:port of the gobgp daemon, for example: gobgpd:5051")
	flag.StringVar(&gatewayPort, "gateway-port", "", "internal container port used by Jalapeno Gateway gRPC server")
	flag.BoolVar(&mockdata, "mock-data", false, "when set to true, uses file testdata.json as a database source")
}

func main() {
	flag.Parse()
	flag.Set("logtostderr", "true")

	grpcPort := defaultGatewayPort
	if gatewayPort != "" {
		if srvPort, err := strconv.Atoi(gatewayPort); err == nil {
			if srvPort != 0 && srvPort < math.MaxUint16 {
				grpcPort = gatewayPort
			}
		}
	}

	// Initialize gRPC server
	conn, err := net.Listen("tcp", ":"+grpcPort)
	if err != nil {
		glog.Errorf("failed to setup listener with with error: %+v", err)
		os.Exit(1)
	}

	if dbAddr == "" {
		glog.Errorf("database address cannot be ''")
		os.Exit(1)
	}

	var dbc srvclient.SrvClient
	if !mockdata {
		dbc, err = makeDBClient(dbAddr)
		if err != nil {
			glog.Errorf("failed to make db client with with error: %+v", err)
			os.Exit(1)
		}
	} else {
		dbc, err = makeMockDBClient()
		if err != nil {
			glog.Errorf("failed to make db client with with error: %+v", err)
			os.Exit(1)
		}
	}
	// In general it is possible to run without gpbgp process, if it is not specified
	// then corresponding client process will not be started
	var bgp srvclient.SrvClient
	if bgpAddr != "" {
		bgp, err = makeBGPClient(bgpAddr)
		if err != nil {
			glog.Errorf("failed to make bgp client with with error: %+v", err)
			os.Exit(1)
		}
	}

	gSrv := gateway.NewGateway(conn, dbc, bgp)
	gSrv.Start()

	// For now just get stuck on stop channel, later add signal processing
	stopCh := setupSignalHandler()
	<-stopCh
	gSrv.Stop()
}

func makeMockDBClient() (srvclient.SrvClient, error) {
	// TODO, Authentication credentials should be passed as a parameters.
	db, err := srvclient.NewSrvClient("", mock.NewMockDBClient("/testdata.json"))
	if err != nil {
		return nil, fmt.Errorf("failed to instantiate new Mock DB client with error: %w", err)
	}

	return db, nil
}

func makeDBClient(addr string) (srvclient.SrvClient, error) {
	// TODO, Authentication credentials should be passed as a parameters.
	db, err := srvclient.NewSrvClient(addr, arango.NewArangoDBClient("root", "jalapeno", "jalapeno", "L3VPN_FIB"))
	if err != nil {
		return nil, fmt.Errorf("failed to instantiate new Arango client with error: %w", err)
	}
	glog.Infof("Connected to ArangoDB")

	return db, nil
}

func makeBGPClient(addr string) (srvclient.SrvClient, error) {
	bgp, err := srvclient.NewSrvClient(addr, bgpclient.NewBGPSrv())
	if err != nil {
		return nil, fmt.Errorf("failed to instantiate new bgp client with error: %w", err)
	}
	return bgp, nil
}

var (
	onlyOneSignalHandler = make(chan struct{})
	shutdownSignals      = []os.Signal{os.Interrupt}
)

func setupSignalHandler() (stopCh <-chan struct{}) {
	close(onlyOneSignalHandler) // panics when called twice

	stop := make(chan struct{})
	c := make(chan os.Signal, 2)
	signal.Notify(c, shutdownSignals...)
	go func() {
		<-c
		close(stop)
		<-c
		os.Exit(1) // second signal. Exit directly.
	}()

	return stop
}
