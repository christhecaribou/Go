/*
CHRIS FELLI, 2019
Given a string, returns whether or not it's a palindrome
*/

//package main ...
package main

import (
	"log"
	"regexp"
	"strings"
)

func isPalindrome(s string) bool {
	// RegExp removes punctuation and spaces
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		log.Fatal(err)
	}
	// Execute RegExp & convert to lowercase
	ps := reg.ReplaceAllString(s, "")
	ps = strings.ToLower(ps)
	//fmt.Println("The raw string is: ", s)
	//fmt.Println("The converted string is: ", ps)
	mid := len(ps) / 2
	last := len(ps) - 1

	for i := 0; i < mid; i++ {
		if ps[i] != ps[last-i] {
			return false
		}
	}
	return true
}