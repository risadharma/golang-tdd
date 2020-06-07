package main

import "testing"

func assertError(t *testing.T, got, want interface{}) {
	t.Helper()
	if got != want {
		t.Errorf("got %.2f, want %.2f", got, want)
	}
}

func checkArea(t *testing.T, shape Shape, want float64) {
	t.Helper()
	got := shape.Area()
	if got != want {
		t.Errorf("%#v got %.2f, want %.2f", shape, got, want)
	}
}

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0
	assertError(t, got, want)
}

func TestArea(t *testing.T) {
	areaTest := map[string]struct {
		shape   Shape
		hasArea float64
	}{
		"rectangle": {
			shape:   Rectangle{4, 6},
			hasArea: 24.0,
		},
		"circle": {
			shape:   Circle{10},
			hasArea: 314.1592653589793,
		},
		"triangle": {
			shape:   Triangle{12, 6},
			hasArea: 36.0,
		},
	}

	for name, tt := range areaTest {
		t.Run(name, func(*testing.T) {
			checkArea(t, tt.shape, tt.hasArea)
		})
	}

}
