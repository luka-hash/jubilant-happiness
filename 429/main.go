// Copyright (c) 2023 Luka Ivanović
// This code is licensed under MIT licence (see LICENCE for details)

// https://leetcode.com/problems/n-ary-tree-level-order-traversal/

package main

import (
	"fmt"
)

type Node struct {
	Val      int
	Children []*Node
}

func levelOrder(root *Node) [][]int {
	if root == nil {
		return nil
	}
	levels := make([][]int, 0)
	queue := []*Node{root}
	for len(queue) > 0 {
		l := len(queue)
		level := make([]int, l)
		for i := range queue {
			level[i] = queue[i].Val
			queue = append(queue, queue[i].Children...)
		}
		levels = append(levels, level)
		queue = queue[l:]
	}
	return levels
}

func main() {
	a6 := Node{
		Val:      6,
		Children: []*Node{},
	}
	a5 := Node{
		Val:      5,
		Children: []*Node{},
	}
	a3 := Node{
		Val:      3,
		Children: []*Node{&a5, &a6},
	}
	a2 := Node{
		Val:      2,
		Children: []*Node{},
	}
	a4 := Node{
		Val:      4,
		Children: []*Node{},
	}
	a1 := Node{
		Val:      1,
		Children: []*Node{&a3, &a2, &a4},
	}
	fmt.Println(levelOrder(&a1))
}
