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

func (g *gateway) AdvBGPVPNv4(ctx context.Context, req *pbapi.VPNv4Prefix) (*empty.Empty, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		client := md.Get("CLIENT_IP")
		if len(client) != 0 {
			glog.Infof("Advertise BGP request from client: %+v", client)
		}
	}
	// Check if Database interface is available, if not then there is no reason to do any processing
	bgp, ok := g.bgp.GetClientInterface().(bgpclient.BGPServices)
	if !ok {
		return nil, fmt.Errorf("request failed, Database service is not available")
	}
	if err := bgp.AdvertiseVPNv4(req.Prefix); err != nil {
		return &empty.Empty{}, err
	}

	return &empty.Empty{}, nil
}

func (g *gateway) WdBGPVPNv4(ctx context.Context, req *pbapi.VPNv4Prefix) (*empty.Empty, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		client := md.Get("CLIENT_IP")
		if len(client) != 0 {
			glog.Infof("Withdraw BGP request from client: %+v", client)
		}
	}
	// Check if Database interface is available, if not then there is no reason to do any processing
	bgp, ok := g.bgp.GetClientInterface().(bgpclient.BGPServices)
	if !ok {
		return nil, fmt.Errorf("request failed, Database service is not available")
	}
	if err := bgp.WithdrawVPNv4(req.Prefix); err != nil {
		return &empty.Empty{}, err
	}

	return &empty.Empty{}, nil
}
