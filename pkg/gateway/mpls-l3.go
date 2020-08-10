package gateway

import (
	"context"
	"fmt"

	"github.com/golang/glog"
	"github.com/golang/protobuf/ptypes/any"
	"github.com/osrg/gobgp/pkg/packet/bgp"
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
	// Check for optional RTs

	var rt []bgp.ExtendedCommunityInterface
	var err error
	if req.Rt != nil {
		rt, err = bgpclient.UnmarshalRT([]*any.Any{req.Rt})
		if err != nil {
			return &pbapi.MPLSL3Response{}, err
		}
	}
	rq := dbclient.NewL3VpnReq("", rt[0].String(), req.Ipv4)

	rs, err := dbi.MPLSL3VpnRequest(context.TODO(), rq)
	if err != nil {
		return nil, err
	}

	return &pbapi.MPLSL3Response{
		MplsPrefix: rs.Prefix}, nil
}
