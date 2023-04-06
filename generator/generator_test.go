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

func Test_camelCase(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "",
			args: args{
				s: "camelCase",
			},
			want: "CamelCase",
		},
		{
			name: "",
			args: args{
				s: "_camelCase",
			},
			want: "CamelCase",
		},
		{
			name: "",
			args: args{
				s: "_camel_Case",
			},
			want: "CamelCase",
		},
		{
			name: "",
			args: args{
				s: "_camel_case",
			},
			want: "CamelCase",
		},
		{
			name: "",
			args: args{
				s: "_camel_2ase",
			},
			want: "Camel_2ase",
		},
		{
			name: "",
			args: args{
				s: "21321_camel_Case",
			},
			want: "CamelCase",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := camelCase(tt.args.s); got != tt.want {
				t.Errorf("camelCase() = %v, want %v", got, tt.want)
			}
		})
	}
}
