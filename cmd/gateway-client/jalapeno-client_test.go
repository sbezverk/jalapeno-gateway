package main

import (
	"testing"

	"github.com/sbezverk/jalapeno-gateway/pkg/bgpclient"
)

func TestValidateRD(t *testing.T) {
	tests := []struct {
		name string
		rd   string
		fail bool
	}{
		{
			name: "2 bytes asn : 4 bytes value",
			rd:   "577:12345",
			fail: false,
		},
		{
			name: "4 bytes asn : 2 bytes value",
			rd:   "57756565:12345",
			fail: false,
		},
		{
			name: "ip address asn : 2 bytes value",
			rd:   "5.7.7.1:12345",
			fail: false,
		},
		{
			name: "malformed 4 bytes asn : 2 bytes value",
			rd:   "57756565:1234598989",
			fail: true,
		},
		{
			name: "malformed ip address asn : 2 bytes value",
			rd:   "5.7.7.1.2:12345",
			fail: true,
		},
	}
	for _, tt := range tests {
		err := bgpclient.RDValidator(tt.rd)
		if err != nil && !tt.fail {
			t.Errorf("test \"%s\" failed as expected to succeed but fail with error: %+v", tt.name, err)
		}
		if err == nil && tt.fail {
			t.Errorf("test \"%s\" failed as expected to fail but succeeded", tt.name)
		}
	}
}
