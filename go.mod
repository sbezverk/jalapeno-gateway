module jalapeno-gateway

go 1.13

require (
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
	github.com/kr/pretty v0.1.0 // indirect
	github.com/sbezverk/jalapeno-gateway/pkg/bgpclient v0.0.0-00010101000000-000000000000
	github.com/sbezverk/jalapeno-gateway/pkg/dbclient v0.0.0-00010101000000-000000000000
	github.com/sbezverk/jalapeno-gateway/pkg/dbclient/arangoclient v0.0.0-00010101000000-000000000000
	github.com/sbezverk/jalapeno-gateway/pkg/gateway v0.0.0-00010101000000-000000000000
	github.com/sbezverk/jalapeno-gateway/pkg/srvclient v0.0.0-00010101000000-000000000000
	github.com/stretchr/testify v1.4.0 // indirect
	gopkg.in/check.v1 v1.0.0-20180628173108-788fd7840127 // indirect
	gopkg.in/yaml.v2 v2.2.4 // indirect
)

replace (
	github.com/sbezverk/jalapeno-gateway/pkg/bgpclient => ./pkg/bgpclient
	github.com/sbezverk/jalapeno-gateway/pkg/dbclient => ./pkg/dbclient
	github.com/sbezverk/jalapeno-gateway/pkg/dbclient/arangoclient => ./pkg/dbclient/arangoclient
	github.com/sbezverk/jalapeno-gateway/pkg/dbclient/dbmockclient => ./pkg/dbclient/dbmockclient
	github.com/sbezverk/jalapeno-gateway/pkg/gateway => ./pkg/gateway
	github.com/sbezverk/jalapeno-gateway/pkg/srvclient => ./pkg/srvclient
)
