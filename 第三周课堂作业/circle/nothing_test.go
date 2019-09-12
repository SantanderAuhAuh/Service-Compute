package circle

import (
	"math"
	"testing"
)

func TestCircle(t *testing.T) {
	cases := []float64{
		float64(1), float64(2), float64(5), float64(8), float64(0.25),
	}
	result1 := []float64{
		math.Pi * 1, math.Pi * 4, math.Pi * 25, math.Pi * 64, math.Pi * math.Pow(0.25, 2),
	}
	result2 := []float64{
		2 * math.Pi * 1, 2 * math.Pi * 2, 2 * math.Pi * 5, 2 * math.Pi * 8, 2 * math.Pi * 0.25,
	}
	for i, n := range cases {
		if GetAreaOfCircle(n) != result1[i] {
			t.Errorf("The area of the circle with radius of %f is not %f ", n, GetPerimeterOfCircle(n))
		}
		if GetPerimeterOfCircle(n) != result2[i] {
			t.Errorf("The perimeter of the with radius of %f is not %f", n, GetPerimeterOfCircle(n))
		}
	}
}
