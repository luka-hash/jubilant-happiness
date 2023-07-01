// Copyright (c) 2023 Luka Ivanovic
// This code is licensed under MIT licence (see LICENCE for details)

// https://leetcode.com/problems/roman-to-integer/
// tldr:
// Given a Roman numeral, convert it to a decimal number (integer).

package main

import "fmt"

// There are other ways to do this, but this is good enough.
func romanToInt(s string) int {
	value := 0
	values := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	prev := 0
	for i := len(s) - 1; i >= 0; i -= 1 {
		curr := values[s[i]]
		if curr < prev {
			curr *= -1
		}
		prev = curr
		value += curr
	}
	return value
}

func main() {
	fmt.Println(romanToInt("III"))     // 3
	fmt.Println(romanToInt("LVIII"))   // 58
	fmt.Println(romanToInt("MCMXCIV")) // 1994
}
