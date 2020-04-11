package arangoclient

import (
	"context"
	"fmt"

	driver "github.com/arangodb/go-driver"
	"github.com/golang/glog"
	"github.com/sbezverk/jalapeno-gateway/pkg/dbclient"
)

func (a *arangoSrv) L3VPNRequest(ctx context.Context, req *dbclient.L3VpnReq) (*dbclient.L3VpnResp, error) {
	glog.V(5).Infof("Arango DB L3 VPN Service was called with the request: %+v", req)

	filters := buildFilter(req)
	query, bindVars := buildQuery(a.collection, filters...)
	if err := a.db.ValidateQuery(context.Background(), query); err != nil {
		return &dbclient.L3VpnResp{}, fmt.Errorf("ValidateQuery failed with error: %+v", err)
	}
	prefix, err := runQuery(a.db, query, bindVars)
	if err != nil {
		return &dbclient.L3VpnResp{}, fmt.Errorf("runQuery failed with error: %+v", err)
	}
	glog.Infof("Prefixes: %+v", prefix)

	return &dbclient.L3VpnResp{
		Prefix: prefix,
	}, nil
}

func buildFilter(req *dbclient.L3VpnReq) []filter {
	filters := make([]filter, 0)
	filters = append(filters, filter{key: "ipv4", value: req.IPv4})
	if req.RD != "" {
		filters = append(filters, filter{key: "rd", value: req.RD})
	}
	if req.Prefix != "" {
		filters = append(filters, filter{key: "prefix", value: req.Prefix})
	}
	if req.MaskLength != 0 {
		filters = append(filters, filter{key: "mask", value: req.MaskLength})
	}

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
			query += fmt.Sprintf("q.RD == @rd ")
			bindVars[f.key] = f.value
		case "ipv4":
			query += fmt.Sprintf("q.IPv4 == @ipv4 ")
			bindVars[f.key] = f.value
		case "mask":
		case "prefix":
		}
		if i < len(filters)-1 {
			query += "and "
		}
	}
	query += "return q"

	return query, bindVars
}

func runQuery(db driver.Database, query string, bindVars map[string]interface{}) ([]dbclient.L3VPNPrefix, error) {
	cursor, err := db.Query(context.TODO(), query, bindVars)
	if err != nil {
		return nil, fmt.Errorf("Query failed with error: %+v", err)
	}
	defer cursor.Close()
	i := make([]dbclient.L3VPNPrefix, 0)
	for {
		var reply dbclient.L3VPNPrefix
		_, err = cursor.ReadDocument(context.Background(), &reply)
		if driver.IsNoMoreDocuments(err) {
			break
		} else if err != nil {
			return nil, fmt.Errorf("ReadDocument failed with error: %+v", err)
		}
		i = append(i, reply)
	}

	return i, nil
}
