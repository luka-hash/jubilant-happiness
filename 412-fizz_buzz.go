// Copyright (c) 2023 Luka Ivanovic
// This code is licensed under MIT licence (see LICENCE for details)

// https://leetcode.com/problems/fizz-buzz/
// tldr:
// FizzBuzz problem :/ LeetCode's "challenges for new users"

package main

import (
	"fmt"
	"strconv"
)

func fizzBuzz(n int) (res []string) {
	for i := 1; i <= n; i += 1 {
		if i%15 == 0 {
			res = append(res, "FizzBuzz")
		} else if i%5 == 0 {
			res = append(res, "Buzz")
		} else if i%3 == 0 {
			res = append(res, "Fizz")
		} else {
			res = append(res, strconv.Itoa(i))
		}
	}
	return res
}

func main() {
	fmt.Println(fizzBuzz(3))
}
