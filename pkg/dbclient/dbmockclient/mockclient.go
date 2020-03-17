package dbmockclient

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/golang/glog"
	"github.com/sbezverk/jalapeno-gateway/pkg/dbclient"
)

// Record represents the database record structure
type Record struct {
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
	PrefixSID       uint32 `json:"PrefixSID,omitempty"`
	VPN             uint32 `json:"VPN_Label,omitempty"`
	RD              string `json:"RD"`
	RT              string `json:"RT,omitempty"`
	Source          string `json:"Source,omitempty"`
	Destination     string `json:"Destination,omitempty"`
}

type mockSrv struct {
	vpnStore map[string][]Record
}

func (m *mockSrv) L3VPNRequest(ctx context.Context, req *dbclient.L3VpnReq) (*dbclient.L3VpnResp, error) {
	glog.V(5).Infof("Mock DB L3 VPN Service was called with the request: %+v", req)

	// Initial lookup for requested RD, if it is not in the store, return error
	records, ok := m.vpnStore[req.RD]
	if !ok {
		return nil, fmt.Errorf("RD %s is not found", req.RD)
	}
	// Requested RD was found, applying RT and Prefix optional constraints
	if req.Prefix != "" {
		records = filterByPrefix(req.Prefix, req.MaskLength, records)
	}

	if len(req.RT) != 0 {
		records = filterByRT(req.RT, records)
	}

	if len(records) == 0 {
		// All filtered, returning error
		return nil, fmt.Errorf("no matching records to found")
	}

	resp := dbclient.L3VpnResp{
		VpnLabel: records[0].VPN,
		SidLabel: records[0].PrefixSID,
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

func filterByPrefix(prefix string, mask uint32, records []Record) []Record {
	result := make([]Record, 0)
	for _, r := range records {
		if r.Prefix == prefix && r.Mask == mask {
			result = append(result, r)
			break
		}

	}
	return result
}

func filterByRT(rts []string, records []Record) []Record {
	result := make([]Record, 0)
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
func NewMockDBClient(fn ...string) dbclient.DBClient {
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
	records := make([]Record, 0)
	if err := json.Unmarshal(b, &records); err != nil {
		glog.Errorf("failed to unmarshal testdata with error: %+v", err)
		return nil
	}

	ds := mockSrv{
		vpnStore: make(map[string][]Record, 0),
	}
	for _, r := range records {
		if _, ok := ds.vpnStore[r.RD]; !ok {
			ds.vpnStore[r.RD] = make([]Record, 0)
		}
		rds := ds.vpnStore[r.RD]
		rds = append(rds, r)
		ds.vpnStore[r.RD] = rds
	}

	return &ds
}
