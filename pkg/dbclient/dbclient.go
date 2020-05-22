package dbclient

import (
	"context"

	"github.com/sbezverk/gobmp/pkg/srv6"
	pbapi "github.com/sbezverk/jalapeno-gateway/pkg/apis"
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
	MPLSL3VpnRequest(context.Context, *L3VpnReq) (*MPLSL3VpnResp, error)
	SRv6L3VpnRequest(context.Context, *L3VpnReq) (*SRv6L3VpnResp, error)
}

// L3VpnReq defines data struct for L3 VPN database request
type L3VpnReq struct {
	RD         string
	IPv4       bool
	RT         []string
	Prefix     string
	MaskLength uint32
}

// MPLSL3VpnPrefix defines L3 VPN prefix Database object
type MPLSL3VpnPrefix struct {
	Prefix     string
	MaskLength uint32
	VpnLabel   uint32
	RT         []string
}

// MPLSL3VpnResp defines data struct for L3 VPN database response
type MPLSL3VpnResp struct {
	Prefix []*pbapi.MPLSL3Prefix
}

// SRv6L3VpnPrefix defines SRv6 L3 VPN prefix Database object
type SRv6L3VpnPrefix struct {
	Prefix     string
	MaskLength uint32
	RT         []string
	PrefixSID  *srv6.L3Service
}

// SRv6L3VpnResp defines data struct for SRv6 L3 VPN database response
type SRv6L3VpnResp struct {
	Prefix []*pbapi.SRv6L3Prefix
}

// NewL3VpnReq instantiates a L3 VPN Databse Request object
func NewL3VpnReq(rd string, rt []string, ipv4 bool, prefix string, masklength uint32) *L3VpnReq {
	r := L3VpnReq{
		IPv4: ipv4,
	}
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
