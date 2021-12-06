package main

import (
	"reflect"
	"testing"
)

func Test_countDepthIncreases(t *testing.T) {
	type args struct {
		depths []int
	}
	tests := []struct {
		name          string
		args          args
		wantIncreases int
	}{
		{"counts increases correctly", args{depths: []int{1, 2, 3, 4, 1, 1, 1, 1}}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotIncreases := countDepthIncreases(tt.args.depths); gotIncreases != tt.wantIncreases {
				t.Errorf("countDepthIncreases() = %v, want %v", gotIncreases, tt.wantIncreases)
			}
		})
	}
}

func Test_parseDepths(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name       string
		args       args
		wantDepths []int
		wantErr    bool
	}{
		{"Parses positive values correctly", args{"1\n2\n3\n4\n5\n6"}, []int{1, 2, 3, 4, 5, 6}, false},
		{"Returns error with negative value", args{"-1\n2\n3\n4\n5\n6"}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotDepths, err := parseDepths(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseDepths() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotDepths, tt.wantDepths) {
				t.Errorf("parseDepths() gotDepths = %v, want %v", gotDepths, tt.wantDepths)
			}
		})
	}
}

func Test_depthsToString(t *testing.T) {
	type args struct {
		depths []int
	}
	tests := []struct {
		name       string
		args       args
		wantOutput string
	}{
		{"prints correctly", args{depths: []int{1, 2, 3, 2, 3}}, "1 (N/A - no previous measurement)\n2 (increased)\n3 (increased)\n2 (decreased)\n3 (increased)\n"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOutput := depthsToString(tt.args.depths); gotOutput != tt.wantOutput {
				t.Errorf("depthsToString() = %v, want %v", gotOutput, tt.wantOutput)
			}
		})
	}
}
