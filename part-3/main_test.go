package main

import "testing"

func TestCalculateGammaRate(t *testing.T) {
	type args struct {
		input [][5]bool
	}
	tests := []struct {
		name       string
		args       args
		wantResult uint
	}{
		{"Calculate gamma rate correctly", args{[][5]bool{
			{false, false, true, false, false},
			{true, true, true, true, false},
			{true, false, true, true, false},
			{true, false, true, true, true},
			{true, false, true, false, true},
			{false, true, true, true, true},
			{false, false, true, true, true},
			{true, true, true, false, false},
			{true, false, false, false, false},
			{true, true, false, false, true},
			{false, false, false, true, false},
			{false, true, false, true, false},
		}}, 22},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := CalculateGammaRate(tt.args.input); gotResult != tt.wantResult {
				t.Errorf("CalculateGammaRate() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
