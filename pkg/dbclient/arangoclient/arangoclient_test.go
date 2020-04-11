package arangoclient

import (
	"reflect"
	"strings"
	"testing"
)

func TestBuildQuery(t *testing.T) {
	tests := []struct {
		name           string
		collection     string
		f              []filter
		expectQuery    string
		expectBindVars map[string]interface{}
	}{
		{
			name:           "empty filters",
			collection:     "L3VPN_FIB",
			expectQuery:    "for q in L3VPN_FIB return q",
			expectBindVars: nil,
		},
		{
			name:        "rd filter",
			collection:  "L3VPN_FIB",
			expectQuery: "for q in L3VPN_FIB filter q.RD == @rd return q",
			expectBindVars: map[string]interface{}{
				"rd": "100:100",
			},
			f: []filter{
				{
					key:   "rd",
					value: "100:100",
				},
			},
		},
		{
			name:        "rd and ipv4 filters",
			collection:  "L3VPN_FIB",
			expectQuery: "for q in L3VPN_FIB filter q.RD == @rd and q.IPv4 == @ipv4 return q",
			expectBindVars: map[string]interface{}{
				"rd":   "100:100",
				"ipv4": true,
			},
			f: []filter{
				{
					key:   "rd",
					value: "100:100",
				},
				{
					key:   "ipv4",
					value: true,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotQ, gotB := buildQuery(tt.collection, tt.f...)
			if strings.Compare(gotQ, tt.expectQuery) != 0 {
				t.Errorf("failed, as expected query %q does not match the actual one %q", tt.expectQuery, gotQ)
			}
			if !reflect.DeepEqual(&gotB, &tt.expectBindVars) {
				t.Errorf("failed, as expected bindVars %+v does not match the actual one %+v", tt.expectBindVars, gotB)
			}
		})
	}
}
