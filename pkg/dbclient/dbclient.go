package dbclient

import (
	"context"
	"fmt"
	"net"
	"strings"

	"github.com/golang/protobuf/ptypes/any"
	"github.com/sbezverk/gobmp/pkg/prefixsid"
	pbapi "github.com/sbezverk/jalapeno-gateway/pkg/apis"
	"github.com/sbezverk/jalapeno-gateway/pkg/bgpclient"
	"github.com/sbezverk/jalapeno-gateway/pkg/srvclient"
	"github.com/sbezverk/jalapeno-gateway/pkg/types"
)

const (
	// RouteTargetPrefix defines a string with a prefix identifying extended community as
	// a route tareget extended community.
	RouteTargetPrefix = "rt="
)

// DBClient defines methods a particular database client must implement
type DBClient interface {
	// Embeding Server interface
	srvclient.Server
	// Additional Database specific methods
	DBServices
}

// DBServices defines interface for Database Services
type DBServices interface {
	MPLSL3VpnRequest(context.Context, *types.L3VpnReq) (*types.MPLSL3VpnResp, error)
	SRv6L3VpnRequest(context.Context, *types.L3VpnReq) (*types.SRv6L3VpnResp, error)
	VPNRTRequest(context.Context, string) (string, error)
}

// NewL3VpnReq instantiates a L3 VPN Databse Request object
func NewL3VpnReq(name string, rt string, ipv4 bool) *types.L3VpnReq {
	r := types.L3VpnReq{
		IPv4: ipv4,
		Name: name,
	}
	r.RT = rt
	return &r
}

// GetRT extracts RT from VPN structure
func GetRT(vpn *types.VRF, name string) (string, error) {
	if vpn.ConfigParameters == nil {
		return "", fmt.Errorf("vpn %s does not have cpnfiguration parameters", name)
	}
	af, ok := vpn.ConfigParameters.AddressFamilies[types.IPv4Unicast]
	if !ok {
		return "", fmt.Errorf("vpn %s does not have IPv4 Unicast address family", name)
	}
	if af == nil {
		return "", fmt.Errorf("vpn %s address family IPv4 Unicast is nil", name)
	}
	rt := af.RouteTargets[types.RouteTargetLocationCore][types.RouteTargetActionImport][types.RouteTargetTypeNative]
	if len(rt) == 0 {
		return "", fmt.Errorf("vpn %s does not have %s %s %s route target", name, types.RouteTargetLocationCore, types.RouteTargetActionImport, types.RouteTargetTypeNative)
	}

	// Returning first route target found on the list of RTs
	return rt[0], nil
}

// GetSRv6Prefixes builds a slice of SRv6L3Prefix from SRv6L3Record records
func GetSRv6Prefixes(records []*types.SRv6L3Record) []*pbapi.SRv6L3Prefix {
	result := make([]*pbapi.SRv6L3Prefix, len(records))
	i := 0
	for _, r := range records {
		p := &pbapi.SRv6L3Prefix{
			Prefix: &pbapi.Prefix{
				Address:    net.ParseIP(r.Prefix).To4(),
				MaskLength: uint32(r.PrefixLen),
			},
			Label:     int32(r.Labels[0]),
			NhAddress: net.ParseIP(r.Nexthop).To16(),
			PrefixSid: &pbapi.PrefixSID{},
		}
		p.Asn = uint32(r.OriginAS)
		p.PrefixSid.Tlvs = bgpclient.MarshalPrefixSID(&prefixsid.PSid{SRv6L3Service: r.SRv6SID})
		rts := make([]*any.Any, 0)
		for _, s := range r.ExtComm {
			if !strings.HasPrefix(s, RouteTargetPrefix) {
				continue
			}
			// Found route target extended community, removing route target prefix and marshal it.
			rt, err := bgpclient.MarshalRTFromString(strings.TrimPrefix(s, RouteTargetPrefix))
			if err != nil {
				continue
			}
			rts = append(rts, rt)
		}
		p.Rt = rts
		result[i] = p
		i++
	}

	return result
}
