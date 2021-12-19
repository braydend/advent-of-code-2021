package tests

import (
	"github.com/braydend/advent-of-code/day-4/bingo"
	"reflect"
	"testing"
)

var (
	testBoardOne   = bingo.Bingo{{22, 13, 17, 11, 0}, {8, 2, 23, 4, 24}, {21, 9, 14, 16, 7}, {6, 10, 3, 18, 5}, {1, 12, 20, 15, 19}}
	testBoardTwo   = bingo.Bingo{{3, 15, 0, 2, 22}, {9, 18, 13, 17, 5}, {19, 8, 7, 25, 23}, {20, 11, 10, 24, 4}, {14, 21, 16, 12, 6}}
	testBoardThree = bingo.Bingo{{14, 21, 17, 24, 4}, {10, 16, 15, 9, 19}, {18, 8, 23, 26, 20}, {22, 11, 13, 6, 5}, {2, 0, 12, 3, 7}}
	completeRow    = []int{18, 8, 23, 26, 20}
	completeCol    = []int{17, 15, 23, 13, 12}
	drawnNumbers   = []int{7, 4, 9, 5, 11, 17, 23, 2, 0, 14, 21, 24, 10, 16, 13, 6, 15, 25, 12, 22, 18, 20, 8, 19, 3, 26, 1}
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
		{"Board is not complete", args{testBoardThree, []int{1, 7, 3, 24, 5}}, false},
		{"Board has complete row", args{testBoardThree, completeRow}, true},
		{"Board has complete col", args{testBoardThree, completeCol}, true},
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
		{"Correctly calculates score", args{testBoardThree, []int{7, 4, 9, 5, 11, 17, 23, 2, 0, 14, 21, 24}}, 4512},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotScore := bingo.CalculateScore(tt.args.board, tt.args.numbers); gotScore != tt.wantScore {
				t.Errorf("CalculateScore() = %v, want %v", gotScore, tt.wantScore)
			}
		})
	}
}

func TestFindWinningBoards(t *testing.T) {
	type args struct {
		boards  []bingo.Bingo
		numbers []int
	}
	tests := []struct {
		name              string
		args              args
		wantWinners       []bingo.Bingo
		wantPlayerNumbers []int
		wantNumbersCalled []int
	}{
		{"Correctly finds the winning board", args{[]bingo.Bingo{testBoardOne, testBoardTwo, testBoardThree}, drawnNumbers}, []bingo.Bingo{testBoardThree}, []int{3}, []int{7, 4, 9, 5, 11, 17, 23, 2, 0, 14, 21, 24}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotWinners, gotPlayerNumbers, gotNumbersCalled := bingo.FindWinningBoards(tt.args.boards, tt.args.numbers)
			if !reflect.DeepEqual(gotWinners, tt.wantWinners) {
				t.Errorf("FindWinningBoards() gotWinners = %v, want %v", gotWinners, tt.wantWinners)
			}
			if !reflect.DeepEqual(gotPlayerNumbers, tt.wantPlayerNumbers) {
				t.Errorf("FindWinningBoards() gotPlayerNumbers = %v, want %v", gotPlayerNumbers, tt.wantPlayerNumbers)
			}
			if !reflect.DeepEqual(gotNumbersCalled, tt.wantNumbersCalled) {
				t.Errorf("FindWinningBoards() gotNumbersCalled = %v, want %v", gotNumbersCalled, tt.wantNumbersCalled)
			}
		})
	}
}

func TestFindLastWinningBoard(t *testing.T) {
	type args struct {
		boards  []bingo.Bingo
		numbers []int
	}
	tests := []struct {
		name              string
		args              args
		wantBoard         bingo.Bingo
		wantNumbersCalled []int
	}{
		{"Correctly find the last winning board", args{[]bingo.Bingo{testBoardOne, testBoardTwo, testBoardThree}, drawnNumbers}, testBoardTwo, []int{7, 4, 9, 5, 11, 17, 23, 2, 0, 14, 21, 24, 10, 16, 13}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotBoard, gotNumbersCalled := bingo.FindLastWinningBoard(tt.args.boards, tt.args.numbers)
			if !reflect.DeepEqual(gotBoard, tt.wantBoard) {
				t.Errorf("FindLastWinningBoard() = %v, want %v", gotBoard, tt.wantBoard)
			}
			if !reflect.DeepEqual(gotNumbersCalled, tt.wantNumbersCalled) {
				t.Errorf("FindWinningBoard() gotNumbersCalled = %v, want %v", gotNumbersCalled, tt.wantNumbersCalled)
			}
		})
	}
}
