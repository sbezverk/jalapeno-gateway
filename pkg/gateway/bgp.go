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

	return nil, bgpi.AddIPv6UnicatRoute(ctx, req.Path)
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

	return nil, bgpi.DelIPv6UnicatRoute(ctx, req.Path)
}
