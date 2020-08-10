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
	"github.com/sbezverk/gobmp/pkg/prefixsid"
	pbapi "github.com/sbezverk/jalapeno-gateway/pkg/apis"
	"github.com/sbezverk/jalapeno-gateway/pkg/bgpclient"
	"github.com/sbezverk/jalapeno-gateway/pkg/dbclient"
	"github.com/sbezverk/jalapeno-gateway/pkg/types"
)

const (
	mplsDataFile = "./testdata/testdata-mpls.json"
	srv6DataFile = "./testdata/testdata-srv6.json"
	vrfDataFile  = "./testdata/vrfs_data.json"
)

type mockSrv struct {
	mplsStore []types.MPLSL3Record
	srv6Store []types.SRv6L3Record
	vrfStore  map[string]types.VRF
}

func (m *mockSrv) MPLSL3VpnRequest(ctx context.Context, req *types.L3VpnReq) (*types.MPLSL3VpnResp, error) {
	if req.RT == "" {
		return nil, fmt.Errorf("route target must be specified in the request")
	}
	glog.V(5).Infof("Mock DB L3 VPN Service was called with the request: %+v", req)
	records := make([]types.MPLSL3Record, 0)

	// Filter by IP Family
	records = filterByIPFamily(req.IPv4, records)

	// Filter by RT
	records = filterByRT([]string{req.RT}, records)

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
	resp := types.MPLSL3VpnResp{
		Prefix: vpnPrefix,
	}

	return &resp, nil
}

func (m *mockSrv) SRv6L3VpnRequest(ctx context.Context, req *types.L3VpnReq) (*types.SRv6L3VpnResp, error) {
	if req.RT == "" {
		return nil, fmt.Errorf("route target must be specified in the request")
	}
	glog.V(5).Infof("Mock DB SRv6 VPN Service was called for VRF RT: %s", req.RT)
	srv6Prefix := make([]*pbapi.SRv6L3Prefix, 0)
	resp := types.SRv6L3VpnResp{
		Prefix: srv6Prefix,
	}
	records := make([]types.SRv6L3Record, 0)
	records = filterByRTSRv6L3Record([]string{req.RT}, records)

	if len(records) == 0 {
		// All filtered, returning error
		return nil, fmt.Errorf("no matching records to found")
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
		p.PrefixSid.Tlvs = bgpclient.MarshalPrefixSID(&prefixsid.PSid{SRv6L3Service: r.SRv6SID})
		rts := make([]*any.Any, 0)
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

func (m *mockSrv) VPNRTRequest(ctx context.Context, name string) (string, error) {
	vpn, ok := m.vrfStore[name]
	if !ok {
		return "", fmt.Errorf("vpn %s does not exist", name)
	}
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

func (m *mockSrv) Connector(addr string) error {

	return nil
}

func (m *mockSrv) Monitor(addr string) error {
	return nil
}

func (m *mockSrv) Validator(addr string) error {
	return nil
}

func filterByIPFamily(ipv4 bool, records []types.MPLSL3Record) []types.MPLSL3Record {
	result := make([]types.MPLSL3Record, 0)
	for _, r := range records {
		if r.IPv4 == ipv4 {
			result = append(result, r)
		}
	}

	return result
}
func filterByPrefix(prefix string, mask uint32, records []types.MPLSL3Record) []types.MPLSL3Record {
	result := make([]types.MPLSL3Record, 0)
	for _, r := range records {
		if r.Prefix == prefix && r.Mask == mask {
			result = append(result, r)
			break
		}

	}

	return result
}

func filterByRT(rts []string, records []types.MPLSL3Record) []types.MPLSL3Record {
	result := make([]types.MPLSL3Record, 0)
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

func filterByRTSRv6L3Record(rts []string, records []types.SRv6L3Record) []types.SRv6L3Record {
	result := make([]types.SRv6L3Record, 0)
	found := false
	for _, r := range records {
		for _, ert := range r.RT {
			for _, rrt := range rts {
				if rrt == ert {
					result = append(result, r)
					found = true
					break
				}
			}
			if found {
				found = false
				break
			}
		}
	}

	return result
}

// NewMockDBClient returns an instance of a new mock database client process
func NewMockDBClient() dbclient.DBClient {
	// Need to load test data
	mplsdata, err := readTestFile(mplsDataFile)
	if err != nil {
		return nil
	}
	srv6data, err := readTestFile(srv6DataFile)
	if err != nil {
		return nil
	}
	vrfdata, err := readTestFile(vrfDataFile)
	if err != nil {
		return nil
	}

	vrfs := make([]types.VRF, 0)
	ds := mockSrv{
		mplsStore: make([]types.MPLSL3Record, 0),
		srv6Store: make([]types.SRv6L3Record, 0),
		vrfStore:  make(map[string]types.VRF),
	}

	if err := json.Unmarshal(mplsdata, &ds.mplsStore); err != nil {
		glog.Errorf("failed to unmarshal mpls test data with error: %+v", err)
		return nil
	}
	if err := json.Unmarshal(srv6data, &ds.srv6Store); err != nil {
		glog.Errorf("failed to unmarshal srv6 test data with error: %+v", err)
		return nil
	}
	if err := json.Unmarshal(vrfdata, &vrfs); err != nil {
		glog.Errorf("failed to unmarshal vrf test data with error: %+v", err)
		return nil
	}
	for _, r := range vrfs {
		ds.vrfStore[r.VRFName] = r
	}

	return &ds
}

func readTestFile(fn string) ([]byte, error) {
	d, err := os.Open(fn)
	if err != nil {
		return nil, fmt.Errorf("failed to open %s with error: %+v", fn, err)
	}
	fi, err := d.Stat()
	if err != nil {
		return nil, fmt.Errorf("failed to get file info of %s with error: %+v", fn, err)
	}
	l := fi.Size()
	b := make([]byte, l)
	if _, err := io.ReadFull(d, b); err != nil {
		return nil, fmt.Errorf("failed to read %s with error: %+v", fn, err)
	}

	return b, nil
}
