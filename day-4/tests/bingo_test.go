package tests

import (
	"github.com/braydend/advent-of-code/day-4/bingo"
	"testing"
)

var (
	testBoard    = bingo.Bingo{{14, 21, 17, 24, 4}, {10, 16, 15, 9, 19}, {18, 8, 23, 26, 20}, {22, 11, 13, 6, 5}, {2, 0, 12, 3, 7}}
	completeRow  = []int{14, 21, 17, 24, 4}
	completeCol  = []int{14, 10, 18, 22, 2}
	drawnNumbers = []int{7, 4, 9, 5, 11, 17, 23, 2, 0, 14, 21, 24}
)

func TestCheckBoard(t *testing.T) {
	type args struct {
		board   bingo.Bingo
		numbers []int
	}
	tests := []struct {
		name           string
		args           args
		wantIsComplete bool
	}{
		{"Board is not complete", args{testBoard, []int{1, 7, 3, 24, 5}}, false},
		{"Board has complete row", args{testBoard, completeRow}, true},
		{"Board has complete col", args{testBoard, completeCol}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotIsComplete := bingo.CheckBoard(tt.args.board, tt.args.numbers); gotIsComplete != tt.wantIsComplete {
				t.Errorf("CheckBoard() = %v, want %v", gotIsComplete, tt.wantIsComplete)
			}
		})
	}
}

func TestCalculateScore(t *testing.T) {
	type args struct {
		board   bingo.Bingo
		numbers []int
	}
	tests := []struct {
		name      string
		args      args
		wantScore int
	}{
		{"Correctly calculates score", args{testBoard, drawnNumbers}, 4512},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotScore := bingo.CalculateScore(tt.args.board, tt.args.numbers); gotScore != tt.wantScore {
				t.Errorf("CalculateScore() = %v, want %v", gotScore, tt.wantScore)
			}
		})
	}
}
