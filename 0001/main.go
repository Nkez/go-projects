package main

import (
	"fmt"
	"strings"
)

func main() {
	word := "aabaa"
	AbbrevName(word)
	CheckPalindrome(word)
	var matrix = [3][4]int{
		{0, 1, 0, 0},
		{1, 0, 0, 1},
		{1, 1, 1, 1}}
	MatrixElementSum(matrix)
	imputString := "foo(bar(baz))blim" // "foo(bar)baz(blim)" // foo
	CreateWord(imputString)
	fmt.Println(DuplicateEncode("din"))
}

func CheckPalindrome(word string) bool {
	if len(word) < 2 {
		return true
	}

	for i := 0; i < len(word)/2; i++ {
		if word[i] != word[len(word)-i-1] {
			return false
		}
	}

	return true
}

func MatrixElementSum(matrix [3][4]int) int {
	sum := 0
	for i := 0; i < 1; i++ {
		for j := 0; j < 4; j++ {
			sum += matrix[i][j]
		}
	}
	for i := 1; i < 3; i++ {
		for j := 0; j < 4; j++ {
			if matrix[i-1][j] != 0 {
				sum += matrix[i][j]
			}
		}
	}
	return sum
}

func CreateWord(word string) string {

	array := make([]rune, len(word))
	emptyArray := make([]rune, len(word))

	for _, c := range word {
		if c == ')' {
			i := len(array) - 1
			emptyArray = emptyArray[:0]
			for ; array[i] != '('; i-- {
				emptyArray = append(emptyArray, array[i])
			}
			array = append(array[:i], emptyArray...)
		} else {
			array = append(array, c)
		}
	}
	return string(array)
}

func AbbrevName(name string) string {
	//your code here
	s := ""
	for i := 1; i < len(name); i++ {
		if string(name[i]) == " " {
			s = string(name[i+1])
		}
	}
	c := strings.ToUpper(string(name[0])) + "." + strings.ToUpper(s) + "."
	return c
}

func DuplicateEncode(word string) string {
	/* 	"din"      =>  "((("
	"recede"   =>  "()()()"
	"Success"  =>  ")())())"
	"(( @"     =>  "))((" */
	word = strings.ToLower(word)
	spl := strings.Split(word, "")
	fmt.Println(spl)
	result := ""
	for _, v := range spl {
		if strings.Count(word, v) > 1 {
			result += ")"
		} else {
			result += "("
		}
	}
	return result

}
