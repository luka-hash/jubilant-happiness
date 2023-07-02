// Copyright (c) 2023 Luka Ivanovic
// This code is licensed under MIT licence (see LICENCE for details)

// https://leetcode.com/problems/string-to-integer-atoi/
// 1. Read in and ignore any leading whitespace.
// 2. Check if the next character (if not already at the end of the string) is '-'
// 		or '+'. Read this character in if it is either. This determines if the
//		final	result is negative or positive respectively. Assume the result is
//    positive if neither is present.
// 3. Read in next the characters until the next non-digit character or the end
// 		of the input is reached. The rest of the string is ignored.
// 4. Convert these digits into an integer (i.e. "123" -> 123, "0032" -> 32). If
// 		no digits were read, then the integer is 0. Change the sign as necessary
// 		(from step 2).
// 5. If the integer is out of the 32-bit signed integer range [-2^31, 2^31 - 1],
// 		then clamp the integer so that it remains in the range. Specifically,
// 		integers less than -2^31 should be clamped to -2^31, and integers greater
// 		than 2^31 - 1 should be clamped to 2^31 - 1.
// 6. Return the integer as the final result.

package main

import (
	"fmt"
)

type State int

const (
	Start State = iota
	Sign
	Number
	After
)

func myAtoi(input string) int {
	state := Start
	sign := 1
	number := 0

loop:
	for i := range input {
		switch state {
		case Start:
			if input[i] == ' ' {
				// state = Start
				continue
			} else if input[i] == '+' {
				state = Sign
				// sign = 1
			} else if input[i] == '-' {
				sign = -1
				state = Sign
			} else if input[i] >= 48 && input[i] <= 57 {
				digit := int(input[i] - '0')
				number = number*10 + digit
				if sign == 1 && number > 2147483647 {
					return 2147483647
				} else if sign == -1 && number > 2147483648 {
					return -2147483648
				}
				state = Number
			} else {
				return 0
			}
		case Sign:
			if input[i] >= 48 && input[i] <= 57 {
				digit := int(input[i] - '0')
				number = number*10 + digit
				if sign == 1 && number > 2147483647 {
					return 2147483647
				} else if sign == -1 && number > 2147483648 {
					return -2147483648
				}
				state = Number
			} else {
				return 0
			}
		case Number:
			if input[i] >= 48 && input[i] <= 57 {
				digit := int(input[i] - '0')
				number = number*10 + digit
				if sign == 1 && number > 2147483647 {
					return 2147483647
				} else if sign == -1 && number > 2147483648 {
					return -2147483648
				}
				// state = Number
			} else {
				state = After
			}
		case After:
			break loop
		}
	}
	return sign * number
}

func main() {
	fmt.Println("herro")
	fmt.Println(myAtoi("9223372036854775808"))
}
