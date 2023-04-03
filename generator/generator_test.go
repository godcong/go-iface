package generator

import (
	"testing"
)

func TestNewGenerator(t *testing.T) {
	tests := []struct {
		name string
		want *Generator
	}{
		{
			name: "",
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewGenerator()
			_, err := got.GenerateFromPath("../test")
			if err != nil {
				t.Fatal(err)
				return
			}

		})
	}
}
