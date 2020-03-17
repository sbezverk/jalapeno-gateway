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
