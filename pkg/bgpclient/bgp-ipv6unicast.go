package bgpclient

import (
	"context"
	"fmt"
	"net"

	"github.com/golang/glog"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/any"
	api "github.com/osrg/gobgp/api"
	pbapi "github.com/sbezverk/jalapeno-gateway/pkg/apis"
)

func (bgp *bgpClient) AddIPv6UnicatRoute(ctx context.Context, path []*pbapi.BGPPath) error {
	for _, p := range path {
		if err := validateIPv6Prefix(p); err != nil {
			return err
		}
		if err := bgp.addIPv6UnicatRoute(p); err != nil {
			return err
		}
	}

	return nil
}

func (bgp *bgpClient) DelIPv6UnicatRoute(ctx context.Context, path []*pbapi.BGPPath) error {
	for _, p := range path {
		if err := validateIPv6Prefix(p); err != nil {
			return err
		}
		if err := bgp.delIPv6UnicatRoute(p); err != nil {
			return err
		}
	}

	return nil
}

func validateIPv6Prefix(p *pbapi.BGPPath) error {
	if p == nil {
		return fmt.Errorf("prefix is nil")
	}
	glog.Infof("ip prefix: %+v", *p)
	// Validating IPv6 address
	if net.IP(p.Prefix.Address).To16() == nil {
		return fmt.Errorf("invalid ipv6 prefix %+v", p.Prefix.Address)
	}
	// Validating Mask
	if p.Prefix.MaskLength <= 0 || p.Prefix.MaskLength > 128 {
		return fmt.Errorf("invalid mask length %d", p.Prefix.MaskLength)
	}
	if net.ParseIP(p.NextHop.NextHop).To16() == nil {
		return fmt.Errorf("invalid next hop address %s", p.NextHop.NextHop)
	}
	switch p.Origin.Origin {
	case 0:
	case 1:
	case 2:
	default:
		return fmt.Errorf("invalid prefix's origin %d", p.Origin.Origin)
	}

	return nil
}

func (bgp *bgpClient) addIPv6UnicatRoute(p *pbapi.BGPPath) error {
	if p == nil {
		return fmt.Errorf("prefix is nil")
	}
	nlri, _ := ptypes.MarshalAny(&api.IPAddressPrefix{
		Prefix:    net.IP(p.Prefix.Address).To16().String(),
		PrefixLen: p.Prefix.MaskLength,
	})

	a1, _ := ptypes.MarshalAny(&api.OriginAttribute{
		Origin: p.Origin.Origin,
	})
	a2, _ := ptypes.MarshalAny(&api.NextHopAttribute{
		NextHop: p.NextHop.NextHop,
	})

	attrs := []*any.Any{a1, a2}
	_, err := bgp.client.AddPath(context.TODO(), &api.AddPathRequest{
		TableType: api.TableType_GLOBAL,
		Path: &api.Path{
			Family: &api.Family{Afi: api.Family_AFI_IP6, Safi: api.Family_SAFI_UNICAST},
			Nlri:   nlri,
			Pattrs: attrs,
		},
	})
	if err != nil {
		return fmt.Errorf("failed to run AddPath call with error: %+v", err)
	}

	return nil
}
func (bgp *bgpClient) delIPv6UnicatRoute(p *pbapi.BGPPath) error {
	if p == nil {
		return fmt.Errorf("prefix is nil")
	}
	nlri, _ := ptypes.MarshalAny(&api.IPAddressPrefix{
		Prefix:    net.IP(p.Prefix.Address).To16().String(),
		PrefixLen: p.Prefix.MaskLength,
	})

	a1, _ := ptypes.MarshalAny(&api.OriginAttribute{
		Origin: p.Origin.Origin,
	})
	a2, _ := ptypes.MarshalAny(&api.NextHopAttribute{
		NextHop: p.NextHop.NextHop,
	})

	attrs := []*any.Any{a1, a2}
	_, err := bgp.client.DeletePath(context.TODO(), &api.DeletePathRequest{
		TableType: api.TableType_GLOBAL,
		Path: &api.Path{
			Family: &api.Family{Afi: api.Family_AFI_IP6, Safi: api.Family_SAFI_UNICAST},
			Nlri:   nlri,
			Pattrs: attrs,
		},
	})
	if err != nil {
		return fmt.Errorf("failed to run DeletePath call with error: %+v", err)
	}

	return nil
}
