// Copyright (c) 2023 Luka Ivanovic
// This code is licensed under MIT licence (see LICENCE for details)

// https://leetcode.com/problems/number-of-steps-to-reduce-a-number-to-zero/
// tldr:
// if number if even, divide it with 2, otherwise sub 1; how many steps till 0

package main

import (
	"fmt"
)

func numberOfSteps(num int) int {
	if num <= 0 {
		return 0
	}
	ones := 0
	zeros := 0
	i := 1
	for i <= num {
		if num&i >= 1 {
			ones += 1
		} else {
			zeros += 1
		}
		i <<= 1
	}
	return ones*2 + zeros - 1
}

func main() {
	fmt.Println(numberOfSteps(123))
}
