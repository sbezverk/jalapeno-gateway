package dbmockclient

package arangoclient

import (
	"context"
	"fmt"
	"math"
	"net"
	"net/url"
	"strconv"
	"time"

	"github.com/golang/glog"
	"github.com/sbezverk/jalapeno-gateway/pkg/dbclient"
)

type Record struct {
	Key string   `json:"_key,omitempty"` 
	Id string   `json:"_id,omitempty"`
	From string   `json:"_from,omitempty"`
	To string   `json:"_to,omitempty"`
	Rev string   `json:"_rev,omitempty"`
	SourceAddr string   `json:"SrcIP,omitempty"`
	DestinationAddr   `json:"DstIP,omitempty"`
	Prefix   string `json:"VPN_Prefix,omitempty"`
	Mask uint32   `json:"VPN_Prefix_Len,omitempty"`
	RouterID string   `json:"RouterID,omitempty"`
	PrefixSID string   `json:"PrefixSID,omitempty"`
	VPN string   `json:"VPN_Label,omitempty"`
	RD string   `json:"RD"`
	RT string    `json:"RT,omitempty"`
	Source string   `json:"Source,omitempty"`
	Destination string   `json:"Destination,omitempty"`
   }

type mockSrv struct {
}

func (m *mockSrv) L3VPNRequest(ctx context.Context, req *dbclient.L3VpnReq) (*dbclient.L3VpnResp, error) {
	glog.V(5).Infof("Arango DB L3 VPN Service was called with the request: %+v", req)
	return &dbclient.L3VpnResp{VpnLabel: uint32(24001), SidLabel: uint32(10004)}, nil
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

// NewMockDBClient returns an instance of a new mock database client process
func NewMockDBClient() dbclient.DBClient {
	return &mockSrv{
	}
}
