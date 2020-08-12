package arangoclient

import (
	"context"
	"fmt"

	driver "github.com/arangodb/go-driver"
	"github.com/golang/glog"
	"github.com/sbezverk/jalapeno-gateway/pkg/apis"
	"github.com/sbezverk/jalapeno-gateway/pkg/dbclient"
	"github.com/sbezverk/jalapeno-gateway/pkg/types"
)

const (
	vrfCollection = "Bell_VRF"
	fibCollection = "L3VPNPrefix"
	rtCollection  = "L3VPN_RT"
)

// RTRecord defines route target record
type RTRecord struct {
	ID       string            `json:"_id,omitempty"`
	Key      string            `json:"_key,omitempty"`
	RT       string            `json:"RT,omitempty"`
	Prefixes map[string]string `json:"Prefixes,omitempty"`
}

func (a *arangoSrv) VPNRTRequest(ctx context.Context, name string) (string, error) {
	glog.V(5).Infof("Arango DB VPN RT request for VPN: %s", name)
	vrf, err := getCollection(ctx, a, vrfCollection)
	if err != nil {
		return "", err
	}
	vpn := &types.VRF{}
	if _, err := vrf.ReadDocument(ctx, name, vpn); err != nil {
		return "", err
	}

	return dbclient.GetRT(vpn, name)
}

func getCollection(ctx context.Context, a *arangoSrv, name string) (driver.Collection, error) {
	found, err := a.db.CollectionExists(ctx, fibCollection)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, fmt.Errorf("collection %s does not exist", fibCollection)
	}

	return a.db.Collection(ctx, name)
}

func (a *arangoSrv) SRv6L3VpnRequest(ctx context.Context, req *types.L3VpnReq) (*types.SRv6L3VpnResp, error) {
	glog.V(5).Infof("Arango DB L3 VPN Service was called with the request: %s", req.RT)
	fib, err := getCollection(ctx, a, fibCollection)
	if err != nil {
		return nil, err
	}
	rt, err := getCollection(ctx, a, rtCollection)
	if err != nil {
		return nil, err
	}

	rtr := &RTRecord{}
	if _, err := rt.ReadDocument(ctx, req.RT, rtr); err != nil {
		return nil, err
	}
	resp := &types.SRv6L3VpnResp{}
	records := make([]*types.SRv6L3Record, len(rtr.Prefixes))
	i := 0
	for _, v := range rtr.Prefixes {
		r := &types.SRv6L3Record{}
		if _, err := fib.ReadDocument(ctx, v, r); err != nil {
			return nil, err
		}
		records[i] = r
		i++
	}
	resp.Prefix = dbclient.GetSRv6Prefixes(records)

	return resp, nil
}

func (a *arangoSrv) MPLSL3VpnRequest(ctx context.Context, req *types.L3VpnReq) (*types.MPLSL3VpnResp, error) {
	glog.V(5).Infof("Arango DB L3 VPN Service was called with the request: %+v", req)

	filters := buildFilter(req)
	query, bindVars := buildQuery(a.collection, filters...)
	if err := a.db.ValidateQuery(context.Background(), query); err != nil {
		return &types.MPLSL3VpnResp{}, fmt.Errorf("ValidateQuery failed with error: %+v", err)
	}
	prefix, err := runQuery(a.db, query, bindVars)
	if err != nil {
		return &types.MPLSL3VpnResp{}, fmt.Errorf("runQuery failed with error: %+v", err)
	}
	glog.Infof("Prefixes: %+v", prefix)

	return &types.MPLSL3VpnResp{
		Prefix: prefix,
	}, nil
}

func buildFilter(req *types.L3VpnReq) []filter {
	filters := make([]filter, 0)
	// filters = append(filters, filter{key: "ipv4", value: req.IPv4})
	// if req.RD != "" {
	// 	filters = append(filters, filter{key: "rd", value: req.RD})
	// }
	// if req.Prefix != "" {
	// 	filters = append(filters, filter{key: "prefix", value: req.Prefix})
	// }
	// if req.MaskLength != 0 {
	// 	filters = append(filters, filter{key: "mask", value: req.MaskLength})
	// }
	// if len(req.RT) != 0 {
	// 	filters = append(filters, filter{key: "rt", value: req.RT})
	// }

	return filters
}

func buildQuery(collection string, filters ...filter) (string, map[string]interface{}) {
	var query string
	var bindVars map[string]interface{}

	// Adding initial part of the query
	query += fmt.Sprintf("for q in %s ", collection)
	// If filters are provided, add corresponding filters to the query
	if len(filters) != 0 {
		bindVars = make(map[string]interface{}, 0)
		query += "filter "
	}
	for i, f := range filters {
		switch f.key {
		case "rd":
			query += "q.RD == @rd "
			bindVars[f.key] = f.value
		case "ipv4":
			query += "q.IPv4 == @ipv4 "
			bindVars[f.key] = f.value
		case "rt":
			// Since RT is a slice of strings, building filtering expression on the fly
			// no need for bindVars.
			query += "["
			rts := f.value.([]string)
			for i, rt := range rts {
				query += fmt.Sprintf("%q", rt)
				if i < len(rts)-1 {
					query += ","
				}
			}
			query += "] all in q.RT "
		case "prefix":
			query += "q.VPN_Prefix == @prefix "
			bindVars[f.key] = f.value
		case "mask":
			query += "q.VPN_Prefix_Len == @mask "
			bindVars[f.key] = f.value
		}
		if i < len(filters)-1 {
			query += "and "
		}
	}
	query += "return q"

	return query, bindVars
}

func runQuery(db driver.Database, query string, bindVars map[string]interface{}) ([]*apis.MPLSL3Prefix, error) {
	// cursor, err := db.Query(context.TODO(), query, bindVars)
	// if err != nil {
	// 	return nil, fmt.Errorf("Query failed with error: %+v", err)
	// }
	// defer cursor.Close()
	i := make([]*apis.MPLSL3Prefix, 0)
	// for {
	// 	var reply dbclient.L3VPNPrefix
	// 	_, err = cursor.ReadDocument(context.Background(), &reply)
	// 	if driver.IsNoMoreDocuments(err) {
	// 		break
	// 	} else if err != nil {
	// 		return nil, fmt.Errorf("ReadDocument failed with error: %+v", err)
	// 	}
	// 	i = append(i, reply)
	// }

	// return i, nil

	return i, nil
}
