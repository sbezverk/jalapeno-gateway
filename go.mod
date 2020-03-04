module jalapeno-gateway

go 1.13

require (
	github.com/sbezverk/jalapeno-gateway/pkg/bgpclient v0.0.0-00010101000000-000000000000
	github.com/sbezverk/jalapeno-gateway/pkg/dbclient v0.0.0-00010101000000-000000000000
	github.com/sbezverk/jalapeno-gateway/pkg/dbclient/dbmockclient v0.0.0-00010101000000-000000000000
	github.com/coredns/coredns v1.6.7 // indirect
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
	github.com/google/go-cmp v0.4.0 // indirect
	github.com/osrg/gobgp v2.0.0+incompatible // indirect
	golang.org/x/net v0.0.0-20200114155413-6afb5195e5aa // indirect
	golang.org/x/sys v0.0.0-20200113162924-86b910548bc1 // indirect
	google.golang.org/genproto v0.0.0-20200115191322-ca5a22157cba // indirect
	google.golang.org/grpc v1.26.0
)

replace (
	github.com/sbezverk/jalapeno-gateway/pkg/bgpclient => ./pkg/bgpclient
	github.com/sbezverk/jalapeno-gateway/pkg/dbclient => ./pkg/dbclient
	github.com/sbezverk/jalapeno-gateway/pkg/dbclient/dbmockclient => ./pkg/dbclient/dbmockclient
)
