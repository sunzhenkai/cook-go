package sample

import "testing"

func TestDecode(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"empty"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Decode()
		})
	}
}
