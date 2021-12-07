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
		{"prints correctly", args{depths: []int{6, 7, 8, 6}}, "A: 6 (N/A - no previous sum)\nB: 7 (increased)\nC: 8 (increased)\nD: 6 (decreased)\n"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOutput := depthsToString(tt.args.depths); gotOutput != tt.wantOutput {
				t.Errorf("depthsToString() = %v, want %v", gotOutput, tt.wantOutput)
			}
		})
	}
}

func Test_createDepthWindows(t *testing.T) {
	type args struct {
		depths []int
	}
	tests := []struct {
		name        string
		args        args
		wantWindows [][3]int
	}{
		{"creates 4 windows correctly", args{depths: []int{1, 2, 3, 4, 5, 6, 7}}, [][3]int{{1, 2, 3}, {2, 3, 4}, {3, 4, 5}, {4, 5, 6}, {5, 6, 7}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotWindows := createDepthWindows(tt.args.depths); !reflect.DeepEqual(gotWindows, tt.wantWindows) {
				t.Errorf("createDepthWindows() = %v, want %v", gotWindows, tt.wantWindows)
			}
		})
	}
}

func Test_sumDepthWindows(t *testing.T) {
	type args struct {
		windows [][3]int
	}
	tests := []struct {
		name     string
		args     args
		wantSums []int
	}{
		{"Sums window correctly", args{windows: [][3]int{{1, 2, 3}, {2, 3, 4}, {3, 4, 5}}}, []int{6, 9, 12}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotSums := sumDepthWindows(tt.args.windows); !reflect.DeepEqual(gotSums, tt.wantSums) {
				t.Errorf("sumDepthWindows() = %v, want %v", gotSums, tt.wantSums)
			}
		})
	}
}

func TestReportDepths(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name       string
		args       args
		wantOutput string
		wantErr    bool
	}{
		{"Errors with negative depths", args{filename: "error.test.txt"}, "", true},
		{"Measures depths correctly", args{filename: "input.test.txt"}, "A: 607 (N/A - no previous sum)\nB: 618 (increased)\nC: 618 (no change)\nD: 617 (decreased)\nE: 647 (increased)\nF: 716 (increased)\nG: 769 (increased)\nH: 792 (increased)\n\n5 depth increases\n", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotOutput, err := ReportDepths(tt.args.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReportDepths() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotOutput != tt.wantOutput {
				t.Errorf("ReportDepths() gotOutput = %v, want %v", gotOutput, tt.wantOutput)
			}
		})
	}
}
