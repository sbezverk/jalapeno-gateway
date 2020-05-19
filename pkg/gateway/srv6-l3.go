package gateway

import (
	"context"
	"fmt"
	"net"

	"github.com/golang/glog"
	pbapi "github.com/sbezverk/jalapeno-gateway/pkg/apis"
	"github.com/sbezverk/jalapeno-gateway/pkg/bgpclient"
	"github.com/sbezverk/jalapeno-gateway/pkg/dbclient"
	"google.golang.org/grpc/metadata"
)

func (g *gateway) SRv6L3VPN(ctx context.Context, req *pbapi.L3VpnRequest) (*pbapi.SRv6L3Response, error) {
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		client := md.Get("CLIENT_IP")
		if len(client) != 0 {
			glog.Infof("SRv6 L3 request from client: %+v", client)
		}
	}
	// Check if Database interface is available, if not then there is no reason to do any processing
	dbi, ok := g.dbc.GetClientInterface().(dbclient.DBServices)
	if !ok {
		return nil, fmt.Errorf("request failed, BGP service is not available")
	}
	// Check if RD is present in the request, if not return error as RD is a mandatory parameter
	if req.Rd == nil {
		return nil, fmt.Errorf("request failed, RD is nil")
	}
	rd, err := bgpclient.UnmarshalRD(req.Rd)
	if err != nil {
		return &pbapi.SRv6L3Response{}, err
	}
	// Check for an optional prefix
	var addr string
	var mask int
	if req.VpnPrefix != nil {
		addr = net.IP(req.VpnPrefix.Address).String()
		mask = int(req.VpnPrefix.MaskLength)
		glog.V(5).Infof("L3VPN request for prefix: %s/%d", addr, mask)
	}
	glog.V(5).Infof("SRv6 L3 request for RD: %s", rd.String())
	rq := dbclient.NewL3VpnReq(rd.String(), nil, req.Ipv4, addr, uint32(mask))

	rs, err := dbi.SRv6L3VpnRequest(context.TODO(), rq)
	if err != nil {
		return nil, err
	}

	return &pbapi.SRv6L3Response{
		Srv6Prefix: rs.Prefix,
	}, nil
}
