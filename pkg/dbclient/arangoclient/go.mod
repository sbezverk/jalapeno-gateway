module arangoclient

go 1.13

replace (
	github.com/sbezverk/jalapeno-gateway => ../../../../jalapeno-gateway
	github.com/sbezverk/jalapeno-gateway/pkg/dbclient => ../../dbclient
	github.com/sbezverk/jalapeno-gateway/pkg/srvclient => ../../srvclient
)

require (
	github.com/arangodb/go-driver v0.0.0-20200226154541-eb7d8400480f
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
	github.com/osrg/gobgp v2.0.0+incompatible
	github.com/sbezverk/jalapeno-gateway/pkg/dbclient v0.0.0-00010101000000-000000000000
	github.com/sbezverk/jalapeno-gateway/pkg/srvclient v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.27.1
	k8s.io/klog v1.0.0
)
