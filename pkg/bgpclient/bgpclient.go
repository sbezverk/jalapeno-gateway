package bgpclient

import (
	"context"
	"fmt"
	"math"
	"net"
	"strconv"
	"sync"

	api "github.com/osrg/gobgp/api"
	pbapi "github.com/sbezverk/jalapeno-gateway/pkg/apis"
	"github.com/sbezverk/jalapeno-gateway/pkg/srvclient"
	"google.golang.org/grpc"
)

// BGPClient defines the interface for communication with gobgpd process
type BGPClient interface {
	// Embeding Server interface
	srvclient.Server
	BGPServices
}

// BGPServices defines interface with BGP services methods
type BGPServices interface {
	AdvertiseVPNv4([]*pbapi.Prefix) error
	WithdrawVPNv4([]*pbapi.Prefix) error
}

type bgpClient struct {
	sync.Mutex
	conn   *grpc.ClientConn
	client api.GobgpApiClient
}

func (bgp *bgpClient) Connector(addr string) error {
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

func (bgp *bgpClient) Monitor(addr string) error {
	// Testing connection to gobgp by requesting its global config
	if _, err := bgp.client.GetBgp(context.TODO(), &api.GetBgpRequest{}); err != nil {
		return err
	}

	return nil
}

func (bgp *bgpClient) Validator(addr string) error {
	host, port, _ := net.SplitHostPort(addr)
	if host == "" || port == "" {
		return fmt.Errorf("host or port cannot be ''")
	}
	// Try to resolve if the hostname was used in the address
	if ip, err := net.LookupIP(host); err != nil || ip == nil {
		// Check if IP address was used in address instead of a host name
		if net.ParseIP(host) == nil {
			return fmt.Errorf("fail to parse host part of address")
		}
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
	return &bgpClient{}
}
