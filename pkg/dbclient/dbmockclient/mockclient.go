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
	pbapi "github.com/sbezverk/jalapeno-gateway/pkg/apis"
	"github.com/sbezverk/jalapeno-gateway/pkg/dbclient"
	"github.com/sbezverk/jalapeno-gateway/pkg/types"
)

const (
	srv6DataFile = "./testdata/testdata-srv6.json"
	vrfDataFile  = "./testdata/vrfs_data.json"
)

type mockSrv struct {
	srv6Store []*types.SRv6L3Record
	vrfStore  map[string]*types.VRF
}

func (m *mockSrv) SRv6L3VpnRequest(ctx context.Context, req *types.L3VpnReq) (*types.SRv6L3VpnResp, error) {
	if req.RT == "" {
		return nil, fmt.Errorf("route target must be specified in the request")
	}
	glog.V(5).Infof("requesting srv6 prefixes for a route target: %s", req.RT)
	srv6Prefix := make([]*pbapi.SRv6L3Prefix, 0)
	resp := types.SRv6L3VpnResp{
		Prefix: srv6Prefix,
	}
	records := make([]*types.SRv6L3Record, 0)
	records = filterByRTSRv6L3Record([]string{req.RT}, m.srv6Store)

	if len(records) == 0 {
		// All filtered, returning error
		return nil, fmt.Errorf("no matching records to found")
	}
	resp.Prefix = dbclient.GetSRv6Prefixes(records)

	return &resp, nil
}

func (m *mockSrv) VPNRTRequest(ctx context.Context, name string) (string, error) {
	glog.V(5).Infof("requesting a route target for vpn %q", name)
	vpn, ok := m.vrfStore[name]
	if !ok {
		return "", fmt.Errorf("vpn %s does not exist", name)
	}
	return dbclient.GetRT(vpn, name)
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

func filterByRTSRv6L3Record(rts []string, records []*types.SRv6L3Record) []*types.SRv6L3Record {
	result := make([]*types.SRv6L3Record, 0)
	found := false
	for _, r := range records {
		for _, extComm := range r.BaseAttributes.ExtCommunityList {
			if !strings.HasPrefix(extComm, bgp.ECPRouteTarget) {
				continue
			}
			for _, rrt := range rts {
				if rrt == strings.TrimPrefix(extComm, bgp.ECPRouteTarget) {
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
	srv6data, err := readTestFile(srv6DataFile)
	if err != nil {
		return nil
	}
	vrfdata, err := readTestFile(vrfDataFile)
	if err != nil {
		return nil
	}

	vrfs := make([]*types.VRF, 0)
	ds := mockSrv{
		srv6Store: make([]*types.SRv6L3Record, 0),
		vrfStore:  make(map[string]*types.VRF),
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
