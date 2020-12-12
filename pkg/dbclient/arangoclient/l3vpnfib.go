package arangoclient

import (
	"context"
	"fmt"

	driver "github.com/arangodb/go-driver"
	"github.com/golang/glog"
	"github.com/sbezverk/jalapeno-gateway/pkg/dbclient"
	"github.com/sbezverk/jalapeno-gateway/pkg/types"
)

// RTRecord defines route target record
type RTRecord struct {
	ID       string            `json:"_id,omitempty"`
	Key      string            `json:"_key,omitempty"`
	RT       string            `json:"RT,omitempty"`
	Prefixes map[string]string `json:"Prefixes,omitempty"`
}

func (a *arangoSrv) VPNRTRequest(ctx context.Context, name string) (string, error) {
	glog.V(5).Infof("requesting a route target for vpn %q", name)
	vrf, err := getCollection(ctx, a, a.vrfCollection)
	if err != nil {
		return "", err
	}
	vpn := &types.VRF{}
	if _, err := vrf.ReadDocument(ctx, name, vpn); err != nil {
		glog.Errorf("failed to get vpn %s's route target with error: %+v", name, err)
		return "", fmt.Errorf("failed to get vpn %s's route target with error: %+v", name, err)
	}

	return dbclient.GetRT(vpn, name)
}

func getCollection(ctx context.Context, a *arangoSrv, name string) (driver.Collection, error) {
	found, err := a.db.CollectionExists(ctx, name)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, fmt.Errorf("collection %s does not exist", name)
	}

	return a.db.Collection(ctx, name)
}

func (a *arangoSrv) SRv6L3VpnRequest(ctx context.Context, req *types.L3VpnReq) (*types.SRv6L3VpnResp, error) {
	glog.V(5).Infof("requesting srv6 prefixes for a route target: %s", req.RT)
	fib, err := getCollection(ctx, a, a.fibCollection)
	if err != nil {
		return nil, err
	}
	rt, err := getCollection(ctx, a, a.rtCollection)
	if err != nil {
		return nil, err
	}

	rtr := &RTRecord{}
	if _, err := rt.ReadDocument(ctx, req.RT, rtr); err != nil {
		return nil, err
	}
	glog.V(5).Infof("found %d prefixes a route target %s", len(rtr.Prefixes), req.RT)
	resp := &types.SRv6L3VpnResp{}
	records := make([]*types.SRv6L3Record, 0, len(rtr.Prefixes))
	for _, v := range rtr.Prefixes {
		r := &types.SRv6L3Record{}
		if _, err := fib.ReadDocument(ctx, v, r); err != nil {
			glog.Errorf("prefix by key %q is not found in collection %q", v, fib.Name())
			glog.Errorf("inconsistency between l3vpn route target collection %q and l3vpn prefix collection %q found, please report a software bug", rt.Name(), fib.Name())
			continue
		}
		records = append(records, r)
	}
	resp.Prefix = dbclient.GetSRv6Prefixes(records)
	glog.V(5).Infof("processed %d prefixes for a route target %s", len(records), req.RT)

	return resp, nil
}
