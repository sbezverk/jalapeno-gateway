module jalapeno-gateway

go 1.13

require (
	github.com/arangodb/go-driver v0.0.0-20200403100147-ca5dd87ffe93 // indirect
	github.com/gogo/protobuf v1.3.1 // indirect
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
	github.com/golang/protobuf v1.4.0 // indirect
	github.com/kr/pretty v0.1.0 // indirect
	github.com/sbezverk/jalapeno-gateway/pkg/apis v0.0.0-20200424214153-bc697394c78d // indirect
	github.com/sbezverk/jalapeno-gateway/pkg/bgpclient v0.0.0-20200424215720-986760d824da
	github.com/sbezverk/jalapeno-gateway/pkg/dbclient v0.0.0-20200424215720-986760d824da // indirect
	github.com/sbezverk/jalapeno-gateway/pkg/dbclient/arangoclient v0.0.0-20200424215720-986760d824da
	github.com/sbezverk/jalapeno-gateway/pkg/dbclient/dbmockclient v0.0.0-20200424215720-986760d824da
	github.com/sbezverk/jalapeno-gateway/pkg/gateway v0.0.0-20200424215720-986760d824da
	github.com/sbezverk/jalapeno-gateway/pkg/srvclient v0.0.0-20200424215720-986760d824da
	golang.org/x/net v0.0.0-20200421231249-e086a090c8fd // indirect
	golang.org/x/sys v0.0.0-20200420163511-1957bb5e6d1f // indirect
	golang.org/x/text v0.3.2 // indirect
	google.golang.org/genproto v0.0.0-20200424135956-bca184e23272 // indirect
	google.golang.org/grpc v1.29.1 // indirect
	gopkg.in/check.v1 v1.0.0-20180628173108-788fd7840127 // indirect
	gopkg.in/yaml.v2 v2.2.4 // indirect
)

replace (
	github.com/sbezverk/jalapeno-gateway/pkg/apis => ./pkg/apis
	github.com/sbezverk/jalapeno-gateway/pkg/bgpclient => ./pkg/bgpclient
	github.com/sbezverk/jalapeno-gateway/pkg/dbclient => ./pkg/dbclient
	github.com/sbezverk/jalapeno-gateway/pkg/dbclient/arangoclient => ./pkg/dbclient/arangoclient
	github.com/sbezverk/jalapeno-gateway/pkg/dbclient/dbmockclient => ./pkg/dbclient/dbmockclient
	github.com/sbezverk/jalapeno-gateway/pkg/gateway => ./pkg/gateway
	github.com/sbezverk/jalapeno-gateway/pkg/srvclient => ./pkg/srvclient
)
