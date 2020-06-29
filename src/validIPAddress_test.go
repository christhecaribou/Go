/*
CHRIS FELLI, 2019
Validates IPv4 and IPv6 addresses
IPv4: 192.168.255.255
IPv6: 2001:0db8:85a3:0:0:8A2E:0370:7334
*/

//package main ...
package main

import (
	"fmt"
	"testing"
)

func TestValidIPAddress1(t *testing.T) {
	testcase := "172.16.254.1"
	got := validIPAddress(testcase)
	want := "IPv4"

	fmt.Printf("Attempting %q ...\n", testcase)
	if got != want {
		t.Errorf("Got %v, wanted %v\nFAILED!\n", got, want)
	} else {
		fmt.Printf("Got %v, wanted %v\nPASSED!\n\n", got, want)
	}
}
func TestValidIPAddress2(t *testing.T) {
	testcase := "2001:0db8:85a3:0:0:8A2E:0370:7334"
	got := validIPAddress(testcase)
	want := "IPv6"

	fmt.Printf("Attempting %q ...\n", testcase)
	if got != want {
		t.Errorf("Got %v, wanted %v\nFAILED!\n", got, want)
	} else {
		fmt.Printf("Got %v, wanted %v\nPASSED!\n\n", got, want)
	}
}
func TestValidIPAddress3(t *testing.T) {
	testcase := "256.256.256.256"
	got := validIPAddress(testcase)
	want := "Neither"

	fmt.Printf("Attempting %q ...\n", testcase)
	if got != want {
		t.Errorf("Got %v, wanted %v\nFAILED!\n", got, want)
	} else {
		fmt.Printf("Got %v, wanted %v\nPASSED!\n\n", got, want)
	}
}
func TestValidIPAddress4(t *testing.T) {
	testcase := "192.168.1.1"
	got := validIPAddress(testcase)
	want := "IPv4"

	fmt.Printf("Attempting %q ...\n", testcase)
	if got != want {
		t.Errorf("Got %v, wanted %v\nFAILED!\n", got, want)
	} else {
		fmt.Printf("Got %v, wanted %v\nPASSED!\n\n", got, want)
	}
}
func TestValidIPAddress5(t *testing.T) {
	testcase := "200441:0db8:85a3:0:0:8A2E:0370:7334"
	got := validIPAddress(testcase)
	want := "Neither"

	fmt.Printf("Attempting %q ...\n", testcase)
	if got != want {
		t.Errorf("Got %v, wanted %v\nFAILED!\n", got, want)
	} else {
		fmt.Printf("Got %v, wanted %v\nPASSED!\n\n", got, want)
	}
}
func TestValidIPAddress6(t *testing.T) {
	testcase := "2001:0Xb8:85a3:0:0:8A2E:0370:7334"
	got := validIPAddress(testcase)
	want := "Neither"

	fmt.Printf("Attempting %q ...\n", testcase)
	if got != want {
		t.Errorf("Got %v, wanted %v\nFAILED!\n", got, want)
	} else {
		fmt.Printf("Got %v, wanted %v\nPASSED!\n\n", got, want)
	}
}
func TestValidIPAddress7(t *testing.T) {
	testcase := "1e1.4.5.6"
	got := validIPAddress(testcase)
	want := "Neither"

	fmt.Printf("Attempting %q ...\n", testcase)
	if got != want {
		t.Errorf("Got %v, wanted %v\nFAILED!\n", got, want)
	} else {
		fmt.Printf("Got %v, wanted %v\nPASSED!\n\n", got, want)
	}
}
func TestValidIPAddress8(t *testing.T) {
	testcase := "2001:0db8:85a3:033:0:8A2E:0370:7334"
	got := validIPAddress(testcase)
	want := "IPv6"

	fmt.Printf("Attempting %q ...\n", testcase)
	if got != want {
		t.Errorf("Got %v, wanted %v\nFAILED!\n", got, want)
	} else {
		fmt.Printf("Got %v, wanted %v\nPASSED!\n\n", got, want)
	}
}
