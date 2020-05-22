package dbmockclient

import (
	"testing"
)

func TestNewMockDBClient(t *testing.T) {
	tests := []struct {
		name         string
		testDataFile string
		fail         bool
	}{
		{
			name:         "default file name",
			testDataFile: "",
			fail:         false,
		},
		{
			name:         "default file name",
			testDataFile: "non-existing.json",
			fail:         true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dbi := NewMockDBClient(true, tt.testDataFile)
			if dbi == nil {
				if !tt.fail {
					t.Error("supposed to succeed but failed")
				}
			}
			if dbi != nil {
				if tt.fail {
					t.Error("supposed to fail but succeeded")
				}
			}
		})
	}
}

func TestFilterByPrefix(t *testing.T) {
	tests := []struct {
		name    string
		records []MPLSL3Record
		prefix  string
		mask    uint32
		// expected defines a number of record left after filtering
		expected int
	}{
		{
			name: "Match found",
			records: []MPLSL3Record{
				{
					Prefix: "1.1.1.0",
					Mask:   24,
				},
				{
					Prefix: "1.1.2.0",
					Mask:   25,
				},
				{
					Prefix: "1.1.3.0",
					Mask:   29,
				},
			},
			prefix:   "1.1.1.0",
			mask:     24,
			expected: 1,
		},
		{
			name: "Match not found",
			records: []MPLSL3Record{
				{
					Prefix: "1.1.1.0",
					Mask:   24,
				},
				{
					Prefix: "1.1.2.0",
					Mask:   25,
				},
				{
					Prefix: "1.1.3.0",
					Mask:   29,
				},
			},
			prefix:   "1.1.4.0",
			mask:     24,
			expected: 0,
		},
		{
			name: "Match found but mask does not match",
			records: []MPLSL3Record{
				{
					Prefix: "1.1.1.0",
					Mask:   24,
				},
				{
					Prefix: "1.1.2.0",
					Mask:   25,
				},
				{
					Prefix: "1.1.3.0",
					Mask:   29,
				},
			},
			prefix:   "1.1.2.0",
			mask:     24,
			expected: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := filterByPrefix(tt.prefix, tt.mask, tt.records)
			if len(f) != tt.expected {
				t.Errorf("expected records length %d, got %d", tt.expected, len(f))
			}
		})
	}
}

func TestFilterByRT(t *testing.T) {
	tests := []struct {
		name    string
		records []MPLSL3Record
		rts     []string
		// expected defines a number of record left after filtering
		expected int
	}{
		{
			name: "Match of 1 RT",
			records: []MPLSL3Record{
				{
					RT: "100:200",
				},
				{
					RT: "100:300",
				},
				{
					RT: "100:400",
				},
			},
			rts:      []string{"100:200"},
			expected: 1,
		},
		{
			name: "Match of 2 RT",
			records: []MPLSL3Record{
				{
					RT: "100:200",
				},
				{
					RT: "100:300,100:400",
				},
				{
					RT: "100:400,100:300",
				},
			},
			rts:      []string{"100:300", "100:400"},
			expected: 2,
		},
		{
			name: "No Match of 2 RT",
			records: []MPLSL3Record{
				{
					RT: "100:200",
				},
				{
					RT: "100:300,100:400",
				},
				{
					RT: "100:400,100:300",
				},
			},
			rts:      []string{"100:300", "100:500"},
			expected: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := filterByRT(tt.rts, tt.records)
			if len(f) != tt.expected {
				t.Errorf("expected records length %d, got %d", tt.expected, len(f))
			}
		})
	}
}

func TestFilterByIPFamily(t *testing.T) {
	tests := []struct {
		name    string
		records []MPLSL3Record
		// expected defines a number of record left after filtering
		ipv4     bool
		expected int
	}{
		{
			name: "IPv4 only",
			records: []MPLSL3Record{
				{
					IPv4: true,
				},
				{
					IPv4: true,
				},
				{
					IPv4: true,
				},
			},
			ipv4:     true,
			expected: 3,
		},
		{
			name: "IPv6 only",
			records: []MPLSL3Record{
				{
					IPv4: false,
				},
				{
					IPv4: false,
				},
				{
					IPv4: false,
				},
			},
			ipv4:     false,
			expected: 3,
		},
		{
			name: "IPv4 only request for IPv6",
			records: []MPLSL3Record{
				{
					IPv4: true,
				},
				{
					IPv4: true,
				},
				{
					IPv4: true,
				},
			},
			ipv4:     false,
			expected: 0,
		},
		{
			name: "IPv6 only request for IPv4",
			records: []MPLSL3Record{
				{
					IPv4: true,
				},
				{
					IPv4: true,
				},
				{
					IPv4: true,
				},
			},
			ipv4:     false,
			expected: 0,
		},
		{
			name: "Mix IPv4 and IPv6 request for IPv4",
			records: []MPLSL3Record{
				{
					IPv4: true,
				},
				{
					IPv4: false,
				},
				{
					IPv4: true,
				},
				{
					IPv4: false,
				},
				{
					IPv4: true,
				},
			},
			ipv4:     true,
			expected: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := filterByIPFamily(tt.ipv4, tt.records)
			if len(f) != tt.expected {
				t.Errorf("expected records length %d, got %d", tt.expected, len(f))
			}
		})
	}
}
