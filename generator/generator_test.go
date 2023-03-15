package generator

import (
	"reflect"
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
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewGenerator() = %v, want %v", got, tt.want)
			}
			_, err := got.GenerateFromPath("../test")
			if err != nil {
				return
			}

		})
	}
}
