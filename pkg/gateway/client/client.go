package client

import "sync"

// Store defines methods to oprate a clients' store
type Store interface {
	Add(string)
	Delete(string)
	Get(string) *Client
	Update(*Client) *Client
}

// Client defines a structure of a Gateway client
type Client struct {
	id string
	// TODO add a list of resources created for a client
	// and required to release/delete/withdraw when a client terminates
	data interface{}
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
		id: id,
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

func (s *store) Update(c *Client) *Client {
	s.Lock()
	defer s.Unlock()

	if _, ok := s.clients[c.id]; !ok {
		return nil
	}
	s.clients[c.id] = c

	return c
}

// NewClientStore return a new instance of the client store.
func NewClientStore() Store {
	return &store{
		clients: make(map[string]*Client, 0),
	}
}
