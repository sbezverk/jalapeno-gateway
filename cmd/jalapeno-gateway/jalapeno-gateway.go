package main

import (
	"fmt"
	"os"
	"os/signal"

	arango "github.com/sbezverk/jalapeno-gateway/pkg/dbclient/arangoclient"
	"github.com/sbezverk/jalapeno-gateway/pkg/srvclient"
)

func main() {
	addr := "http://10.200.99.3:30852"
	db, err := srvclient.NewSrvClient(addr, arango.NewArangoSrv("root", "jalapeno", "jalapeno", "L3VPN_Prefixes"))
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
