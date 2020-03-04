package dbclient

import (
	"github.com/sbezverk/jalapeno-gateway/pkg/srvclient"
)

// DBClient defines methods a particular database client must implement
type DBClient interface {
	// Embeding Server interface
	srvclient.Server
	// Additional Database specific methods
}
