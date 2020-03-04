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
	"github.com/sbezverk/jalapeno-gateway/pkg/gateway"
	"github.com/sbezverk/jalapeno-gateway/pkg/srvclient"
)

const (
	// DefaultGatewayPort defines default port Gateway's gRPC server listens on
	// this port is a container port, not the port used for Gateway Kubernetes Service.
	defaultGatewayPort = "15151"
)

func main() {
	flag.Parse()
	flag.Set("logtostderr", "true")

	// Getting port for gRPC server to listen on, from environment varialbe
	// GATEWAY_PORT
	strPort := os.Getenv("GATEWAY_PORT")
	if strPort == "" {
		// TODO, should it fallback to the default port?
		strPort = defaultGatewayPort
		glog.Warningf("env variable \"GATEWAY_PORT\" is not defined, using default Gateway port: %s", strPort)
	}
	srvPort, err := strconv.Atoi(strPort)
	if err != nil {
		glog.Warningf("env variable \"GATEWAY_PORT\" containes an invalid value %s, using default Gateway port instead: %s", strPort, defaultGatewayPort)
		srvPort, _ = strconv.Atoi(defaultGatewayPort)
	}
	// The value of port cannot be more than max uint16
	if srvPort == 0 || srvPort > math.MaxUint16 {
		glog.Warningf("env variable \"GATEWAY_PORT\" containes an invalid value %d, using default Gateway port instead: %s\n", srvPort, defaultGatewayPort)
		srvPort, _ = strconv.Atoi(defaultGatewayPort)
	}

	// Initialize gRPC server
	conn, err := net.Listen("tcp", ":"+strPort)
	if err != nil {
		glog.Errorf("failed to setup listener with with error: %+v", err)
		os.Exit(1)
	}
	dbc, err := makeDBClient()
	if err != nil {
		glog.Errorf("failed to make db client with with error: %+v", err)
		os.Exit(1)
	}
	bgp, err := makeBGPClient()
	if err != nil {
		glog.Errorf("failed to make bgp client with with error: %+v", err)
		os.Exit(1)
	}
	gSrv := gateway.NewGateway(conn, dbc, bgp)
	gSrv.Start()

	// For now just get stuck on stop channel, later add signal processing
	stopCh := setupSignalHandler()
	<-stopCh
	gSrv.Stop()
}

func makeDBClient() (srvclient.SrvClient, error) {
	addr := "http://10.200.99.3:30852"
	db, err := srvclient.NewSrvClient(addr, arango.NewArangoDBClient("root", "jalapeno", "jalapeno", "L3VPN_Prefixes"))
	if err != nil {
		return nil, fmt.Errorf("failed to instantiate new Arango client with error: %w", err)
	}
	return db, nil
}

func makeBGPClient() (srvclient.SrvClient, error) {
	addr := "192.168.80.103:5051"
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
