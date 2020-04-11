package arangoclient

import (
	"context"
	"fmt"
	"math"
	"net"
	"net/url"
	"strconv"
	"time"

	driver "github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"
	"github.com/golang/glog"

	"github.com/sbezverk/jalapeno-gateway/pkg/dbclient"
)

var (
	arangoDBConnectTimeout = time.Duration(time.Second * 10)
)

type arangoSrv struct {
	user       string
	pass       string
	dbName     string
	conn       driver.Connection
	client     driver.Client
	db         driver.Database
	collection string
}

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

func (a *arangoSrv) Connector(addr string) error {
	conn, err := http.NewConnection(http.ConnectionConfig{
		Endpoints: []string{addr},
	})
	if err != nil {
		return err
	}
	c, err := driver.NewClient(driver.ClientConfig{
		Connection:     conn,
		Authentication: driver.BasicAuthentication(a.user, a.pass),
	})
	if err != nil {
		return err
	}
	a.conn = conn
	a.client = c
	ctx, cancel := context.WithTimeout(context.TODO(), arangoDBConnectTimeout)
	defer cancel()
	db, err := c.Database(ctx, a.dbName)
	if err != nil {
		return err
	}
	a.db = db

	return nil
}

func (a *arangoSrv) Monitor(addr string) error {
	_, err := a.db.CollectionExists(context.TODO(), a.collection)
	if err != nil {
		return err
	}

	return nil
}

func (a *arangoSrv) Validator(addr string) error {
	endpoint, err := url.Parse(addr)
	if err != nil {
		return err
	}
	host, port, _ := net.SplitHostPort(endpoint.Host)
	if host == "" || port == "" {
		return fmt.Errorf("host or port cannot be ''")
	}
	// Try to resolve if the hostname was used in the address
	if ip, err := net.LookupIP(host); err != nil || ip == nil {
		// Check if IP address was used in address instead of a host name
		if net.ParseIP(host) == nil {
			return fmt.Errorf("fail to parse host part of address")
		}
	}
	np, err := strconv.Atoi(port)
	if err != nil {
		return fmt.Errorf("fail to parse port with error: %w", err)
	}
	if np == 0 || np > math.MaxUint16 {
		return fmt.Errorf("the value of port is invalid")
	}
	return nil
}

// NewArangoDBClient returns an instance of a new Arango database client process
func NewArangoDBClient(user string, pass string, dbName string, collection string) dbclient.DBClient {
	return &arangoSrv{
		user:       user,
		pass:       pass,
		dbName:     dbName,
		collection: collection,
	}
}

type filter struct {
	key   string
	value interface{}
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
