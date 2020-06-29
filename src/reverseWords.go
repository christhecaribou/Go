/*
CHRIS FELLI, 2019
This algorithm reverses words, such that
"Hello world!" => "world! Hello"
" My name is   Chris" => "Chris is name My"
*/

//package main ...
package main

import "strings"

func reverseWords(s string) string {
	cur := 0
	strs := []string{}
	for cur < len(s) {
		for cur < len(s) && s[cur] == ' ' {
			cur++
		}

		if cur >= len(s) {
			break
		}

		next := cur + 1
		for next < len(s) && s[next] != ' ' {
			next++
		}

		strs = append([]string{s[cur:next]}, strs...)
		cur = next
	}

	return strings.Join(strs, " ")
}
