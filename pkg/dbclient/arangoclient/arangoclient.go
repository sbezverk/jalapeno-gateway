package arangoclient

import (
	"context"
	"fmt"
	"math"
	"net"
	"net/url"
	"os"
	"os/signal"
	"strconv"
	"time"

	driver "github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"

	"github.com/sbezverk/jalapeno-gateway/pkg/srvclient"
)

var (
	arangoDBConnectTimeout = time.Duration(time.Second * 10)
)

// DBClient defines the interface for communication with DB process
type DBClient interface {
}

type arangoSrv struct {
	user       string
	pass       string
	dbName     string
	conn       driver.Connection
	client     driver.Client
	db         driver.Database
	collection string
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
	if net.ParseIP(host) == nil {
		return fmt.Errorf("fail to parse host part of address")
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

// NewArangoSrv returns an instance of a new Arango server process
func NewArangoSrv(user string, pass string, dbName string, collection string) srvclient.Server {
	return &arangoSrv{
		user:       user,
		pass:       pass,
		dbName:     dbName,
		collection: collection,
	}
}
func main() {
	addr := "http://10.200.99.3:30852"
	db, err := srvclient.NewSrvClient(addr, NewArangoSrv("root", "jalapeno", "jalapeno", "L3VPN_Prefixes"))
	if err != nil {
		fmt.Printf("failed to instantiate new Arango client with error: %+v\n", err)
		os.Exit(1)

	}
	db.Connect()
	sigc := make(chan os.Signal)
	signal.Notify(sigc, os.Interrupt)
	sig := <-sigc
	fmt.Printf("received %v\n", sig)
	db.Disconnect()
}
