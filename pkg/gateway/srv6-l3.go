package gateway

import (
	"context"

	"github.com/golang/glog"
	pbapi "github.com/sbezverk/jalapeno-gateway/pkg/apis"
	"google.golang.org/grpc/metadata"
)

func (g *gateway) SRv6L3VPN(ctx context.Context, req *pbapi.L3VpnRequest) (*pbapi.SRv6L3Response, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		client := md.Get("CLIENT_IP")
		if len(client) != 0 {
			glog.Infof("VPN request from client: %+v", client)
		}
	}
	return &pbapi.SRv6L3Response{}, nil
}
