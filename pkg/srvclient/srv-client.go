package srvclient

import (
	"fmt"
	"sync"
	"time"

	"github.com/golang/glog"
)

// SrvClient defines methods to interact with a server process
type SrvClient interface {
	GetStatus() SrvStatus
	SetStatus(SrvStatus)
	GetClientInterface() interface{}
	Connect()
	Disconnect()
	Addr() string
}

var (
	reconnectTimeout  int = 5
	keepaliveInterval int = 10
)

// SrvStatus defines a server status enumeration.
type SrvStatus int32

const (
	// UP indicates the client is connected to a server process
	UP SrvStatus = iota
	// DOWN indicates theclient is not connected to a server process
	DOWN
	// CONNECT indicates the client is not connected to a server process, but in the process of reconnecting
	CONNECT
)

type srvClient struct {
	srvAddr       string
	srv           Server
	reconnect     chan struct{}
	work          chan struct{}
	stopConnector chan struct{}
	stopMonitor   chan struct{}
	sync.Mutex
	status      SrvStatus
	connectorUP bool
	monitorUP   bool
	sync.WaitGroup
}

func (b *srvClient) GetClientInterface() interface{} {
	return b.srv
}

func (b *srvClient) GetStatus() SrvStatus {
	b.Lock()
	defer b.Unlock()
	s := b.status
	return s
}

func (b *srvClient) SetStatus(s SrvStatus) {
	b.Lock()
	defer b.Unlock()
	b.status = s
}

func (b *srvClient) Connect() {
	if b.GetStatus() == DOWN {
		b.stopConnector = make(chan struct{})
		b.stopMonitor = make(chan struct{})
		go b.connector()
		go b.monitor()
		b.WaitGroup.Add(2)
	}
}

func (b *srvClient) Disconnect() {
	glog.V(5).Infof("Shutting down the client for the server: %s", b.Addr())
	b.stopMonitor <- struct{}{}
	b.stopConnector <- struct{}{}
	b.WaitGroup.Wait()
	glog.V(5).Infof("Shutdown of the client for the server: %s completed.", b.Addr())
	b.SetStatus(DOWN)
}

func (b *srvClient) connector() {
	if b.connectorUP {
		return
	}
	b.Lock()
	b.connectorUP = true
	b.Unlock()
	defer func() {
		b.Lock()
		b.connectorUP = false
		b.Unlock()
		glog.V(5).Infof("Connector for the server: %s exiting", b.Addr())
		b.WaitGroup.Done()
	}()
	// Main connector loop
	for {
		// Initially wait on either reconnect request or stop
		select {
		case <-b.stopConnector:
			return
		case <-b.reconnect:
			glog.V(5).Infof("Connector for the server: %s received connect request.", b.Addr())
			b.SetStatus(CONNECT)
			for b.GetStatus() != UP {
				// Attempting to connect to server
				if err := b.srv.Connector(b.Addr()); err == nil {
					b.SetStatus(UP)
					glog.V(5).Infof("Connection to the server: %s succeeded", b.Addr())
				} else {
					timeout := time.NewTimer(time.Second * time.Duration(reconnectTimeout))
					glog.Errorf("Connector failed to reconnect to the server: %s with error: %+v, retrying in %d seconds.", b.Addr(), err, reconnectTimeout)
					select {
					case <-b.stopConnector:
						return
					case <-timeout.C:
					}
				}
			}
		}
		// Connection came to gobgpd was restored, sending messge to monitor
		b.work <- struct{}{}
	}
}

func (b *srvClient) monitor() {
	if b.monitorUP {
		return
	}
	b.Lock()
	b.monitorUP = true
	b.Unlock()
	defer func() {
		b.Lock()
		b.monitorUP = false
		b.Unlock()
		glog.V(5).Infof("Monitor for the server: %s exiting", b.Addr())
		b.WaitGroup.Done()
	}()
	for {
		status := b.GetStatus()
		switch status {
		case DOWN:
			// Sending message to connector to reconnect to gobgpd
			b.reconnect <- struct{}{}
			select {
			case <-b.work:
				glog.V(5).Infof("Connection to the server: %s succeeded resuming keepalives", b.Addr())
			case <-b.stopMonitor:
				return
			}
		case UP:
			if err := b.srv.Monitor(b.Addr()); err != nil {
				glog.Errorf("keepalive with the server %s failed with error: %+v", b.Addr(), err)
				b.SetStatus(DOWN)
			} else {
				timeout := time.NewTimer(time.Second * time.Duration(keepaliveInterval))
				glog.V(6).Infof("keepalive with the server %s succeeded, next keepalive in %d seconds.", b.Addr(), keepaliveInterval)
				select {
				case <-b.stopMonitor:
					return
				case <-timeout.C:
				}
			}
		case CONNECT:
		}
	}
}

func (b *srvClient) Addr() string {
	return b.srvAddr
}

// NewSrvClient return a new instance of bgp client
func NewSrvClient(addr string, srv Server) (SrvClient, error) {
	if srv == nil {
		return nil, fmt.Errorf("server's interface cannot be nil")
	}
	if err := srv.Validator(addr); err != nil {
		return nil, fmt.Errorf("fail to validate address with error %w", err)
	}
	return &srvClient{
		srvAddr:   addr,
		srv:       srv,
		status:    DOWN,
		work:      make(chan struct{}),
		reconnect: make(chan struct{}),
	}, nil
}

// Server defines methods to connect and to check liveness of the server process
type Server interface {
	Validator(string) error
	Monitor(string) error
	Connector(string) error
}
