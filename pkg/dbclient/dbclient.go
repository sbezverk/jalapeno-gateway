package dbclient

import (
	"context"

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
}

// NewL3VpnReq instantiates a L3 VPN Databse Request object
func NewL3VpnReq(name string, rd string, rt []string, ipv4 bool, prefix string, masklength uint32) *types.L3VpnReq {
	r := types.L3VpnReq{
		IPv4: ipv4,
		Name: name,
	}
	r.RD = rd
	r.RT = rt
	if prefix != "" {
		r.Prefix = prefix
		if masklength != 0 {
			r.MaskLength = masklength
		}
	}

	return &r
}
