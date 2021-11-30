package main

import (
	"go/lesson1/first-lesson/tasks"
	"testing"
)

func TestPalindrome(t *testing.T) {
	t.Run("positive", func(t *testing.T) {
		expectation := true
		actual := tasks.CheckPalindrome("aabaa")
		if actual != expectation {
			t.Errorf("Expected %t but got %t", expectation, actual)
		}
	})
	t.Run("negative", func(t *testing.T) {
		expectation := false
		actual := tasks.CheckPalindrome("abac")
		if actual != expectation {
			t.Errorf("Expected %t but got %t", expectation, actual)
		}
	})
	t.Run("withOneLetter", func(t *testing.T) {
		expectation := true
		actual := tasks.CheckPalindrome("a")

		if actual != expectation {
			t.Errorf("Expected %t but got %t", expectation, actual)
		}
	})
}

func TestMatrixElement(t *testing.T) {

	expectation := 9
	matrix := [3][4]int{{0, 1, 1, 2}, {0, 5, 0, 0}, {2, 0, 0, 3}}
	actual := tasks.MatrixElementSum(matrix)
	if actual != expectation {
		t.Errorf("Expected %d but got %d", expectation, actual)
	}
}

func TestCreateWord(t *testing.T) {
	t.Run("one", func(t *testing.T) {
		expectation := "rab"
		actual := tasks.CreateWord("(bar)")
		retWord := ""
		for _, c := range actual {
			if c != 0 {
				retWord += string(c)
			}
		}
		if expectation != retWord {
			t.Errorf("Expected %s but got %s", expectation, actual)
		}
	})
	t.Run("two", func(t *testing.T) {
		expectation := "foorabbaz"
		actual := tasks.CreateWord("foo(bar)baz")
		retWord := ""
		for _, c := range actual {
			if c != 0 {
				retWord += string(c)
			}
		}
		if expectation != retWord {
			t.Errorf("Expected %s but got %s", expectation, actual)
		}
	})
	t.Run("three", func(t *testing.T) {
		expectation := "foobazrabblim"
		actual := tasks.CreateWord("foo(bar(baz))blim")
		retWord := ""
		for _, c := range actual {
			if c != 0 {
				retWord += string(c)
			}
		}
		if expectation != retWord {
			t.Errorf("Expected %s but got %s", expectation, actual)
		}
	})
}
