package dbclient

import (
	"context"
	"fmt"

	"github.com/sbezverk/jalapeno-gateway/pkg/srvclient"
	"github.com/sbezverk/jalapeno-gateway/pkg/types"
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
