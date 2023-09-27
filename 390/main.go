// Copyright (c) 2023 Luka Ivanović
// This code is licensed under MIT licence (see LICENCE for details)

// https://leetcode.com/problems/elimination-game/

package main

import (
	"fmt"
)

func lastRemaining(n int) int {
	p := 1
	s := 1
	v := true
	for n > 1 {
		if v {
			p += s
		} else {
			if n%2 == 1 {
				p += s
			}
		}
		s *= 2
		n /= 2
		v = !v
	}
	return p
}

func main() {
	fmt.Println(lastRemaining(9))
	fmt.Println(lastRemaining(1))
}
