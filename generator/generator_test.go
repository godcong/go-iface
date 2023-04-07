package generator

import (
	"fmt"
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
			got := New()
			ifaces, err := got.GenerateFromPath("../test")
			if err != nil {
				t.Fatal(err)
				return
			}
			for name, iface := range ifaces {
				fmt.Printf("file:%s,ifaces:%s", name, iface)
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

func Test_snakeCase(t *testing.T) {
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
				s: "SnakeCase",
			},
			want: "snake_case",
		},
		{
			name: "",
			args: args{
				s: "SnakeCASE",
			},
			want: "snake_case",
		},
		{
			name: "",
			args: args{
				s: "Snake_CASE",
			},
			want: "snake_case",
		},
		{
			name: "",
			args: args{
				s: "_Snake_CASE",
			},
			want: "snake_case",
		},
		{
			name: "",
			args: args{
				s: "_Snake_CaseJSON",
			},
			want: "snake_case_json",
		},
		{
			name: "",
			args: args{
				s: "___Snake_CaseJSON",
			},
			want: "snake_case_json",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := snakeCase(tt.args.s); got != tt.want {
				t.Errorf("snakeCase() = %v, want %v", got, tt.want)
			}
		})
	}
}
