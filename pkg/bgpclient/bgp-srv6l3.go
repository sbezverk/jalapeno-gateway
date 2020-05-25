package bgpclient

import (
	"context"
	"fmt"
	"net"

	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/any"
	api "github.com/osrg/gobgp/api"
	pbapi "github.com/sbezverk/jalapeno-gateway/pkg/apis"
)

func (bgp *bgpClient) AddSRv6L3Route(ctx context.Context, path []*pbapi.SRv6L3Prefix) error {
	for _, p := range path {
		if err := validateSRv6L3Route(p); err != nil {
			return err
		}
		if err := bgp.addSRv6L3Route(p); err != nil {
			return err
		}
	}

	return nil
}

func (bgp *bgpClient) DelSRv6L3Route(ctx context.Context, path []*pbapi.SRv6L3Prefix) error {
	for _, p := range path {
		if err := validateSRv6L3Route(p); err != nil {
			return err
		}
		if err := bgp.delSRv6L3Route(p); err != nil {
			return err
		}
	}

	return nil
}

func validateSRv6L3Route(p *pbapi.SRv6L3Prefix) error {
	if p == nil {
		return fmt.Errorf("prefix is nil")
	}
	// Validating IP address
	if p.Prefix == nil {
		return fmt.Errorf("prefix is nil")
	}
	if net.IP(p.Prefix.Address).To4() == nil {
		return fmt.Errorf("invalid ipv4 address %+v", p.Prefix.Address)
	}
	// Validating Mask
	if p.Prefix.MaskLength <= 0 || p.Prefix.MaskLength > 32 {
		return fmt.Errorf("invalid mask length %d", p.Prefix.MaskLength)
	}
	// Validating vpn Label that it is not excedding 2^20
	if p.Label <= 0 || p.Label > 1048576 {
		return fmt.Errorf("invalid vpn label %d", p.Label)
	}
	if net.IP(p.NhAddress).To16() == nil {
		return fmt.Errorf("invalid next hop address %+v", p.NhAddress)
	}
	if p.PrefixSid == nil {
		return fmt.Errorf("prefix sid is nil")
	}
	if len(p.PrefixSid.Tlvs) < 1 {
		return fmt.Errorf("prefix sid has 0 TLVs")
	}

	return nil
}

func (bgp *bgpClient) addSRv6L3Route(p *pbapi.SRv6L3Prefix) error {
	if p == nil {
		return fmt.Errorf("prefix is nil")
	}

	nlrivpn, _ := ptypes.MarshalAny(&api.LabeledVPNIPAddressPrefix{
		Labels:    []uint32{uint32(p.Label)},
		Rd:        p.Rd,
		PrefixLen: p.Prefix.MaskLength,
		Prefix:    net.IP(p.Prefix.Address).To4().String(),
	})
	// Origin attribute
	origin, _ := ptypes.MarshalAny(&api.OriginAttribute{
		Origin: 0,
	})
	// Next hop attribute
	nh, _ := ptypes.MarshalAny(&api.NextHopAttribute{
		NextHop: net.IP(p.NhAddress).To16().String(),
	})
	// Extended communities attribute
	rt, _ := ptypes.MarshalAny(&api.ExtendedCommunitiesAttribute{
		Communities: p.Rt,
	})
	// Inject Prefix SID attribute
	prefixSID, _ := ptypes.MarshalAny(&api.PrefixSID{
		Tlvs: p.PrefixSid.Tlvs,
	})
	attrs := []*any.Any{origin, nh, rt, prefixSID}
	if _, err := bgp.client.AddPath(context.TODO(), &api.AddPathRequest{
		TableType: api.TableType_GLOBAL,
		Path: &api.Path{
			Nlri:      nlrivpn,
			Pattrs:    attrs,
			Family:    &api.Family{Afi: api.Family_AFI_IP, Safi: api.Family_SAFI_MPLS_VPN},
			Best:      true,
			SourceAsn: p.Asn,
		},
	}); err != nil {
		return fmt.Errorf("failed to run AddPath call with error: %v", err)
	}

	return nil
}

func (bgp *bgpClient) delSRv6L3Route(p *pbapi.SRv6L3Prefix) error {
	if p == nil {
		return fmt.Errorf("prefix is nil")
	}

	nlrivpn, _ := ptypes.MarshalAny(&api.LabeledVPNIPAddressPrefix{
		Labels:    []uint32{uint32(p.Label)},
		Rd:        p.Rd,
		PrefixLen: p.Prefix.MaskLength,
		Prefix:    net.IP(p.Prefix.Address).To4().String(),
	})
	// Origin attribute
	origin, _ := ptypes.MarshalAny(&api.OriginAttribute{
		Origin: 0,
	})
	// Next hop attribute
	nh, _ := ptypes.MarshalAny(&api.NextHopAttribute{
		NextHop: net.IP(p.NhAddress).To16().String(),
	})
	// Extended communities attribute
	rt, _ := ptypes.MarshalAny(&api.ExtendedCommunitiesAttribute{
		Communities: p.Rt,
	})
	// Inject Prefix SID attribute
	prefixSID, _ := ptypes.MarshalAny(&api.PrefixSID{
		Tlvs: p.PrefixSid.Tlvs,
	})
	attrs := []*any.Any{origin, nh, rt, prefixSID}
	if _, err := bgp.client.DeletePath(context.TODO(), &api.DeletePathRequest{
		TableType: api.TableType_GLOBAL,
		Path: &api.Path{
			Nlri:      nlrivpn,
			Pattrs:    attrs,
			Family:    &api.Family{Afi: api.Family_AFI_IP, Safi: api.Family_SAFI_MPLS_VPN},
			Best:      true,
			SourceAsn: p.Asn,
		},
	}); err != nil {
		return fmt.Errorf("failed to run DelPath call with error: %v", err)
	}

	return nil
}
