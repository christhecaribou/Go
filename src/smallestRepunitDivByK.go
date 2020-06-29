/*
CHRIS FELLI, 2019
Given a positive integer K, you need find the smallest 
positive integer N such that N is divisible by K, and N only contains the digit 1.
Return the length of N.  If there is no such N, return -1.
*/

package main

func smallestRepunitDivByK(K int) int {
	if(K%2==0 || K%5==0){
		return -1
	}
	module := 1
	for i:=1; i<=K; i++ {
		if(module % K == 0){
			return i
		}
		module = (module*10 +1)%K
	}
	return 0
}