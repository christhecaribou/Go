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
	got := isPalindrome(testcase)
	want := true

	fmt.Printf("Attempting %q ...\n", testcase)
	if got != want {
		t.Errorf("Got %v, wanted %v\nFAILED!\n", got, want)
	} else {
		fmt.Printf("Got %v, wanted %v\nPASSED!\n\n", got, want)
	}
}

func TestIsPalindrome2(t *testing.T) {
	testcase := "Race a car"
	got := isPalindrome(testcase)
	want := false

	fmt.Printf("Attempting %q ...\n", testcase)
	if got != want {
		t.Errorf("Got %v, wanted %v\nFAILED!\n", got, want)
	} else {
		fmt.Printf("Got %v, wanted %v\nPASSED!\n\n", got, want)
	}
}
func TestIsPalindrome3(t *testing.T) {
	testcase := "0P"
	got := isPalindrome(testcase)
	want := false

	fmt.Printf("Attempting %q ...\n", testcase)
	if got != want {
		t.Errorf("Got %v, wanted %v\nFAILED!\n", got, want)
	} else {
		fmt.Printf("Got %v, wanted %v\nPASSED!\n\n", got, want)
	}
}
func TestIsPalindrome4(t *testing.T) {
	testcase := "P"
	got := isPalindrome(testcase)
	want := true
	
	fmt.Printf("Attempting %q ...\n", testcase)
	if got != want {
		t.Errorf("Got %v, wanted %v\nFAILED!\n", got, want)
	} else {
		fmt.Printf("Got %v, wanted %v\nPASSED!\n\n", got, want)
	}
}