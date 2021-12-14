package tests

import (
	"github.com/braydend/advent-of-code/day-4/bingo"
	"testing"
)

var (
	testBoard   = bingo.Bingo{{1, 2, 3, 4, 5}, {6, 7, 8, 9, 0}, {10, 11, 12, 13, 14}, {15, 16, 17, 18, 19}, {20, 21, 22, 23, 24}}
	completeRow = []int{1, 2, 3, 4, 5}
	completeCol = []int{1, 6, 10, 15, 20}
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
