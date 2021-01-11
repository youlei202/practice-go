package main

import (
	"fmt"
	"practice-go/binary_search_tree/binary_tree"
)

var arr = [10]int32{40, 50, 30, 60, 20, 50, 10, 15, 32, 53}

func main() {
	tree := &binary_tree.BinaryTree{Root: nil, Size: 0}

	for i := 0; i < len(arr); i++ {
		tree.InsertValue(arr[i])
		fmt.Printf("Node with value %d is inserted.\n", arr[i])
	}
	tree.Display("in_order")
	fmt.Printf("Tree size: %d\n", tree.Size)

	tree.DeleteValue(20)
	tree.Display("in_order")
	tree.DeleteValue(40)
	tree.Display("in_order")

	tree.InsertValue(100)
	tree.Display("in_order")
	fmt.Printf("Tree size: %d\n", tree.Size)
}
