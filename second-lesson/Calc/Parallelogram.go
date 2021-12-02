package calc

import (
	"fmt"
	"math"

	"github.com/fatih/color"
	"gonum.org/v1/gonum/mat"
)

type PointStruct struct {
	X, Y  int
	Color int
}

func AddPoint(x, y int) PointStruct {
	var p PointStruct
	p.X = x
	p.Y = y
	p.Color = (x - y) / 2
	return p
}

type MatrixStruct struct {
	Matrix *mat.Dense
}

func CreateMatrix(size int) *mat.Dense {
	zero := mat.NewDense(size, size, nil)
	return zero
}

type ParallelogramStruct struct {
	Area           float64
	P1, P2, P3, P4 PointStruct
	Center         PointStruct
}

func FindCenterPar(p1, p2, p4 PointStruct) PointStruct {
	var center PointStruct
	center.X = (p2.X + p4.X) / 2
	center.Y = (p2.Y + p4.Y) / 2
	return center
}

func FindFour(p1, p2, p4 PointStruct) PointStruct {
	var p3 PointStruct
	center := FindCenterPar(p1, p2, p4)

	p3.X = 0
	p3.Y = 0
	for i := 0; i < 100000; i++ {
		if center.X != (p1.X+p3.X)/2 {
			p3.X++
		}
	}
	for i := 0; i < 100000; i++ {
		if center.Y != (p1.Y+p3.Y)/2 {
			p3.Y++
		}
	}
	p3.X = 2*center.X - p1.X
	p3.Y = 2*center.Y - p1.Y
	if p3.X < 0 || p3.Y < 0 {
		color.Red("matrix pos must be more then 0")
	}
	return p3

}
func AddPointToMatrix(p1, p2, p4 PointStruct) *mat.Dense {
	var p3 PointStruct
	var center PointStruct
	center = FindCenterPar(p1, p2, p4)
	p3 = FindFour(p1, p2, p4)
	arr := []int{p1.X, p1.Y, p2.Y, p2.X, p4.X, p4.Y, p3.X, p3.Y}
	max := 0
	for _, value := range arr {
		if max < value {
			max = value
		}
	}

	matrix := CreateMatrix(max + 1)
	matrix.Set(p1.Y, p1.X, 1)
	matrix.Set(p2.Y, p2.X, 1)
	matrix.Set(p4.Y, p4.X, 1)
	matrix.Set(p3.Y, p3.X, 1)
	matrix.Set(center.Y, center.X, 1)

	return matrix
}

func FindArea(p1, p2, p3 PointStruct) float64 {
	area := math.Abs(float64(p1.X*(p2.Y-p3.Y) - p1.Y*(p2.X-p3.X) + p2.X*p3.Y - p2.Y*p3.X))

	return area
}

func ParCalculation(p1, p2, p4 PointStruct) ParallelogramStruct {
	var par ParallelogramStruct
	var x1, x2, x3, x4 PointStruct
	par.Area = FindArea(p1, p2, p4)
	fmt.Printf("Parallelogram area : %v\n", par.Area)
	par.Center = FindCenterPar(p1, p2, p4)
	fmt.Printf("Parallelogram center : %v\n", par.Center)
	x3 = FindFour(p1, p2, p4)
	x1.Color = (p1.X - p1.Y) / 2
	x2.Color = (p2.X - p2.Y) / 2
	x3.Color = (x3.X - x3.Y) / 2
	x4.Color = (p4.X - p4.Y) / 2
	ChangeColor(int(math.Abs(float64(x1.Color))))
	fmt.Printf("First point : {%v : %v}\n", p1.X, p1.Y)
	color.Unset()
	ChangeColor(int(math.Abs(float64(x2.Color))))
	fmt.Printf("Second point : {%v : %v}\n", p2.X, p2.Y)
	color.Unset()
	ChangeColor(int(math.Abs(float64(x3.Color))))
	fmt.Printf("Third point : {%v : %v}\n", x3.X, x3.Y)
	color.Unset()
	ChangeColor(int(math.Abs(float64(x4.Color))))
	fmt.Printf("Fourth point : {%v : %v}\n", p4.X, p4.Y)
	color.Unset()
	return par

}
