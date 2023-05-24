// Copyright (c) 2023 Luka Ivanovic
// This code is licensed under MIT licence (see LICENCE for details)

// https://leetcode.com/problems/contains-duplicate/
// tldr:
// given an integer array nums, return true if any value appears at least twice
// in the array, and false otherwise (if every element is distinct)

package main

import "fmt"

// This solution consumes a bit more memory (it creates an additional map), but,
// it is considerably faster, plus it does not modify original slice (sorting an
// array and then checking for neighbors will modify original slice or consume
// around the same amount of memory, plus it will be slower).
func containsDuplicate(nums []int) bool {
    nums_set := make(map[int]int)
    for _, num:= range nums {
        if nums_set[num]>0 {
            return true
        }
        nums_set[num]=1
    }
    return false
}
func main() {
    fmt.Println(containsDuplicate([]int{1,1,1,3,3,4,3,2,4,2}))
}
