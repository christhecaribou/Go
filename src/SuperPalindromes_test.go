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

func TestSuperPalindromes1(t *testing.T) {
	got := superpalindromesInRange("4","1000")
	expected := 4

	if got != expected {
		t.Errorf("Got %v, expected %v\nFAILED!\n", got, expected)
	} else {
		fmt.Printf("Got %v, expected %v\nPASSED!\n\n", got, expected)
	}
}
func TestSuperPalindromes2(t *testing.T) {
	got := superpalindromesInRange("1","2")
	expected := 1

	if got != expected {
		t.Errorf("Got %v, expected %v\nFAILED!\n", got, expected)
	} else {
		fmt.Printf("Got %v, expected %v\nPASSED!\n\n", got, expected)
	}
}