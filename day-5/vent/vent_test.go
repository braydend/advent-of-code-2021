package vent

import (
	"testing"
)

func TestVent_IsHorizontal(t *testing.T) {
	type fields struct {
		startingPosition position
		endingPosition   position
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{"Correctly marks vent as horizontal with matching x-axis", fields{position{1, 1}, position{1, 100}}, true},
		{"Correctly marks vent as horizontal with matching y-axis", fields{position{1, 1}, position{100, 1}}, true},
		{"Correctly marks vent as not horizontal with differing x and y-axis", fields{position{1, 1}, position{100, 100}}, false},
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
