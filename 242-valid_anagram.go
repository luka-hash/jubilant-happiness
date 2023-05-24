// Copyright (c) 2023 Luka Ivanovic
// This code is licensed under MIT licence (see LICENCE for details)

// https://leetcode.com/problems/valid-anagram/
// tldr:
// given two strings s and t, return true if t is an anagram of s, and false
// otherwise

package main

import "fmt"

func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }
    var s_freq [26]int // zero initialized array to serve as primitive freq map
    var t_freq [26]int
    // since we know we are handling only lowercase english letters, we can
    // ignore rune iteration and just iterate over bytes
    for i := range s { // iterating over s AND t
        s_freq[s[i]-'a'] += 1
        t_freq[t[i]-'a'] += 1
    }
    return s_freq == t_freq
}

func main() {
    fmt.Println(isAnagram("anagram","nagaram"))
    fmt.Println(isAnagram("anagram","naguram"))
}
