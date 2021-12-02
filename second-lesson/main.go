package main

import (
	"fmt"
	calc "go/lesson2/Calc"

	"gonum.org/v1/gonum/mat"
)

func main() {
	p1 := calc.PointStruct{9, 8, 3}
	p2 := calc.PointStruct{14, 8, 3}
	p4 := calc.PointStruct{7, 12, 3}

	h := calc.AddPointToMatrix(p1, p2, p4)
	fa := mat.Formatted(h, mat.Prefix(" "), mat.Squeeze())
	fmt.Printf(" % v", fa)
	fmt.Println()

	calc.ParCalculation(p1, p2, p4)

	p5 := calc.PointStruct{7, 6, 3}
	p6 := calc.PointStruct{20, 15, 3}
	p7 := calc.PointStruct{3, 13, 3}

	d := calc.AddPointToMatrix(p5, p6, p7)

	dd := mat.Formatted(d, mat.Prefix(" "), mat.Squeeze())
	fmt.Printf(" % v", dd)
	fmt.Println()

	calc.ParCalculation(p5, p6, p7)

}
