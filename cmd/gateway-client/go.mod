module gateway-client

go 1.13

require (
	github.com/gogo/protobuf v1.3.1 // indirect
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
	github.com/golang/protobuf v1.4.0 // indirect
	github.com/sbezverk/jalapeno-gateway/pkg/apis v0.0.0-20200424215720-986760d824da
	github.com/sbezverk/jalapeno-gateway/pkg/bgpclient v0.0.0-20200424215720-986760d824da // indirect
	golang.org/x/net v0.0.0-20200421231249-e086a090c8fd // indirect
	golang.org/x/sys v0.0.0-20200420163511-1957bb5e6d1f // indirect
	golang.org/x/text v0.3.2 // indirect
	google.golang.org/genproto v0.0.0-20200424135956-bca184e23272 // indirect
	google.golang.org/grpc v1.29.1
)

replace (
	github.com/sbezverk/jalapeno-gateway/pkg/apis => ../../pkg/apis
	github.com/sbezverk/jalapeno-gateway/pkg/bgpclient => ../../pkg/bgpclient
	github.com/sbezverk/jalapeno-gateway/pkg/srvclient => ../../pkg/srvclient
)
