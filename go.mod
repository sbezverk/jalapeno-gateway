module github.com/sbezverk/jalapeno-gateway

go 1.13

require (
	github.com/arangodb/go-driver v0.0.0-20201202080739-c41c94f2de00
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
	github.com/golang/protobuf v1.4.0
	github.com/osrg/gobgp v0.0.0-20200513035851-833188f52610
	github.com/sbezverk/gobmp v0.0.1-beta.0.20200921121857-d571e95382ce
	golang.org/x/sys v0.0.0-20200420163511-1957bb5e6d1f // indirect
	google.golang.org/genproto v0.0.0-20200424135956-bca184e23272 // indirect
	google.golang.org/grpc v1.29.1
)

replace (
	github.com/osrg/gobgp => ../gobgp
	github.com/sbezverk/gobmp => ../gobmp
)
