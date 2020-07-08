/*
CHRIS FELLI, 2019
Calculate x^n
*/

//package main ...
package main

import (
	"fmt"
	"testing"
)

func TestMyPow1(t *testing.T) {
	got := myPow(2, 10)
	expected := 1024.0

	if got != expected {
		t.Errorf("Got %f, expected %f\nFAILED!\n", got, expected)
	} else {
		fmt.Printf("Got %f, expected %f\nPASSED!\n\n", got, expected)
	}
}
