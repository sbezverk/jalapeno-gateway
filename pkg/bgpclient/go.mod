module github.com/sbezverk/jalapeno-gateway/pkg/bgpclient

go 1.14

require (
	github.com/osrg/gobgp v2.0.0+incompatible
	github.com/sbezverk/jalapeno-gateway/pkg/srvclient v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.27.1
)

replace github.com/sbezverk/jalapeno-gateway/pkg/srvclient => ../srvclient
