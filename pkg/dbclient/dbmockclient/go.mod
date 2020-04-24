module github.com/sbezverk/jalapeno-gateway/pkg/dbmockclient

go 1.14

replace (
	github.com/sbezverk/jalapeno-gateway => ../../../../jalapeno-gateway
	github.com/sbezverk/jalapeno-gateway/pkg/dbclient => ../../dbclient
	github.com/sbezverk/jalapeno-gateway/pkg/srvclient => ../../srvclient
)

require (
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
	github.com/sbezverk/jalapeno-gateway/pkg/dbclient v0.0.0-20200424215720-986760d824da
	github.com/sbezverk/jalapeno-gateway/pkg/srvclient v0.0.0-20200424215720-986760d824da // indirect
)
