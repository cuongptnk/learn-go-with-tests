package structs

import "testing"

func TestPerimeter(t *testing.T) {
	checkPerimeter := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Perimeter()

		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	}

	t.Run("Rectangles", func(t *testing.T) {
		r := Rectangle{Width: 10.0, Height: 10.0}
		checkPerimeter(t, r, 40.0)
	})

	t.Run("Circles", func(t *testing.T) {
		c := Circle{Radius: 10.0}
		checkPerimeter(t, c, 62.83185307179586)
	})

}

func TestArea(t *testing.T) {
	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	}

	t.Run("rectangles", func(t *testing.T) {
		r := Rectangle{Width: 10.0, Height: 20.0}
		checkArea(t, r, 200.0)
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{10}
		checkArea(t, circle, 314.1592653589793)
	})

}
