package main

import (
	"fmt"
	"strings"
)

func main() {

	CheckPalindrome("aabaa")
	var matrix = [3][4]int{
		{0, 1, 0, 0},
		{1, 0, 0, 1},
		{1, 1, 1, 1}}
	MatrixElementSum(matrix)
	imputString := "foo(bar(baz))blim"
	fmt.Println(CreateWord(imputString))
	DuplicateEncode("din")
	number := []int{1, -2, 4, -4}
	PositiveSum(number)
	ToJadenCase("most trees are blue")
	Divisors(500000)
	kConcatenationMaxSum([]int{-1, 1, 1}, 3)
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

func GetCount(word string) int {

	count := 0

	vowels := []string{"a", "e", "i", "o", "u"}

	for _, vowel := range vowels {
		count += strings.Count(word, vowel)
	}

	return count
}

func IsTriangle(a, b, c int) bool {
	/*  Expect(IsTriangle(5, 1, 2)).To(Equal(false))
	Expect(IsTriangle(1, 2, 5)).To(Equal(false))
	Expect(IsTriangle(2, 5, 1)).To(Equal(false))
	Expect(IsTriangle(4, 2, 3)).To(Equal(true))
	Expect(IsTriangle(5, 1, 5)).To(Equal(true)) */
	if a+b > c && a+c > b && b+c > a {
		return true
	}

	return false
}

func PositiveSum(numbers []int) int {
	result := 0
	for _, c := range numbers {
		if c > 0 {
			result += c
		}
	}
	return result
}

func Grow(arr []int) int {
	result := 1
	for _, c := range arr {
		result *= c
	}
	return result
}

func Feast(beast string, dish string) bool {
	return beast[0] == dish[0] && beast[len(beast)-1] == dish[len(dish)-1]
}

func ToJadenCase(str string) string {
	return strings.Title(str)
	/*  strSlice:= strings.Split(str, " ")

	result := []string{}
	for _, c := range strSlice {
	  result = append(result, strings.Title(w))
	}

	return strings.Join(result, " ") */
}

func Divisors(n int) int {
	if n < 2 {
		return 1
	}

	result := 2
	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			result++
		}
	}

	return result
}

func MakeNegative(x int) int {
	if x > 0 {
		return x * -1
	}
	return x
}

func Capitalize(st string) []string {
	str1 := ""
	str2 := ""
	for i, c := range st {
		if i%2 == 0 {
			str1 += strings.ToUpper(string(c))
			str2 += strings.ToLower(string(c))
		} else {
			str1 += strings.ToLower(string(c))
			str2 += strings.ToUpper(string(c))
		}
	}

	return []string{str1, str2}
}

func kConcatenationMaxSum(arr []int, k int) int {
	slice := []int{}

	for i := 0; i < k; i++ {
		slice = append(slice, arr...)
	}
	if len(slice) == 0 {
		return 0
	}

	currentMaxSum := slice[0]
	finalMaxSum := slice[0]
	for _, c := range slice[1:] {
		currentMaxSum = Max(c, currentMaxSum+c)

		finalMaxSum = Max(finalMaxSum, currentMaxSum)
	}
	return finalMaxSum

}

func Max(x, y int) int {
	if x >= y {
		return x
	}
	return y
}
