package gateway

import (
	"context"
	"fmt"
	"net"

	"github.com/golang/glog"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/osrg/gobgp/pkg/packet/bgp"
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
		return &pbapi.SRv6L3Response{}, fmt.Errorf("request failed, BGP service is not available")
	}

	// RD, VPN Name or RTs can be used as primary selection criteria, one of them must be
	// present in the request.
	if req.Rd == nil && req.VpnName == "" && len(req.Rt) == 0 {
		return &pbapi.SRv6L3Response{}, fmt.Errorf("request failed, at leat one of RD or VPN name or RTs defined")
	}
	// Check each attribute and prepare it for NewL3VpnReq call
	var rd bgp.RouteDistinguisherInterface
	var err error
	m := "SRv6 L3 request for "
	if req.VpnName != "" {
		m += "VPN Name: " + req.VpnName
	}
	if req.Rd != nil {
		rd, err = bgpclient.UnmarshalRD(req.Rd)
		if err != nil {
			return &pbapi.SRv6L3Response{}, err
		}
		m += "RD: " + rd.String()
	}
	var rts []string
	if len(req.Rt) != 0 {
		rt, err := bgpclient.UnmarshalRT(req.Rt)
		if err != nil {
			return &pbapi.SRv6L3Response{}, err
		}
		m += "RTs: "
		for _, r := range rt {
			rts = append(rts, r.String())
			m += r.String() + " "
		}
	}
	// Check for an optional prefix
	var addr string
	var mask int
	if req.VpnPrefix != nil {
		addr = net.IP(req.VpnPrefix.Address).String()
		mask = int(req.VpnPrefix.MaskLength)
		m += fmt.Sprintf("VPN Prefix: %s/%d", addr, mask)
	}
	glog.V(5).Infof("%s", m)
	// Check if RD is not nil before getting its string representation
	rds := ""
	if rd != nil {
		rds = rd.String()
	}
	rq := dbclient.NewL3VpnReq(req.VpnName, rds, rts, req.Ipv4, addr, uint32(mask))

	rs, err := dbi.SRv6L3VpnRequest(context.TODO(), rq)
	if err != nil {
		return nil, err
	}

	return &pbapi.SRv6L3Response{
		Srv6Prefix: rs.Prefix,
	}, nil
}

func (g *gateway) AddSRv6L3Route(ctx context.Context, req *pbapi.SRv6L3Route) (*empty.Empty, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		client := md.Get("CLIENT_IP")
		if len(client) != 0 {
			glog.Infof("Add SRv6 L3 Route request from client: %+v", client)
		}
	}
	bgpi, ok := g.bgp.GetClientInterface().(bgpclient.BGPServices)
	if !ok {
		return nil, fmt.Errorf("gateway bgp interface is not initialized")
	}
	if err := bgpi.AddSRv6L3Route(ctx, req.Path); err != nil {
		return nil, err
	}
	// Request to programm SRv6 L3 route succeeded, next is to check if the request came from a monitored
	// client, if it is the case, then store programmed prefixes in the client's info to delete them after
	// the client is reported as gone.
	c := g.clientMgmt.Get(string(req.Id))
	if c == nil {
		// Request came from a non monitored client
		return &empty.Empty{}, nil
	}
	// Add callback to delete all programmed routes
	c.AddRouteCleanup(func() error {
		return bgpi.DelSRv6L3Route(context.TODO(), req.Path)
	})

	return &empty.Empty{}, nil
}

func (g *gateway) DelSRv6L3Route(ctx context.Context, req *pbapi.SRv6L3Route) (*empty.Empty, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		client := md.Get("CLIENT_IP")
		if len(client) != 0 {
			glog.Infof("Delete SRv6 L3 Route request from client: %+v", client)
		}
	}
	bgpi, ok := g.bgp.GetClientInterface().(bgpclient.BGPServices)
	if !ok {
		return nil, fmt.Errorf("gateway bgp interface is not initialized")
	}

	return &empty.Empty{}, bgpi.DelSRv6L3Route(ctx, req.Path)
}
