package main

import "testing"

var testInput = [][]bool{
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
}

func TestCalculateGammaRate(t *testing.T) {
	type args struct {
		input [][]bool
	}
	tests := []struct {
		name       string
		args       args
		wantResult uint
	}{
		{"Calculate gamma rate correctly", args{testInput}, 22},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := CalculateGammaRate(tt.args.input); gotResult != tt.wantResult {
				t.Errorf("CalculateGammaRate() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func TestCalculateEpsilonRate(t *testing.T) {
	type args struct {
		input [][]bool
	}
	tests := []struct {
		name       string
		args       args
		wantResult uint
	}{
		{"Calculates sigma correctly", args{testInput}, 9},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := CalculateEpsilonRate(tt.args.input); gotResult != tt.wantResult {
				t.Errorf("CalculateEpsilonRate() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func TestCalculateOxygenGeneratorRating(t *testing.T) {
	type args struct {
		input [][]bool
	}
	tests := []struct {
		name       string
		args       args
		wantOutput uint
	}{
		{"Calculates oxygen generator rating correctly", args{testInput}, 23},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOutput := CalculateOxygenGeneratorRating(tt.args.input); gotOutput != tt.wantOutput {
				t.Errorf("CalculateOxygenGeneratorRating() = %v, want %v", gotOutput, tt.wantOutput)
			}
		})
	}
}
