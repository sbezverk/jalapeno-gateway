package gateway

import (
	"net"
	"time"

	"github.com/golang/glog"
	pbapi "github.com/sbezverk/jalapeno-gateway/pkg/apis"
	"github.com/sbezverk/jalapeno-gateway/pkg/srvclient"
	"google.golang.org/grpc"
)

var (
	// maxRequestProcessTime defines a maximum wait time for a client request processing by DB client
	// and returning results
	maxRequestProcessTime = time.Millisecond * 2000
)

// Gateway defines interface to Gateway gRPC server
type Gateway interface {
	Start()
	Stop()
}
type gateway struct {
	gSrv *grpc.Server
	conn net.Listener
	dbc  srvclient.SrvClient // dbclient.DBClient
	bgp  srvclient.SrvClient // bgpclient.BGPClient
}

func (g *gateway) Start() {
	glog.V(3).Infof("Starting Gateway's gRPC on %s", g.conn.Addr().String())
	if g.dbc != nil {
		glog.V(5).Infof("Starting GraphDB client process for server: %s", g.dbc.Addr())
		g.dbc.(srvclient.SrvClient).Connect()
	}
	if g.bgp != nil {
		glog.V(5).Infof("Starting GoBGPD client process for server: %s", g.bgp.Addr())
		g.bgp.(srvclient.SrvClient).Connect()
	}
	go g.gSrv.Serve(g.conn)
}

func (g *gateway) Stop() {
	glog.V(3).Infof("Stopping Gateway's gRPC server...")
	// First stopping grpc server
	g.gSrv.Stop()
	// Disconnecting Database client if it exists
	if g.dbc != nil {
		glog.V(5).Infof("Stopping GraphDB client process for server: %s", g.dbc.Addr())
		g.dbc.Disconnect()
	}
	// Disconnecting BGP client if it exists
	if g.bgp != nil {
		glog.V(5).Infof("Stopping GoBGPD client process for server: %s", g.bgp.Addr())
		g.bgp.Disconnect()
	}
}

// NewGateway return an instance of Gateway interface
func NewGateway(conn net.Listener, dbc srvclient.SrvClient, bgpc srvclient.SrvClient) Gateway {
	gSrv := gateway{
		conn: conn,
		gSrv: grpc.NewServer([]grpc.ServerOption{}...),
		dbc:  dbc,
		bgp:  bgpc,
	}
	pbapi.RegisterGatewayServiceServer(gSrv.gSrv, &gSrv)

	return &gSrv
}
