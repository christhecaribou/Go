/*
CHRIS FELLI, 2019
Let's say a positive integer is a superpalindrome if it is a palindrome, 
and it is also the square of a palindrome.
*/
//package main ...
package main

import (
	"fmt"
	"testing"
)

func TestOptimalSuperPalindromes1(t *testing.T) {
	got := optimalSuperPalindromesInRange("4", "1000")
	expected := 4

	if got != expected {
		t.Errorf("Got %v, expected %v\nFAILED!\n", got, expected)
	} else {
		fmt.Printf("Got %v, expected %v\nPASSED!\n\n", got, expected)
	}
}

func TestOptimalSuperPalindromes2(t *testing.T) {
	got := optimalSuperPalindromesInRange("1", "999999999999999999")
	expected := 70

	if got != expected {
		t.Errorf("Got %v, expected %v\nFAILED!\n", got, expected)
	} else {
		fmt.Printf("Got %v, expected %v\nPASSED!\n\n", got, expected)
	}
}