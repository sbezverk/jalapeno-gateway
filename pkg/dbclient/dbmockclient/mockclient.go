package dbmockclient

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"os"
	"strconv"
	"strings"

	"github.com/golang/glog"
	"github.com/golang/protobuf/ptypes/any"
	pbapi "github.com/sbezverk/jalapeno-gateway/pkg/apis"
	"github.com/sbezverk/jalapeno-gateway/pkg/bgpclient"
	"github.com/sbezverk/jalapeno-gateway/pkg/dbclient"
)

type mockSrv struct {
	vpnStore  map[string][]dbclient.MPLSL3Record
	srv6Store map[string][]dbclient.SRv6L3Record
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
	glog.V(5).Infof("Mock DB SRv6 VPN Service was called for RD: %s", req.RD)
	srv6Prefix := make([]*pbapi.SRv6L3Prefix, 0)
	resp := dbclient.SRv6L3VpnResp{
		Prefix: srv6Prefix,
	}
	// Initial lookup for requested RD, if it is not in the store, return an empty response
	records, ok := m.srv6Store[req.RD]
	if !ok {
		return &resp, nil
	}
	// All filtered, return an empty response
	if len(records) == 0 {
		return &resp, nil
	}

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
		asn, err := strconv.Atoi(r.OriginAS)
		if err != nil {
			continue
		}
		p.Asn = uint32(asn)
		p.PrefixSid.Tlvs = bgpclient.MarshalPrefixSID(r.PrefixSID)
		rd, err := bgpclient.MarshalRDFromString(r.RD)
		if err != nil {
			continue
		}
		p.Rd = rd
		rts := make([]*any.Any, 0)
		// for _, extcomm := range strings.Split(r.BaseAttributes.ExtCommunityList, ",") {
		// 	if !strings.HasPrefix(extcomm, "rt=") {
		// 		continue
		// 	}
		// 	rt, err := bgpclient.MarshalRTFromString(strings.Split(extcomm, "=")[1])
		// 	if err != nil {
		// 		continue
		// 	}
		// 	rts = append(rts, rt)
		// }
		for _, s := range r.RT {
			rt, err := bgpclient.MarshalRTFromString(s)
			if err != nil {
				continue
			}
			rts = append(rts, rt)
		}
		p.Rt = rts
		srv6Prefix = append(srv6Prefix, p)
	}
	resp.Prefix = srv6Prefix

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

func filterByIPFamily(ipv4 bool, records []dbclient.MPLSL3Record) []dbclient.MPLSL3Record {
	result := make([]dbclient.MPLSL3Record, 0)
	for _, r := range records {
		if r.IPv4 == ipv4 {
			result = append(result, r)
		}
	}

	return result
}
func filterByPrefix(prefix string, mask uint32, records []dbclient.MPLSL3Record) []dbclient.MPLSL3Record {
	result := make([]dbclient.MPLSL3Record, 0)
	for _, r := range records {
		if r.Prefix == prefix && r.Mask == mask {
			result = append(result, r)
			break
		}

	}

	return result
}

func filterByRT(rts []string, records []dbclient.MPLSL3Record) []dbclient.MPLSL3Record {
	result := make([]dbclient.MPLSL3Record, 0)
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
	vpn := make([]dbclient.MPLSL3Record, 0)
	srv6 := make([]dbclient.SRv6L3Record, 0)
	ds := mockSrv{
		vpnStore:  make(map[string][]dbclient.MPLSL3Record),
		srv6Store: make(map[string][]dbclient.SRv6L3Record),
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
				ds.vpnStore[r.RD] = make([]dbclient.MPLSL3Record, 0)
			}
			ds.vpnStore[r.RD] = append(ds.vpnStore[r.RD], r)
		}
	} else {
		for _, r := range srv6 {
			if _, ok := ds.srv6Store[r.RD]; !ok {
				ds.srv6Store[r.RD] = make([]dbclient.SRv6L3Record, 0)
			}
			ds.srv6Store[r.RD] = append(ds.srv6Store[r.RD], r)
		}
	}

	return &ds
}
