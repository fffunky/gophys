package vector

import (
	"testing"
)

func TestMakeVector(t *testing.T) {
	tests := []struct {
		xIn      float64
		yIn      float64
		expected *struct {
			x float64
			y float64
		}
	}{
		{
			1,
			1,
			&struct {
				x float64
				y float64
			}{x: 1, y: 1},
		},
		{
			11.390,
			10.1231,
			&struct {
				x float64
				y float64
			}{x: 11.390, y: 10.1231},
		},
		{
			-26,
			-31,
			&struct {
				x float64
				y float64
			}{x: -26, y: -31},
		},
	}

	for _, tt := range tests {
		actual := MakeVector2(tt.xIn, tt.yIn)

		if actual.X() != tt.expected.x {
			t.Errorf("Expected vector.x to be %.2f, got=%.2f",
				tt.expected.x, actual.X())
		}

		if actual.Y() != tt.expected.y {
			t.Errorf("Expected vector.y to be %.2f, got=%.2f",
				tt.expected.y, actual.Y())
		}
	}
}

func TestSetXY(t *testing.T) {
	tests := []struct {
		name     string
		xIn      float64
		yIn      float64
		expected *struct {
			x float64
			y float64
		}
	}{
		{
			"SetX(1) / SetY(0)",
			1,
			0,
			&struct {
				x float64
				y float64
			}{x: 1, y: 0},
		},
		{
			"SetX(132) / SetY(-41)",
			132,
			-41,
			&struct {
				x float64
				y float64
			}{x: 132, y: -41},
		},
		{
			"SetX(-10.421) / SetY(22.41)",
			-10.421,
			22.41,
			&struct {
				x float64
				y float64
			}{x: -10.421, y: 22.41},
		},
	}

	for _, tt := range tests {
		actual := MakeVector2(100, 100)
		actual.SetX(tt.xIn)
		actual.SetY(tt.yIn)

		if actual.X() != tt.expected.x {
			t.Errorf("Expected vector.x to be %.2f, got=%.2f",
				tt.expected.x, actual.X())
		}

		if actual.Y() != tt.expected.y {
			t.Errorf("Expected vector.y to be %.2f, got=%.2f",
				tt.expected.y, actual.Y())
		}
	}
}
