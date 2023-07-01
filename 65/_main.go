// Copyright (c) 2023 Luka Ivanovic
// This code is licensed under MIT licence (see LICENCE for details)

// https://leetcode.com/problems/valid-number/
// tldr:
// Check if string is in a valid number format

package main

type State int
type Transition byte

const (
	// states should be given propper names
	S1 State = iota
	S2
	S3
	S4
	S5
	S6
	S7
	S8
	S9
	ERR
)

// this can be compressed
// e.g.
// by changing Transition from byte to state (or even int) and creating
// special Transition consts, like dot, e, and digit, and converting
// bytes to those Transitions; then map could be shortened to something like
//
//	S1: {
//		DOT: S3,
//		SIGN: S4,
//		DIGIT: S2
//	}
var TransitionTable = map[State]map[Transition]State{
	S1: {
		'.': S3,
		'-': S4,
		'+': S4,
		'0': S2,
		'1': S2,
		'2': S2,
		'3': S2,
		'4': S2,
		'5': S2,
		'6': S2,
		'7': S2,
		'8': S2,
		'9': S2,
	},
	S2: {
		'.': S6,
		'e': S7,
		'E': S7,
		'0': S2,
		'1': S2,
		'2': S2,
		'3': S2,
		'4': S2,
		'5': S2,
		'6': S2,
		'7': S2,
		'8': S2,
		'9': S2,
	},
	S3: {
		'0': S5,
		'1': S5,
		'2': S5,
		'3': S5,
		'4': S5,
		'5': S5,
		'6': S5,
		'7': S5,
		'8': S5,
		'9': S5,
	},
	S4: {
		'0': S2,
		'1': S2,
		'2': S2,
		'3': S2,
		'4': S2,
		'5': S2,
		'6': S2,
		'7': S2,
		'8': S2,
		'9': S2,
		'.': S3,
	},
	S5: {
		'0': S5,
		'1': S5,
		'2': S5,
		'3': S5,
		'4': S5,
		'5': S5,
		'6': S5,
		'7': S5,
		'8': S5,
		'9': S5,
		'e': S7,
		'E': S7,
	},
	S6: {
		'0': S6,
		'1': S6,
		'2': S6,
		'3': S6,
		'4': S6,
		'5': S6,
		'6': S6,
		'7': S6,
		'8': S6,
		'9': S6,
		'e': S7,
		'E': S7,
	},
	S7: {
		'-': S8,
		'+': S8,
		'0': S9,
		'1': S9,
		'2': S9,
		'3': S9,
		'4': S9,
		'5': S9,
		'6': S9,
		'7': S9,
		'8': S9,
		'9': S9,
	},
	S8: {
		'0': S9,
		'1': S9,
		'2': S9,
		'3': S9,
		'4': S9,
		'5': S9,
		'6': S9,
		'7': S9,
		'8': S9,
		'9': S9,
	},
	S9: {
		'0': S9,
		'1': S9,
		'2': S9,
		'3': S9,
		'4': S9,
		'5': S9,
		'6': S9,
		'7': S9,
		'8': S9,
		'9': S9,
	},
}

func transition(s State, t Transition) State {
	ns, v := TransitionTable[s][t]
	if v {
		return ns
	}
	return ERR
}

func isNumber(s string) bool {
	state := S1
	for i := range s {
		state = transition(state, Transition(s[i]))
		if state == ERR {
			return false
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
