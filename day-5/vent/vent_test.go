package vent

import (
	"reflect"
	"testing"
)

func TestVent_GetCoveredCoordinates(t *testing.T) {
	type fields struct {
		startingPosition coordinates
		endingPosition   coordinates
	}
	tests := []struct {
		name   string
		fields fields
		want   []coordinates
	}{
		{"Correctly finds covered coordinates across same x-axis", fields{coordinates{1, 1}, coordinates{1, 3}}, []coordinates{{1, 1}, {1, 2}, {1, 3}}},
		{"Correctly finds covered coordinates across same y-axis", fields{coordinates{9, 7}, coordinates{7, 7}}, []coordinates{{7, 7}, {8, 7}, {9, 7}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := Vent{
				startingPosition: tt.fields.startingPosition,
				endingPosition:   tt.fields.endingPosition,
			}
			if got := v.GetCoveredCoordinates(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCoveredCoordinates() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVent_IsHorizontal(t *testing.T) {
	type fields struct {
		startingPosition coordinates
		endingPosition   coordinates
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{"Correctly marks vent as horizontal with matching x-axis", fields{coordinates{1, 1}, coordinates{1, 100}}, true},
		{"Correctly marks vent as horizontal with matching y-axis", fields{coordinates{1, 1}, coordinates{100, 1}}, true},
		{"Correctly marks vent as not horizontal with differing x and y-axis", fields{coordinates{1, 1}, coordinates{100, 100}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := Vent{
				startingPosition: tt.fields.startingPosition,
				endingPosition:   tt.fields.endingPosition,
			}
			if got := v.IsHorizontal(); got != tt.want {
				t.Errorf("IsHorizontal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindOverlappingVents(t *testing.T) {
	inputVents := []Vent{
		NewVent(0, 9, 5, 9),
		NewVent(8, 0, 0, 8),
		NewVent(9, 4, 3, 4),
		NewVent(2, 2, 2, 1),
		NewVent(7, 0, 7, 4),
		NewVent(6, 4, 2, 0),
		NewVent(0, 9, 2, 9),
		NewVent(3, 4, 1, 4),
		NewVent(0, 0, 8, 8),
		NewVent(5, 5, 8, 2),
	}
	expectedOutput := map[coordinates]int{
		coordinates{7, 0}: 1,
		coordinates{2, 1}: 1,
		coordinates{7, 1}: 1,
		coordinates{2, 2}: 1,
		coordinates{7, 2}: 1,
		coordinates{7, 3}: 1,
		coordinates{1, 4}: 1,
		coordinates{2, 4}: 1,
		coordinates{3, 4}: 2,
		coordinates{4, 4}: 1,
		coordinates{5, 4}: 1,
		coordinates{6, 4}: 1,
		coordinates{7, 4}: 2,
		coordinates{8, 4}: 1,
		coordinates{9, 4}: 1,
		coordinates{0, 9}: 2,
		coordinates{1, 9}: 2,
		coordinates{2, 9}: 2,
		coordinates{3, 9}: 1,
		coordinates{4, 9}: 1,
		coordinates{5, 9}: 1,
	}
	type args struct {
		vents []Vent
	}
	tests := []struct {
		name              string
		args              args
		wantVentLocations map[coordinates]int
	}{
		{"correctly finds overlapping vents", args{inputVents}, expectedOutput},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotVentLocations := FindOverlappingVents(tt.args.vents); !reflect.DeepEqual(gotVentLocations, tt.wantVentLocations) {
				t.Errorf("FindOverlappingVents() = %v, want %v", gotVentLocations, tt.wantVentLocations)
			}
		})
	}
}
