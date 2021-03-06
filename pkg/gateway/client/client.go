package client

import (
	"sync"
)

// Store defines methods to oprate a clients' store
type Store interface {
	Add(string)
	Delete(string)
	Get(string) *Client
}

// Client defines a structure of a Gateway client
type Client struct {
	id string
	// TODO add a list of resources created for a client
	// and required to release/delete/withdraw when a client terminates
	sync.Mutex
	routeCleanup []func() error
}

// AddRouteCleanup appends a function which will be called upon client's termination. The
// idea is to call corresponding route delete for all routes which were added for the life time of the client.
func (c *Client) AddRouteCleanup(f func() error) {
	c.Lock()
	defer c.Unlock()
	c.routeCleanup = append(c.routeCleanup, f)
}

// GetRouteCleanup returns all recorded route cleanup callbacks
func (c *Client) GetRouteCleanup() []func() error {
	c.Lock()
	defer c.Unlock()

	return c.routeCleanup
}

type store struct {
	sync.Mutex
	// client is a store of clients in a form of a map, the key is a unique client id
	clients map[string]*Client
}

func (s *store) Add(id string) {
	s.Lock()
	defer s.Unlock()
	s.clients[id] = &Client{
		id:           id,
		routeCleanup: make([]func() error, 0),
	}
}

func (s *store) Delete(id string) {
	s.Lock()
	defer s.Unlock()
	delete(s.clients, id)
}

func (s *store) Get(id string) *Client {
	s.Lock()
	defer s.Unlock()
	c, ok := s.clients[id]
	if !ok {
		return nil
	}

	return c
}

// NewClientStore return a new instance of the client store.
func NewClientStore() Store {
	return &store{
		clients: make(map[string]*Client, 0),
	}
}
