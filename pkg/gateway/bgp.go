package gateway

import (
	"context"
	"fmt"

	"github.com/golang/glog"
	"github.com/golang/protobuf/ptypes/empty"
	pbapi "github.com/sbezverk/jalapeno-gateway/pkg/apis"
	"github.com/sbezverk/jalapeno-gateway/pkg/bgpclient"
	"google.golang.org/grpc/metadata"
)

func (g *gateway) AddIPv6UnicatRoute(ctx context.Context, req *pbapi.IPv6UnicastRoute) (*empty.Empty, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		client := md.Get("CLIENT_IP")
		if len(client) != 0 {
			glog.Infof("Add IPv6 Unicast Route request from client: %+v", client)
		}
	}
	bgpi, ok := g.bgp.GetClientInterface().(bgpclient.BGPServices)
	if !ok {
		return nil, fmt.Errorf("gateway bgp interface is not initialized")
	}
	if err := bgpi.AddIPv6UnicatRoute(ctx, req.Path); err != nil {
		return nil, err
	}
	// Request to programm IPv6 Unicast succeeded, next is to check if the request came from a monitored
	// client, if it is the case, then store programmed prefixes in the client's info to delete them after
	// the client is reported as gone.
	c := g.clientMgmt.Get(string(req.Id))
	if c == nil {
		// Request came from a non monitored client
		return &empty.Empty{}, nil
	}
	// Add callback to delete all programmed routes
	c.AddRouteCleanup(func() error {
		return bgpi.DelIPv6UnicatRoute(context.TODO(), req.Path)
	})

	return &empty.Empty{}, nil
}

func (g *gateway) DelIPv6UnicatRoute(ctx context.Context, req *pbapi.IPv6UnicastRoute) (*empty.Empty, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		client := md.Get("CLIENT_IP")
		if len(client) != 0 {
			glog.Infof("Delete IPv6 Unicast Route request from client: %+v", client)
		}
	}
	bgpi, ok := g.bgp.GetClientInterface().(bgpclient.BGPServices)
	if !ok {
		return nil, fmt.Errorf("gateway bgp interface is not initialized")
	}

	return &empty.Empty{}, bgpi.DelIPv6UnicatRoute(ctx, req.Path)
}
