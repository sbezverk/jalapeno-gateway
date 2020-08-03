package dbclient

import (
	"context"

	"github.com/sbezverk/gobmp/pkg/bgp"
	"github.com/sbezverk/gobmp/pkg/prefixsid"
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
	Name       string
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

// MPLSL3Record represents the database record structure
type MPLSL3Record struct {
	Key             string `json:"_key,omitempty"`
	ID              string `json:"_id,omitempty"`
	From            string `json:"_from,omitempty"`
	To              string `json:"_to,omitempty"`
	Rev             string `json:"_rev,omitempty"`
	SourceAddr      string `json:"SrcIP,omitempty"`
	DestinationAddr string `json:"DstIP,omitempty"`
	Prefix          string `json:"VPN_Prefix,omitempty"`
	Mask            uint32 `json:"VPN_Prefix_Len,omitempty"`
	RouterID        string `json:"RouterID,omitempty"`
	VPN             uint32 `json:"VPN_Label,omitempty"`
	RD              string `json:"RD"`
	IPv4            bool   `json:"IPv4,omitempty"`
	RT              string `json:"RT,omitempty"`
	Source          string `json:"Source,omitempty"`
	Destination     string `json:"Destination,omitempty"`
}

// SRv6L3Record represents the database record structure
type SRv6L3Recordv2 struct {
	Key       string          `json:"_key,omitempty"`
	ID        string          `json:"_id,omitempty"`
	From      string          `json:"_from,omitempty"`
	To        string          `json:"_to,omitempty"`
	Rev       string          `json:"_rev,omitempty"`
	Prefix    string          `json:"VPN_Prefix,omitempty"`
	PrefixLen int32           `json:"VPN_Prefix_Len,omitempty"`
	IsIPv4    bool            `json:"IPv4"`
	OriginAS  string          `json:"origin_as,omitempty"`
	Nexthop   string          `json:"SrcIP,omitempty"`
	Labels    []uint32        `json:"labels,omitempty"`
	RD        string          `json:"RD,omitempty"`
	RT        []string        `json:"RT,omitempty"`
	PrefixSID *prefixsid.PSid `json:"prefix_sid,omitempty"`
}

// {
//     "_key": "2001:1:1:f003::1_10.0.0.3:2_1.1.103.0",
//     "_id": "L3VPN_FIB/2001:1:1:f003::1_10.0.0.3:2_1.1.103.0",
//     "_from": "L3VPNode/2001:1:1:f003::1",
//     "_to": "L3VPNPrefix/1.1.103.0",
//     "_rev": "_a5-F4OS--_",
//     "SrcIP": "2001:1:1:f003::1",
//     "DstIP": "1.1.103.0",
//     "VPN_Prefix": "1.1.103.0",
//     "VPN_Prefix_Len": 24,
//     "RouterID": "2001:1:1:f003::1",
//     "PrefixSID": null,
//     "VPN_Label": 1104,
//     "RD": "10.0.0.3:2",
//     "RT": [
//       "98:98",
//       "99:99",
//       "100:100"
//     ],
//     "IPv4": true,
//     "SRv6_SID": {
//       "sub_tlvs": {
//         "1": [
//           {
//             "sid": "2001:1:1:f003::",
//             "endpoint_behavior": 19,
//             "sub_sub_tlvs": {
//               "1": [
//                 {
//                   "local_block_length": 40,
//                   "local_node_length": 24,
//                   "function_length": 16,
//                   "transposition_length": 16,
//                   "transposition_offset": 64
//                 }
//               ]
//             }
//           }
//         ]
//       }
//     }
//   },

// SRv6L3Record represents the database record structure
type SRv6L3Record struct {
	Key            string              `json:"_key,omitempty"`
	ID             string              `json:"_id,omitempty"`
	From           string              `json:"_from,omitempty"`
	To             string              `json:"_to,omitempty"`
	Rev            string              `json:"_rev,omitempty"`
	BaseAttributes *bgp.BaseAttributes `json:"base_attrs,omitempty"`
	Prefix         string              `json:"prefix,omitempty"`
	PrefixLen      int32               `json:"prefix_len,omitempty"`
	IsIPv4         bool                `json:"is_ipv4"`
	OriginAS       string              `json:"origin_as,omitempty"`
	Nexthop        string              `json:"nexthop,omitempty"`
	IsNexthopIPv4  bool                `json:"is_nexthop_ipv4"`
	Labels         []uint32            `json:"labels,omitempty"`
	VPNRD          string              `json:"vpn_rd,omitempty"`
	VPNRDType      uint16              `json:"vpn_rd_type"`
	PrefixSID      *prefixsid.PSid     `json:"prefix_sid,omitempty"`
}

// NewL3VpnReq instantiates a L3 VPN Databse Request object
func NewL3VpnReq(name string, rd string, rt []string, ipv4 bool, prefix string, masklength uint32) *L3VpnReq {
	r := L3VpnReq{
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
