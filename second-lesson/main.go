package main

import (
	"fmt"
	"math"
	"math/rand"

	"github.com/fatih/color"
	"github.com/fatih/structs"
)

func main() {

	p := PointStruct{}

	f, a := p.AddThrePoint(4, 4, 4, 15, 15, 25)
	y, x, c := p.FindLastCoordinate(f, a)
	p.DrawMatrix(y, x, c)
	// p.SquareArea(f, a)
	// fmt.Println(ChangeColor())
	// color.Unset()

}

type PointStruct struct {
	PointPositionX int
	PointPositionY int
}

type ColorStruct struct {
	White, Red, Yellow, Blue, Cyan, Green, Magneta, Black string
}

func (p PointStruct) AddThrePoint(p1X, p1Y, p2X, p2Y, p3X, p3Y int) (y, x []int) {

	p1 := PointStruct{}
	p2 := PointStruct{}
	p3 := PointStruct{}

	p1.PointPositionX = p1X
	p1.PointPositionY = p1Y
	p2.PointPositionX = p2X
	p2.PointPositionY = p2Y
	p3.PointPositionX = p3X
	p3.PointPositionY = p3Y

	x = make([]int, 0)
	x = append(x, p1X, p2X, p3X)
	y = make([]int, 0)
	y = append(y, p1Y, p2Y, p3Y)

	return y, x
}

func (p PointStruct) SquareArea(y, x []int) float64 {
	area := math.Abs(float64(x[0]*(y[1]-y[2]) - y[0]*(x[1]-x[2]) + x[1]*y[2] - y[1]*x[2]))
	fmt.Println("SquareArea = ", area)
	return area
}

func (p PointStruct) FindLastCoordinate(y, x []int) (j, f, c []int) {
	fmt.Println(x)
	fmt.Println(y)
	centerX := (x[0] + x[2]) / 2
	centerY := (y[0] + y[2]) / 2

	p4X := 2*centerX - x[1]
	p4Y := 2*centerY - y[1]
	fmt.Println(p4X, p4Y)
	x = append(x, p4X)
	y = append(y, p4Y)
	center := make([]int, 0)
	center = append(center, centerY, centerX)

	return y, x, center
}

func (p *PointStruct) DrawMatrix(y, x, c []int) {
	maxX := MaxValue(x)
	maxY := MaxValue(y)

	matrixSize := 0
	if maxX > maxY {
		matrixSize = maxX + 2
	} else {
		matrixSize = maxY + 2
	}

	a := make([][]string, matrixSize)
	for i := range a {
		a[i] = make([]string, matrixSize)
	}

	for i := 0; i < len(x); i++ {
		a[y[i]][x[i]] = "*"
	}
	fmt.Println(c)
	ChangeColor()
	a[c[0]][c[1]] = "*"
	color.Unset()

	for i := 0; i < matrixSize; i++ {

		for j := 0; j < matrixSize; j++ {

			if a[i][j] == "*" {
				ChangeColor()
				fmt.Print("*" + " ")
				color.Unset()
			} else {
				fmt.Print("-" + " ")
			}
		}
		fmt.Print("\n")
	}

}

func MaxValue(x []int) int {
	max := 0
	for _, value := range x {
		if value > max {
			max = value
		}
	}
	return max
}

func RefStruct(c ColorStruct) []interface{} {
	c = ColorStruct{"white", "red", "yellow", "blue", "cyan", "green", "magneta", "black"}
	val := structs.Values(c)
	return val
}
func ChangeColor() string {
	randInt := GenerateRandomInt()
	col := RefStruct(ColorStruct{})

	switch col[randInt] {
	case "white":
		color.Set(color.BgWhite)
		return "white"
	case "red":
		color.Set(color.BgRed)
		return "Red"
	case "yellow":
		color.Set(color.BgYellow)
		return "yellow"
	case "blue":
		color.Set(color.BgBlue)
		return "blue"
	case "cyan":
		color.Set(color.BgCyan)
		return "cyan"
	case "green":
		color.Set(color.BgGreen)
		return "green"
	case "magneta":
		color.Set(color.BgMagenta)
		return "magneta/pink"
	case "black":
		color.Set(color.BgBlack)
		return "black"
	default:
		fmt.Println("We dont this Color in sturct")
	}
	return ""
}

func GenerateRandomInt() int {
	min := 0
	max := 7
	return rand.Intn(max-min) + min
}
