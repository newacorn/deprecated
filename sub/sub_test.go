package sub

import "testing"

func TestInfo(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"first"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Info()
		})
	}
}
