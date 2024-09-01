package main

import (
	"fmt"
)

func getHeight(root int, tree map[int][]int) int {
	height := 1
	if tree[root] == nil {
		return height
	}
	for _, child := range tree[root] {
		height = max(height, 1+getHeight(child, tree))
	}
	return height

}

func main() {
	var n, root int
	fmt.Scan(&n)
	input := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&input[i])
		if input[i] == -1 {
			root = i
		}
	}
	tree := make(map[int][]int)

	for child, parent := range input {
		if parent == -1 {
			continue
		}
		tree[parent] = append(tree[parent], child)
	}

	fmt.Println(getHeight(root, tree))
}
