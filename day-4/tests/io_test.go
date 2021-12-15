package tests

import (
	"github.com/braydend/advent-of-code/day-4/bingo"
	"github.com/braydend/advent-of-code/day-4/io"
	"log"
	"os"
	"reflect"
	"testing"
)

func readTestInput(filename string) []byte {
	data, err := os.ReadFile(filename)

	if err != nil {
		log.Fatal(err)
	}

	return data
}

func TestParseNumbers(t *testing.T) {
	type args struct {
		input []byte
	}
	tests := []struct {
		name        string
		args        args
		wantNumbers []int
		wantErr     bool
	}{
		{"Correctly parses numbers from test input", args{readTestInput("testinput.txt")}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotNumbers, err := io.ParseNumbers(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseNumbers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotNumbers, tt.wantNumbers) {
				t.Errorf("ParseNumbers() gotNumbers = %v, want %v", gotNumbers, tt.wantNumbers)
			}
		})
	}
}

func TestParseBoards(t *testing.T) {
	type args struct {
		input []byte
	}
	tests := []struct {
		name       string
		args       args
		wantBoards []bingo.Bingo
		wantErr    bool
	}{
		{
			"Correctly parses single board from a file",
			args{readTestInput("testinput.txt")},
			[]bingo.Bingo{bingo.Bingo{{1, 2, 3, 4, 5}, {6, 7, 8, 9, 10}, {11, 12, 13, 14, 15}, {16, 17, 18, 19, 20}, {21, 22, 23, 24, 25}}},
			false,
		},
		{
			"Correctly parses multiple boards from a file",
			args{readTestInput("testinput2.txt")},
			[]bingo.Bingo{
				bingo.Bingo{{1, 2, 3, 4, 5}, {6, 7, 8, 9, 10}, {11, 12, 13, 14, 15}, {16, 17, 18, 19, 20}, {21, 22, 23, 24, 25}},
				bingo.Bingo{{31, 2, 3, 4, 5}, {6, 37, 8, 9, 10}, {11, 12, 33, 14, 15}, {16, 17, 18, 39, 20}, {21, 22, 23, 24, 35}},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotBoards, err := io.ParseBoards(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseNumbers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotBoards, tt.wantBoards) {
				t.Errorf("ParseNumbers() gotBoards = %v, want %v", gotBoards, tt.wantBoards)
			}
		})
	}
}
