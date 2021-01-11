package binary_tree

import "fmt"

type TreeNode struct {
	data int32

	left  *TreeNode
	right *TreeNode
}

type BinaryTree struct {
	Root *TreeNode
	Size int32
}

func (tree *BinaryTree) InsertValue(val int32) *TreeNode {
	if tree == nil {
		return nil
	}
	node := tree.Root.traverseToInsertNode(val)
	if node == nil {
		return nil
	}
	if tree.Root == nil {
		tree.Root = node
	}
	tree.Size += 1
	return node
}

func (tree *BinaryTree) Display(order string) {
	if tree == nil {
		return
	}
	if order == "in_order" {
		fmt.Println("In-order traverse:")
		tree.Root.inOrderTraverse()
	} else if order == "pre_order" {
		fmt.Println("Pre-order traverse:")
		tree.Root.preOrderTraverse()
	} else if order == "post_order" {
		fmt.Println("Post-order traverse:")
		tree.Root.postOrderTraverse()
	} else {
		panic("Traverse order is not correctly specified!")
	}
	fmt.Println()
}

func (tree *BinaryTree) DeleteValue(val int32) {
	if tree == nil {
		return
	}

	parent, to_delete := tree.Root.traverseToFindNode(val)
	if to_delete == nil {
		return
	}

	tree.Size -= 1
	if to_delete.right == nil {
		// if the node only has no child or only left child
		if to_delete == tree.Root {
			tree.Root = to_delete.left
		} else if to_delete == parent.left {
			parent.left = to_delete.left
		} else {
			parent.right = to_delete.left
		}
	} else if to_delete.left == nil {
		// if the node only has right child
		if to_delete == tree.Root {
			tree.Root = to_delete.right
		} else if to_delete == parent.left {
			parent.left = to_delete.right
		} else {
			parent.right = to_delete.right
		}
	} else {
		// if the node has two valid children
		successor_parent, successor := to_delete.traverseToFindSuccessor(parent)
		if successor == successor_parent.left {
			successor_parent.left = successor.right
		} else {
			successor_parent.right = successor.right
		}
		to_delete.data = successor.data
	}
}

func (root *TreeNode) traverseToInsertNode(val int32) *TreeNode {
	new_node := &TreeNode{data: val, left: nil, right: nil}

	if root == nil {
		root = new_node
		return new_node
	}
	curr := root
	if val == curr.data {
		return nil
	} else if val < curr.data {
		if curr.left == nil {
			curr.left = new_node
			return new_node
		} else {
			curr = curr.left
		}
	} else {
		if curr.right == nil {
			curr.right = new_node
			return new_node
		} else {
			curr = curr.right
		}
	}
	return curr.traverseToInsertNode(val)
}

func (root *TreeNode) inOrderTraverse() {
	if root == nil {
		return
	}
	root.left.inOrderTraverse()
	fmt.Printf("%d ", root.data)
	root.right.inOrderTraverse()
}

func (root *TreeNode) preOrderTraverse() {
	if root == nil {
		return
	}
	fmt.Printf("%d ", root.data)
	root.left.preOrderTraverse()
	root.right.preOrderTraverse()
}

func (root *TreeNode) postOrderTraverse() {
	if root == nil {
		return
	}
	root.left.postOrderTraverse()
	root.right.postOrderTraverse()
	fmt.Printf("%d ", root.data)
}

func (root *TreeNode) traverseToFindNode(val int32) (*TreeNode, *TreeNode) {
	if root == nil {
		return nil, nil
	}
	var parent *TreeNode = nil
	current := root
	for {
		if current == nil {
			return parent, current
		}
		if val < current.data {
			parent = current
			current = current.left
		} else if val > current.data {
			parent = current
			current = current.right
		} else {
			return parent, current
		}
	}
}

func (node *TreeNode) traverseToFindSuccessor(parent *TreeNode) (*TreeNode, *TreeNode) {
	if node == nil {
		return nil, nil
	}
	if node.right == nil {
		return parent, node
	}
	parent = node
	node = node.right
	for {
		if node.left == nil {
			break
		}
		parent = node
		node = node.left
	}
	return parent, node
}
