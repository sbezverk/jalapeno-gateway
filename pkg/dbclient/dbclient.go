package dbclient

import (
	"context"

	"github.com/sbezverk/jalapeno-gateway/pkg/srvclient"
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
	L3VPNRequest(context.Context, *L3VpnReq) (*L3VpnResp, error)
}

// L3VpnReq defines data struct for L3 VPN database request
type L3VpnReq struct {
	RD         string
	RT         []string
	Prefix     string
	MaskLength uint32
}

// L3VPNPrefix defines L3 VPN prefix Database object
type L3VPNPrefix struct {
	Prefix     string
	MaskLength uint32
	RT         []string
}

// L3VpnResp defines data struct for L3 VPN database response
type L3VpnResp struct {
	VpnLabel uint32
	SidLabel uint32
	Prefix   []L3VPNPrefix
}

// NewL3VpnReq instantiates a L3 VPN Databse Request object
func NewL3VpnReq(rd string, rt []string, prefix string, masklength uint32) *L3VpnReq {
	r := L3VpnReq{}
	r.RD = rd
	if len(rt) != 0 {
		r.RT = make([]string, len(rt))
		copy(r.RT, rt)
	}
	if prefix != "" {
		r.Prefix = prefix
		if masklength != 0 {
			r.MaskLength = masklength
		}
	}

	return &r
}
