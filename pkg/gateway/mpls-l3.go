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

func (g *gateway) MPLSL3VPN(ctx context.Context, req *pbapi.L3VpnRequest) (*pbapi.MPLSL3Response, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		client := md.Get("CLIENT_IP")
		if len(client) != 0 {
			glog.Infof("VPN request from client: %+v", client)
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
		return &pbapi.MPLSL3Response{}, err
	}
	glog.V(5).Infof("L3VPN request for RD: %s", rd.String())

	// Check for optional RTs
	var rts []string
	if req.Rt != nil {
		rt, err := bgpclient.UnmarshalRT(req.Rt)
		if err != nil {
			return &pbapi.MPLSL3Response{}, err
		}
		for _, r := range rt {
			rts = append(rts, r.String())
		}
	}
	// Check for an optional prefix
	var addr string
	var mask int
	if req.VpnPrefix != nil {
		addr = net.IP(req.VpnPrefix.Address).String()
		mask = int(req.VpnPrefix.MaskLength)
		glog.V(5).Infof("L3VPN request for prefix: %s/%d", addr, mask)
	}

	rq := dbclient.NewL3VpnReq(rd.String(), rts, req.Ipv4, addr, uint32(mask))

	rs, err := dbi.MPLSL3VpnRequest(context.TODO(), rq)
	if err != nil {
		return nil, err
	}

	// vpnPrefix := make([]*pbapi.MPLSL3Prefix, 0)
	// for _, p := range rs.Prefix {
	// 	vpnPrefix = append(vpnPrefix, &pbapi.MPLSL3Prefix{
	// 		Prefix: &pbapi.Prefix{
	// 			Address:    net.ParseIP(p.Prefix),
	// 			MaskLength: p.MaskLength,
	// 		},
	// 		VpnLabel: p.VpnLabel,
	// 	})
	// }

	return &pbapi.MPLSL3Response{
		MplsPrefix: rs.Prefix}, nil
}
