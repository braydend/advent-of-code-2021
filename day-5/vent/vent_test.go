package vent

import (
	"reflect"
	"testing"
)

func TestVent_GetCoveredCoordinates(t *testing.T) {
	type fields struct {
		startingPosition Coordinates
		endingPosition   Coordinates
	}
	tests := []struct {
		name   string
		fields fields
		want   []Coordinates
	}{
		{"Correctly finds covered Coordinates across same X-axis", fields{Coordinates{1, 1}, Coordinates{1, 3}}, []Coordinates{{1, 1}, {1, 2}, {1, 3}}},
		{"Correctly finds covered Coordinates across same Y-axis", fields{Coordinates{9, 7}, Coordinates{7, 7}}, []Coordinates{{7, 7}, {8, 7}, {9, 7}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := Vent{
				StartingPosition: tt.fields.startingPosition,
				EndingPosition:   tt.fields.endingPosition,
			}
			if got := v.GetCoveredCoordinates(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCoveredCoordinates() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVent_IsHorizontal(t *testing.T) {
	type fields struct {
		startingPosition Coordinates
		endingPosition   Coordinates
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{"Correctly marks vent as horizontal with matching X-axis", fields{Coordinates{1, 1}, Coordinates{1, 100}}, true},
		{"Correctly marks vent as horizontal with matching Y-axis", fields{Coordinates{1, 1}, Coordinates{100, 1}}, true},
		{"Correctly marks vent as not horizontal with differing X and Y-axis", fields{Coordinates{1, 1}, Coordinates{100, 100}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := Vent{
				StartingPosition: tt.fields.startingPosition,
				EndingPosition:   tt.fields.endingPosition,
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
	expectedHorizontalOutput := map[Coordinates]int{
		Coordinates{7, 0}: 1,
		Coordinates{2, 1}: 1,
		Coordinates{7, 1}: 1,
		Coordinates{2, 2}: 1,
		Coordinates{7, 2}: 1,
		Coordinates{7, 3}: 1,
		Coordinates{1, 4}: 1,
		Coordinates{2, 4}: 1,
		Coordinates{3, 4}: 2,
		Coordinates{4, 4}: 1,
		Coordinates{5, 4}: 1,
		Coordinates{6, 4}: 1,
		Coordinates{7, 4}: 2,
		Coordinates{8, 4}: 1,
		Coordinates{9, 4}: 1,
		Coordinates{0, 9}: 2,
		Coordinates{1, 9}: 2,
		Coordinates{2, 9}: 2,
		Coordinates{3, 9}: 1,
		Coordinates{4, 9}: 1,
		Coordinates{5, 9}: 1,
	}
	expectedDiagonalOutput := map[Coordinates]int{
		Coordinates{0, 0}: 1,
		Coordinates{2, 0}: 1,
		Coordinates{7, 0}: 1,
		Coordinates{8, 0}: 1,
		Coordinates{1, 1}: 1,
		Coordinates{2, 1}: 1,
		Coordinates{3, 1}: 1,
		Coordinates{7, 1}: 2,
		Coordinates{2, 2}: 2,
		Coordinates{4, 2}: 1,
		Coordinates{6, 2}: 1,
		Coordinates{7, 2}: 1,
		Coordinates{8, 2}: 1,
		Coordinates{4, 3}: 2,
		Coordinates{6, 3}: 2,
		Coordinates{8, 3}: 1,
		Coordinates{2, 4}: 1,
		Coordinates{3, 4}: 2,
		Coordinates{4, 4}: 3,
		Coordinates{5, 4}: 1,
		Coordinates{6, 4}: 3,
		Coordinates{7, 4}: 2,
		Coordinates{8, 4}: 1,
		Coordinates{9, 4}: 1,
		Coordinates{4, 5}: 1,
		Coordinates{6, 5}: 2,
		Coordinates{6, 6}: 1,
		Coordinates{2, 6}: 1,
		Coordinates{6, 6}: 1,
		Coordinates{1, 7}: 1,
		Coordinates{7, 7}: 1,
		Coordinates{0, 8}: 1,
		Coordinates{8, 8}: 1,
		Coordinates{0, 9}: 2,
		Coordinates{1, 9}: 2,
		Coordinates{2, 9}: 2,
		Coordinates{3, 9}: 1,
		Coordinates{4, 9}: 1,
		Coordinates{5, 9}: 1,
	}
	type args struct {
		vents          []Vent
		onlyHorizontal bool
	}
	tests := []struct {
		name              string
		args              args
		wantVentLocations map[Coordinates]int
	}{
		{"correctly finds overlapping horizontal vents", args{inputVents, true}, expectedHorizontalOutput},
		{"correctly finds overlapping diagonal vents", args{inputVents, false}, expectedDiagonalOutput},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotVentLocations := FindOverlappingVents(tt.args.vents, tt.args.onlyHorizontal); !reflect.DeepEqual(gotVentLocations, tt.wantVentLocations) {
				t.Errorf("FindOverlappingVents() = %v, want %v", BuildVentLocationMap(gotVentLocations), BuildVentLocationMap(tt.wantVentLocations))
			}
		})
	}
}
