package sample

import "testing"

func TestTA(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"empty"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			TA()
		})
	}
}
