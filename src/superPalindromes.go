/*
CHRIS FELLI, 2019
Let's say a positive integer is a superpalindrome if it is a palindrome,
and it is also the square of a palindrome.
NOTE:
Though functional, this solution times out LeetCode and is not performant.
A better solution can be found in "optimalSuperPalindromes"
*/

//package main ...
package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
)

func superpalindromesInRange(L string, R string) int {
	counter := 0
	Lint, err := strconv.Atoi(L)
	Rint, err2 := strconv.Atoi(R)
	if err != nil || err2 != nil {
		log.Fatal(err)
	}
	Lsqrt := int(math.Round(math.Sqrt(float64(Lint))))
	Rsqrt := int(math.Round(math.Sqrt(float64(Rint))))
	fmt.Println("Left: ", L, Lsqrt)
	fmt.Println("Right: ", R, Rsqrt)

	// For values between sqrt(L) and sqrt(R)
	for i := Lsqrt; i < Rsqrt+1; i++ {
		if isPalindrome(strconv.Itoa(i)) && isPalindrome(strconv.Itoa(i*i)) {
			counter++
		}
	}
	return counter
}
