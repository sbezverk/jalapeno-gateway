package dbmockclient

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/golang/glog"
	"github.com/sbezverk/gobmp/pkg/bgp"
	"github.com/sbezverk/gobmp/pkg/srv6"
	pbapi "github.com/sbezverk/jalapeno-gateway/pkg/apis"
	"github.com/sbezverk/jalapeno-gateway/pkg/dbclient"
)

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
type SRv6L3Record struct {
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
	PrefixSID      *srv6.L3Service     `json:"srv6_l3_service,omitempty"`
}

type mockSrv struct {
	vpnStore  map[string][]MPLSL3Record
	srv6Store map[string][]SRv6L3Record
}

func (m *mockSrv) MPLSL3VpnRequest(ctx context.Context, req *dbclient.L3VpnReq) (*dbclient.MPLSL3VpnResp, error) {
	glog.V(5).Infof("Mock DB L3 VPN Service was called with the request: %+v", req)

	// Initial lookup for requested RD, if it is not in the store, return error
	records, ok := m.vpnStore[req.RD]
	if !ok {
		return nil, fmt.Errorf("RD %s is not found", req.RD)
	}

	// Filter by IP Family
	records = filterByIPFamily(req.IPv4, records)

	// Filter by Prefix
	if req.Prefix != "" {
		records = filterByPrefix(req.Prefix, req.MaskLength, records)
	}
	// Filter by RT
	if len(req.RT) != 0 {
		records = filterByRT(req.RT, records)
	}

	if len(records) == 0 {
		// All filtered, returning error
		return nil, fmt.Errorf("no matching records to found")
	}

	vpnPrefix := make([]*pbapi.MPLSL3Prefix, 0)
	glog.Infof("number of prefixes retrieved: %d", len(records))
	for _, r := range records {
		vpnPrefix = append(vpnPrefix, &pbapi.MPLSL3Prefix{
			Prefix: &pbapi.Prefix{
				Address:    []byte(r.Prefix),
				MaskLength: r.Mask,
			},
			VpnLabel: r.VPN,
		})
	}
	resp := dbclient.MPLSL3VpnResp{
		Prefix: vpnPrefix,
	}

	return &resp, nil
}

func (m *mockSrv) SRv6L3VpnRequest(ctx context.Context, req *dbclient.L3VpnReq) (*dbclient.SRv6L3VpnResp, error) {
	glog.V(5).Infof("Mock DB SRv6 VPN Service was called with the request: %+v", req)

	// Initial lookup for requested RD, if it is not in the store, return error
	records, ok := m.srv6Store[req.RD]
	if !ok {
		return nil, fmt.Errorf("RD %s is not found", req.RD)
	}

	if len(records) == 0 {
		// All filtered, returning error
		return nil, fmt.Errorf("no matching records to found")
	}

	srv6Prefix := make([]*pbapi.SRv6L3Prefix, 0)
	glog.Infof("number of prefixes retrieved: %d", len(records))
	for _, r := range records {
		srv6Prefix = append(srv6Prefix, &pbapi.SRv6L3Prefix{
			Prefix: &pbapi.Prefix{
				Address:    []byte(r.Prefix),
				MaskLength: uint32(r.PrefixLen),
			},
			Label: int32(r.Labels[0]),
		})
	}
	resp := dbclient.SRv6L3VpnResp{
		Prefix: srv6Prefix,
	}

	return &resp, nil
}

func (m *mockSrv) Connector(addr string) error {

	return nil
}

func (m *mockSrv) Monitor(addr string) error {
	return nil
}

func (m *mockSrv) Validator(addr string) error {
	return nil
}

func filterByIPFamily(ipv4 bool, records []MPLSL3Record) []MPLSL3Record {
	result := make([]MPLSL3Record, 0)
	for _, r := range records {
		if r.IPv4 == ipv4 {
			result = append(result, r)
		}
	}

	return result
}
func filterByPrefix(prefix string, mask uint32, records []MPLSL3Record) []MPLSL3Record {
	result := make([]MPLSL3Record, 0)
	for _, r := range records {
		if r.Prefix == prefix && r.Mask == mask {
			result = append(result, r)
			break
		}

	}

	return result
}

func filterByRT(rts []string, records []MPLSL3Record) []MPLSL3Record {
	result := make([]MPLSL3Record, 0)
	match := 0
	for _, r := range records {
		for _, rrt := range strings.Split(r.RT, ",") {
			for _, rt := range rts {
				if rt == rrt {
					match++
				}
			}
		}
		// If number of matches == length of requested rts, then all requested rts were found
		// within a record's RT list.
		if match == len(rts) {
			result = append(result, r)
		}
		match = 0
	}

	return result
}

// NewMockDBClient returns an instance of a new mock database client process
func NewMockDBClient(mpls bool, fn ...string) dbclient.DBClient {
	// Need to load test data
	tfn := "./testdata.json"
	if fn[0] != "" {
		tfn = fn[0]
	}
	d, err := os.Open(tfn)
	if err != nil {
		glog.Errorf("failed to open %s with error: %+v", tfn, err)
		return nil
	}
	fi, err := d.Stat()
	if err != nil {
		glog.Errorf("failed to get file info of %s with error: %+v", tfn, err)
		return nil
	}
	l := fi.Size()
	b := make([]byte, l)
	if _, err := io.ReadFull(d, b); err != nil {
		glog.Errorf("failed to read testdata.json with error: %+v", err)
		return nil
	}
	vpn := make([]MPLSL3Record, 0)
	srv6 := make([]SRv6L3Record, 0)
	ds := mockSrv{
		vpnStore:  make(map[string][]MPLSL3Record),
		srv6Store: make(map[string][]SRv6L3Record),
	}
	if mpls {
		if err := json.Unmarshal(b, &vpn); err != nil {
			glog.Errorf("failed to unmarshal testdata with error: %+v", err)
			return nil
		}
	} else {
		if err := json.Unmarshal(b, &srv6); err != nil {
			glog.Errorf("failed to unmarshal testdata with error: %+v", err)
			return nil
		}
	}
	if mpls {
		for _, r := range vpn {
			if _, ok := ds.vpnStore[r.RD]; !ok {
				ds.vpnStore[r.RD] = make([]MPLSL3Record, 0)
			}
			ds.vpnStore[r.RD] = append(ds.vpnStore[r.RD], r)
		}
	} else {
		for _, r := range srv6 {
			if _, ok := ds.srv6Store[r.VPNRD]; !ok {
				ds.srv6Store[r.VPNRD] = make([]SRv6L3Record, 0)
			}
			ds.srv6Store[r.VPNRD] = append(ds.srv6Store[r.VPNRD], r)
		}
	}

	return &ds
}
