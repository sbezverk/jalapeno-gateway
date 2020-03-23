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

func (bgp *bgpClient) AdvertiseVPNv4(prefix []*pbapi.VPNPrefix) error {
	if err := validateVPNv4Prefix(prefix); err != nil {
		return err
	}
	for _, p := range prefix {
		if err := bgp.advertiseVPNv4Prefix(context.TODO(), p); err != nil {
			return err
		}
	}

	return nil
}

func (bgp *bgpClient) WithdrawVPNv4(prefix []*pbapi.VPNPrefix) error {
	if err := validateVPNv4Prefix(prefix); err != nil {
		return err
	}

	return nil
}

func (bgp *bgpClient) AdvertiseLUPrefix(prefix []*pbapi.LUPrefix) error {
	if err := validateLUPrefix(prefix); err != nil {
		return err
	}
	for _, p := range prefix {
		if err := bgp.advertiseLUPrefix(context.TODO(), p); err != nil {
			return err
		}
	}

	return nil
}

func (bgp *bgpClient) WithdrawLUPrefix(prefix []*pbapi.LUPrefix) error {
	if err := validateLUPrefix(prefix); err != nil {
		return err
	}

	return nil
}

func validateVPNv4Prefix(prefix []*pbapi.VPNPrefix) error {
	for _, p := range prefix {
		if p == nil {
			continue
		}
		glog.Infof("vpnv4 prefix: %+v", *p)
		// Validating IP address
		if net.IP(p.Prefix.Address).To4() == nil {
			return fmt.Errorf("invalid ipv4 address %+v", p.Prefix.Address)
		}
		// Validating Mask
		if p.Prefix.MaskLength <= 0 || p.Prefix.MaskLength > 32 {
			return fmt.Errorf("invalid mask length %d", p.Prefix.MaskLength)
		}
		// Validating vpn Label that it is not excedding 2^20
		if p.VpnLabel <= 0 || p.VpnLabel > 1048576 {
			return fmt.Errorf("invalid vpn label %d", p.VpnLabel)
		}
		if net.IP(p.NhAddress).To4() == nil {
			return fmt.Errorf("invalid next hop address %+v", p.NhAddress)
		}
	}

	return nil
}

func validateLUPrefix(prefix []*pbapi.LUPrefix) error {
	for _, p := range prefix {
		if p == nil {
			continue
		}
		glog.Infof("vpnv4 prefix: %+v", *p)
		// Validating IP address
		if net.IP(p.Prefix.Address).To4() == nil {
			return fmt.Errorf("invalid ipv4 address %+v", p.Prefix.Address)
		}
		// Validating Mask
		if p.Prefix.MaskLength <= 0 || p.Prefix.MaskLength > 32 {
			return fmt.Errorf("invalid mask length %d", p.Prefix.MaskLength)
		}
		// Validating vpn Label that it is not excedding 2^20
		if p.UcastLabel <= 0 || p.UcastLabel > 1048576 {
			return fmt.Errorf("invalid vpn label %d", p.UcastLabel)
		}
	}

	return nil
}

func (bgp *bgpClient) advertiseVPNv4Prefix(ctx context.Context, prefix *pbapi.VPNPrefix) error {
	nlrivpn, _ := ptypes.MarshalAny(&api.LabeledVPNIPAddressPrefix{
		Labels:    []uint32{prefix.VpnLabel},
		Rd:        prefix.Rd,
		PrefixLen: prefix.Prefix.MaskLength,
		Prefix:    net.IP(prefix.Prefix.Address).To4().String(),
	})

	a1, _ := ptypes.MarshalAny(&api.OriginAttribute{
		Origin: 0,
	})
	a2, _ := ptypes.MarshalAny(&api.NextHopAttribute{
		NextHop: net.IP(prefix.NhAddress).To4().String(),
	})
	communities := &api.ExtendedCommunitiesAttribute{}
	communities.Communities = append(communities.Communities, prefix.Rt...)
	a3, _ := ptypes.MarshalAny(communities)

	attrs := []*any.Any{a1, a2, a3}
	_, err := bgp.client.AddPath(ctx, &api.AddPathRequest{
		TableType: api.TableType_GLOBAL,
		Path: &api.Path{
			Nlri:      nlrivpn,
			Pattrs:    attrs,
			Family:    &api.Family{Afi: api.Family_AFI_IP, Safi: api.Family_SAFI_MPLS_VPN},
			Best:      true,
			SourceAsn: prefix.Asn,
		},
	})
	if err != nil {
		return err
	}

	return nil
}

func (bgp *bgpClient) advertiseLUPrefix(ctx context.Context, prefix *pbapi.LUPrefix) error {
	nlrilu, _ := ptypes.MarshalAny(&api.LabeledIPAddressPrefix{
		Labels:    []uint32{prefix.UcastLabel},
		PrefixLen: prefix.Prefix.MaskLength,
		Prefix:    net.IP(prefix.Prefix.Address).To4().String(),
	})
	a1, _ := ptypes.MarshalAny(&api.OriginAttribute{
		Origin: 0,
	})
	a2, _ := ptypes.MarshalAny(&api.NextHopAttribute{
		NextHop: net.IP(prefix.Prefix.Address).To4().String(),
	})
	attrs := []*any.Any{a1, a2}
	_, err := bgp.client.AddPath(context.TODO(), &api.AddPathRequest{
		TableType: api.TableType_GLOBAL,
		Path: &api.Path{
			Nlri:   nlrilu,
			Pattrs: attrs,
			Family: &api.Family{Afi: api.Family_AFI_IP, Safi: api.Family_SAFI_MPLS_LABEL},
			Best:   true,
		},
	})
	if err != nil {
		return err
	}

	return nil
}
