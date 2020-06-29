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
	want := 1024.0

	if got != want {
		t.Errorf("Got %f, wanted %f\nFAILED!\n", got, want)
	} else {
		fmt.Printf("Got %f, wanted %f\nPASSED!\n\n", got, want)
	}
}
