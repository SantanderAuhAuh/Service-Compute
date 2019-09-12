package circle

import (
	"math"
)

//GetAreaOfCircle compute the area of a circle
func GetAreaOfCircle(r float64) float64 {
	return math.Pi * math.Pow(r, 2)
}

//GetPerimeterOfCircle compute the perimeter of a circle
func GetPerimeterOfCircle(r float64) float64 {
	return 2 * math.Pi * r
}
