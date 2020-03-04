package bgpclient

import (
	"context"
	"fmt"
	"math"
	"net"
	"os"
	"os/signal"
	"strconv"
	"sync"

	api "github.com/osrg/gobgp/api"
	"google.golang.org/grpc"

	"github.com/sbezverk/jalapeno-gateway/pkg/srvclient"
)

// BGPClient defines the interface for communication with gobgpd process
type BGPClient interface {
	GetPrefix() error
	AddPrefix() error
}

type bgpSrv struct {
	sync.Mutex
	conn   *grpc.ClientConn
	client api.GobgpApiClient
}

func (bgp *bgpSrv) Connector(addr string) error {
	conn, err := grpc.DialContext(context.TODO(), addr, grpc.WithInsecure())
	if err != nil {
		return err
	}
	client := api.NewGobgpApiClient(conn)
	// Testing connection to gobgp by requesting its global config
	if _, err := client.GetBgp(context.TODO(), &api.GetBgpRequest{}); err != nil {
		return err
	}
	bgp.Lock()
	defer bgp.Unlock()
	bgp.conn = conn
	bgp.client = client

	return nil
}

func (bgp *bgpSrv) Monitor(addr string) error {
	// Testing connection to gobgp by requesting its global config
	if _, err := bgp.client.GetBgp(context.TODO(), &api.GetBgpRequest{}); err != nil {
		return err
	}

	return nil
}

func (bgp *bgpSrv) Validator(addr string) error {
	host, port, _ := net.SplitHostPort(addr)
	if host == "" || port == "" {
		return fmt.Errorf("host or port cannot be ''")
	}
	if net.ParseIP(host) == nil {
		return fmt.Errorf("fail to parse host part of address")
	}
	np, err := strconv.Atoi(port)
	if err != nil {
		return fmt.Errorf("fail to parse port with error: %w", err)
	}
	if np == 0 || np > math.MaxUint16 {
		return fmt.Errorf("the value of port is invalid")
	}
	return nil
}

// NewBGPSrv returns an instance of a new bgp server process
func NewBGPSrv() srvclient.Server {
	return &bgpSrv{}
}
func main() {
	addr := "192.168.80.103:5051"
	b, err := srvclient.NewSrvClient(addr, NewBGPSrv())
	if err != nil {
		fmt.Printf("failed to instantiate new bgp client with error: %+v\n", err)
		os.Exit(1)

	}
	b.Connect()
	sigc := make(chan os.Signal)
	signal.Notify(sigc, os.Interrupt)
	sig := <-sigc
	fmt.Printf("received %v\n", sig)
	b.Disconnect()
}
