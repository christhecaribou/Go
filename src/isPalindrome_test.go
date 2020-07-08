/*
CHRIS FELLI, 2019
Given a string, returns whether or not it's a palindrome
*/

//package main ...
package main

import (
	"fmt"
	"testing"
)

func TestIsPalindrome1(t *testing.T) {
	testcase := "A man, a plan, a canal: Panama"
	actual := isPalindrome(testcase)
	expected := true

	fmt.Printf("Attempting %q ...\n", testcase)
	if actual != expected {
		t.Errorf("actual %v, expected %v\nFAILED!\n", actual, expected)
	} else {
		fmt.Printf("actual %v, expected %v\nPASSED!\n\n", actual, expected)
	}
}

func TestIsPalindrome2(t *testing.T) {
	testcase := "Race a car"
	actual := isPalindrome(testcase)
	expected := false

	fmt.Printf("Attempting %q ...\n", testcase)
	if actual != expected {
		t.Errorf("actual %v, expected %v\nFAILED!\n", actual, expected)
	} else {
		fmt.Printf("actual %v, expected %v\nPASSED!\n\n", actual, expected)
	}
}
func TestIsPalindrome3(t *testing.T) {
	testcase := "0P"
	actual := isPalindrome(testcase)
	expected := false

	fmt.Printf("Attempting %q ...\n", testcase)
	if actual != expected {
		t.Errorf("actual %v, expected %v\nFAILED!\n", actual, expected)
	} else {
		fmt.Printf("actual %v, expected %v\nPASSED!\n\n", actual, expected)
	}
}
func TestIsPalindrome4(t *testing.T) {
	testcase := "P"
	actual := isPalindrome(testcase)
	expected := true

	fmt.Printf("Attempting %q ...\n", testcase)
	if actual != expected {
		t.Errorf("actual %v, expected %v\nFAILED!\n", actual, expected)
	} else {
		fmt.Printf("actual %v, expected %v\nPASSED!\n\n", actual, expected)
	}
}
