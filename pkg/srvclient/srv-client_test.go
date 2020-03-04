package srvclient

import (
	"fmt"
	"math"
	"math/rand"
	"net"
	"strconv"
	"testing"
	"time"
)

type testSrv struct {
}

func (t *testSrv) Connector(addr string) error {
	// Simulating failing to reconnect if random generator return 3 or 7 or 9
	if n := rand.Intn(10); n == 3 || n == 7 || n == 9 {
		return fmt.Errorf("fail to reconnect")
	}

	return nil
}

func (t *testSrv) Monitor(addr string) error {
	// Simulating lost of connectivity if random generator returns 5 or 9
	if n := rand.Intn(10); n == 5 || n == 9 {
		return fmt.Errorf("connection lost")
	}

	return nil
}

func (t *testSrv) Validator(bgp string) error {
	host, port, _ := net.SplitHostPort(bgp)
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

// NewTestSrv returns an instance of a new test server process
func NewTestSrv() Server {
	return &testSrv{}
}

func TestSrvClient(t *testing.T) {
	addr := "192.168.80.103:5051"
	b, err := NewSrvClient(addr, NewTestSrv())
	if err != nil {
		t.Errorf("failed to instantiate new bgp client with error: %+v\n", err)

	}
	b.Connect()
	time.Sleep(time.Second * 20)
	//	sigc := make(chan os.Signal)
	//	signal.Notify(sigc, os.Interrupt)
	//	sig := <-sigc
	//	fmt.Printf("received %v\n", sig)
	b.Disconnect()
}
