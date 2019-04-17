/*
CHRIS FELLI, 2019
Let's say a positive integer is a superpalindrome if it is a palindrome, 
and it is also the square of a palindrome.
*/

//Package gocode ...
package gocode

import (
	"fmt"
	"testing"
)

func TestSuperPalindromes1(t *testing.T) {
	got := superpalindromesInRange("4","1000")
	want := 4

	if got != want {
		t.Errorf("Got %v, wanted %v\nFAILED!\n", got, want)
	} else {
		fmt.Printf("Got %v, wanted %v\nPASSED!\n\n", got, want)
	}
}
func TestSuperPalindromes2(t *testing.T) {
	got := superpalindromesInRange("1","2")
	want := 1

	if got != want {
		t.Errorf("Got %v, wanted %v\nFAILED!\n", got, want)
	} else {
		fmt.Printf("Got %v, wanted %v\nPASSED!\n\n", got, want)
	}
}