package main

import (
	"reflect"
	"testing"
)

func TestNavigate(t *testing.T) {
	type args struct {
		instructions []Instruction
	}
	tests := []struct {
		name string
		args args
		want Position
	}{
		{"Correctly navigates", args{instructions: []Instruction{{Forward, 5}, {Down, 5}, {Forward, 8}, {Up, 3}, {Down, 8}, {Forward, 2}}}, Position{15, 60}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Navigate(tt.args.instructions); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Navigate() = %v, want %v", got, tt.want)
			}
		})
	}
}
