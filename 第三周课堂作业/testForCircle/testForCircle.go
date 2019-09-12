package main

import (
	"fmt"

	"github.com/github-user/circle"
)

func main() {
	var r float64
	fmt.Printf("Please input a radius:")
	fmt.Scan(&r)
	fmt.Printf("The area of circle with a radius of %f is %f\n", r, circle.GetAreaOfCircle(r))
}
