package calc

import (
	"fmt"
	"math"
)

type CircleStruct struct {
	Center PointStruct
	Length float64
	Radius float64
	Area   float64
}

func (c *CircleStruct) CalcRadius(s float64) {

	var par ParallelogramStruct
	a := par.Area
	c.Radius = math.Sqrt(a / math.Pi)
	c.Length = 2 * (math.Pi) * c.Radius
	fmt.Printf("Radius = %v and lengh = %v", c.Radius, c.Length)
}
