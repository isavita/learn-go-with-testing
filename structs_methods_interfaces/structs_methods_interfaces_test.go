package structs_methods_interfaces

import (
	"math"
	"testing"
)

func TestPerimeter(t *testing.T) {
	checkPerimeter := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		got := shape.Perimeter()
		eps := 0.001
		if math.Abs(got-want) > eps {
			t.Errorf("got %.5f want %.5f", got, want)
		}
	}

	t.Run("reactangles", func(t *testing.T) {
		r := Reactangle{2.0, 4.5}
		want := 13.0
		checkPerimeter(t, r, want)
	})

	t.Run("circles", func(t *testing.T) {
		c := Circle{2.0}
		want := 12.5664
		checkPerimeter(t, c, want)
	})
}

func TestArea(t *testing.T) {
	checkArea := func(t *testing.T, name string, shape Shape, want float64) {
		t.Helper()
		t.Run(name, func(t *testing.T) {
			got := shape.Area()
			eps := 0.01
			if math.Abs(got-want) > eps {
				t.Errorf("%#v got %g want %g", shape, got, want)
			}
		})
	}

	areaTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{name: "Reactangle", shape: Reactangle{2.0, 4.5}, want: 9.0},
		{name: "Circle", shape: Circle{10}, want: 314.15},
		{name: "Triangle", shape: Triangle{12, 6}, want: 36.0},
	}

	for _, tt := range areaTests {
		checkArea(t, tt.name, tt.shape, tt.want)
	}
}

func BenchmarkPerimeter(b *testing.B) {
	r := Reactangle{10.0, 20.0}
	for i := 0; i < b.N; i++ {
		r.Perimeter()
	}
}

func BenchmarkArea(b *testing.B) {
	r := Reactangle{10.0, 20.0}
	for i := 0; i < b.N; i++ {
		r.Area()
	}
}
