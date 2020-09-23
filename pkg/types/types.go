package types

import (
	"encoding/json"

	"github.com/sbezverk/gobmp/pkg/bgp"
	"github.com/sbezverk/gobmp/pkg/prefixsid"
	"github.com/sbezverk/gobmp/pkg/srv6"
	pbapi "github.com/sbezverk/jalapeno-gateway/pkg/apis"
)

// BGPAddressFamily defines parameters of BGP address family
type BGPAddressFamily struct {
	RedistStatic          bool `json:"redist_static"`
	MaxPathEBGP           int  `json:"max_paths_ebgp,omitempty"`
	MaxPathIBGP           int  `json:"max_paths_ibgp,omitempty"`
	LabelAllocationPerVRF bool `json:"label_allocation_mode"`
	RedistConnected       bool `json:"redist_connected"`
}

// RouteTargetElement defines a single instance of a route target
type RouteTargetElement struct {
	Assignment string `json:"-"`               // `json:"assignment"`
	AS         []byte `json:"as,omitempty"`    // "as": 577,
	Index      []byte `json:"index,omitempty"` // "index": 1128
}

func (rte RouteTargetElement) String() string {
	return string(rte.AS) + ":" + string(rte.Index)
}

// Policy defines structure of vrf's related policy
type Policy struct {
	SiteID         int    `json:"side_id"`
	InterfaceGroup string `json:"interface_group,omitempty"`
	Export         string `json:"export,omitempty"`
	Import         string `json:"import,omitempty"`
}

// StaticRoute defines structure of vrf's assigned static routes
type StaticRoute struct {
	Description  string `json:"description,omitempty"`
	SiteID       int    `json:"site_id,omitempty"`
	Prefix       []byte `json:"prefix,omitempty"`
	NextHop      []byte `json:"next_hop,omitempty"`
	PrefixLength int    `json:"prefix_length"`
}

// RouteTargetType defines map of types of route targets (types are native or leaked)
type RouteTargetType map[string][]string

// RouteTargetAction defines map of actions of route targets (actions are import or export)
type RouteTargetAction map[string]RouteTargetType

// RouteTargetLocation defines map of locations of route targets (locations are core or dc)
type RouteTargetLocation map[string]RouteTargetAction

// AddressFamily defines structure of an instance of address family
type AddressFamily struct {
	Policies         *Policy             `json:"policies,omitempty"`
	StaticRoutes     []*StaticRoute      `json:"static_routes,omitempty"`
	ConfigNeed       bool                `json:"bgp_address_family_config_needed"`
	Enabled          bool                `json:"enabled"`
	AFName           string              `json:"af_name,omitempty"`
	RouteTargets     RouteTargetLocation `json:"route_targets"`
	SAFIName         string              `json:"saf_name,omitempty"`
	BGPAddressFamily *BGPAddressFamily   `json:"bgp_address_family"`
}

// BGP defines structure of vrf's bgp parameters
type BGP struct {
	DefaultInfoOriginate bool `json:"default_info_originate"`
	RDAuto               bool `json:"rd_auto"`
}

const (
	// IPv4Unicast defines a type of address family for IPv4 Unicast
	IPv4Unicast = "ipv4unicast"
)

const (
	RouteTargetLocationCore = "core"
	RouteTargetLocationDC   = "dc"
	RouteTargetActionImport = "import"
	RouteTargetActionExport = "export"
	RouteTargetTypeNative   = "native"
	RouteTargetTypeLeaked   = "leaked"
)

// ConfigParameters defines structure of vrf's configuration parameters
type ConfigParameters struct {
	BGP             *BGP                      `json:"bgp"`
	AddressFamilies map[string]*AddressFamily `json:"address_families"`
}

// VRF defines structure of vrf table
type VRF struct {
	Created          string            `json:"-"` // `json:"created,omitempty"`
	VRFName          string            `json:"vrf_name,omitempty"`
	SecurityZone     string            `json:"security_zone,omitempty"`
	ConfigParameters *ConfigParameters `json:"config_parameters,omitempty"`
	Version          *int64            `json:"-"`
	Key              string            `json:"_key"`
}

// Hit defines a structre of information received from ElasticSearch
type Hit struct {
	Score   *float64        `json:"_score"`
	Source  json.RawMessage `json:"_source"`
	Index   string          `json:"_index"`
	Type    string          `json:"_type"`
	ID      string          `json:"_id"`
	Version *int64          `json:"_version"`
}

// SRv6L3Record represents the database record structure
type SRv6L3Record struct {
	Key            string              `json:"_key,omitempty"`
	ID             string              `json:"_id,omitempty"`
	Prefix         string              `json:"prefix,omitempty"`
	PrefixLen      int32               `json:"prefix_len,omitempty"`
	IsIPv4         bool                `json:"is_ipv4"`
	OriginAS       int32               `json:"origin_as,omitempty"`
	Nexthop        string              `json:"nexthop,omitempty"`
	Labels         []uint32            `json:"labels,omitempty"`
	RD             string              `json:"vpn_rd,omitempty"`
	BaseAttributes *bgp.BaseAttributes `json:"base_attrs,omitempty"`
	PrefixSID      *prefixsid.PSid     `json:"prefix_sid,omitempty"`
}

// L3VpnReq defines data struct for L3 VPN database request
type L3VpnReq struct {
	Name string
	IPv4 bool
	RT   string
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
