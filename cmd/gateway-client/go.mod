module gateway-client

go 1.13

require (
	github.com/gogo/protobuf v1.3.1 // indirect
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
	github.com/sbezverk/jalapeno-gateway/pkg/apis v0.0.0-00010101000000-000000000000
	github.com/sbezverk/jalapeno-gateway/pkg/bgpclient v0.0.0-00010101000000-000000000000 // indirect
	google.golang.org/grpc v1.27.1
)

replace (
	github.com/sbezverk/jalapeno-gateway/pkg/apis => ../../pkg/apis
	github.com/sbezverk/jalapeno-gateway/pkg/bgpclient => ../../pkg/bgpclient
	github.com/sbezverk/jalapeno-gateway/pkg/srvclient => ../../pkg/srvclient
)
