package main

import (
	"fmt"
	"go/lesson1/first-lesson/tasks"
)

func main() {

	pal := tasks.CheckPalindrome("abbba")
	fmt.Println(pal)

	matrix := [3][4]int{{0, 1, 1, 2}, {0, 5, 0, 0}, {2, 0, 0, 3}}
	fmt.Println(tasks.MatrixElementSum(matrix))

	fmt.Println(tasks.CreateWord("foo(bar(baz))blim"))
	fmt.Println(tasks.CreateWord("foo(bar)baz"))
	fmt.Println(tasks.CreateWord("(bar)"))

}
