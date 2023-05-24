// Copyright (c) 2023 Luka Ivanovic
// This code is licensed under MIT licence (see LICENCE for details)

package main

import "fmt"

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
