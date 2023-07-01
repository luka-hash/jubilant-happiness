// Copyright (c) 2023 Luka Ivanovic
// This code is licensed under MIT licence (see LICENCE for details)

// https://leetcode.com/problems/valid-number/
// tldr:
// Check if string is in a valid number format

package main

const (
	S1 int = iota
	S2
	S3
	S4
	S5
	S6
	S7
	S8
	S9
)

func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
}

func isSign(b byte) bool {
	return b == '-' || b == '+'
}

func isDot(b byte) bool {
	return b == '.'
}

func isE(b byte) bool {
	return b == 'e' || b == 'E'
}

func isNumber(s string) bool {
	state := S1
	for i := range s {
		t := s[i]
		switch state {
		case S1:
			if isDot(t) {
				state = S3
			} else if isDigit(t) {
				state = S2
			} else if isSign(t) {
				state = S4
			} else {
				return false
			}
		case S2:
			if isDigit(t) {
				// state = S2
				continue
			} else if isDot(t) {
				state = S6
			} else if isE(t) {
				state = S7
			} else {
				return false
			}
		case S3:
			if isDigit(t) {
				state = S5
			} else {
				return false
			}
		case S4:
			if isDigit(t) {
				state = S2
			} else if isDot(t) {
				state = S3
			} else {
				return false
			}
		case S5:
			if isDigit(t) {
				// state = S5
				continue
			} else if isE(t) {
				state = S7
			} else {
				return false
			}
		case S6:
			if isDigit(t) {
				// state = S6
				continue
			} else if isE(t) {
				state = S7
			} else {
				return false
			}
		case S7:
			if isSign(t) {
				state = S8
			} else if isDigit(t) {
				state = S9
			} else {
				return false
			}
		case S8:
			if isDigit(t) {
				state = S9
			} else {
				return false
			}
		case S9:
			if isDigit(t) {
				// state = S9
				continue
			} else {
				return false
			}
		}
	}
	if state == S2 || state == S5 || state == S6 || state == S9 {
		return true
	}
	return false
}

func assert(expr bool) {
	if !expr {
		panic("")
	}
}

func main() {
	assert(isNumber("0") == true)
	assert(isNumber("e") == false)
	assert(isNumber(".") == false)
	assert(isNumber("-1.") == true)
}
