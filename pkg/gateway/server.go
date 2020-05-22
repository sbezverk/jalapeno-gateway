package gateway

import (
	"fmt"
	"net"
	"time"

	"github.com/golang/glog"
	pbapi "github.com/sbezverk/jalapeno-gateway/pkg/apis"
	"github.com/sbezverk/jalapeno-gateway/pkg/gateway/client"
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
	gSrv       *grpc.Server
	conn       net.Listener
	dbc        srvclient.SrvClient // dbclient.DBClient
	bgp        srvclient.SrvClient // bgpclient.BGPClient
	clientMgmt client.Store
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

func (g *gateway) Monitor(c pbapi.GatewayService_MonitorServer) error {
	cm, err := c.Recv()
	if err != nil {
		return err
	}
	if cm == nil {
		return fmt.Errorf("client info is nil")
	}
	glog.V(5).Infof("request from client with id %s to monitor.", string(cm.Id))
	if m := g.clientMgmt.Get(string(cm.Id)); m == nil {
		glog.V(5).Infof("adding client with id %s to the store.", string(cm.Id))
		g.clientMgmt.Add(string(cm.Id))
	} else {
		// TODO add handling of such condition
		glog.Warningf("duplicate monitor request, client with id: %s already in the store", string(cm.Id))
	}
	for {
		_, err := c.Recv()
		if err != nil {
			// Error indicates that the client is no longer functional, sending command to
			// the clients manager to remove the client and exit
			glog.V(5).Infof("client with id %s is no longer alive, error: %+v, deleting it from the store.", string(cm.Id), err)
			c := g.clientMgmt.Get(string(cm.Id))
			for _, f := range c.GetRouteCleanup() {
				if err := f(); err != nil {
					glog.Errorf("route cleanup encountered error: %+v", err)
				}
			}
			g.clientMgmt.Delete(string(cm.Id))
			return err
		}
	}
}

// NewGateway return an instance of Gateway interface
func NewGateway(conn net.Listener, dbc srvclient.SrvClient, bgpc srvclient.SrvClient) Gateway {
	gSrv := gateway{
		conn:       conn,
		gSrv:       grpc.NewServer([]grpc.ServerOption{}...),
		dbc:        dbc,
		bgp:        bgpc,
		clientMgmt: client.NewClientStore(),
	}
	pbapi.RegisterGatewayServiceServer(gSrv.gSrv, &gSrv)

	return &gSrv
}
