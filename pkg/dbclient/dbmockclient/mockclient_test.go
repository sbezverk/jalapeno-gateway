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
			dbi := NewMockDBClient(tt.testDataFile)
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
		records []Record
		prefix  string
		mask    uint32
		// expected defines a number of record left after filtering
		expected int
	}{
		{
			name: "Match found",
			records: []Record{
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
			records: []Record{
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
			records: []Record{
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
