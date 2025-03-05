package main

import (
	"fmt"
)

// Structure for AVL tree node
type Node struct {
	value  int
	height int
	left   *Node
	right  *Node
}

// Helper function to get the height of a node
func height(n *Node) int {
	if n == nil {
		return 0
	}
	return n.height
}

// Calculate Balance Factor of a node
func getBalance(n *Node) int {
	if n == nil {
		return 0
	}
	return height(n.left) - height(n.right)
}

// Right Rotation
func rightRotate(y *Node) *Node {
	if y == nil || y.left == nil {
		return y
	}
	x := y.left
	T2 := x.right

	x.right = y
	y.left = T2

	y.height = max(height(y.left), height(y.right)) + 1
	x.height = max(height(x.left), height(x.right)) + 1

	return x
}

// Left Rotation
func leftRotate(x *Node) *Node {
	if x == nil || x.right == nil {
		return x
	}
	y := x.right
	T2 := y.left

	y.left = x
	x.right = T2

	x.height = max(height(x.left), height(x.right)) + 1
	y.height = max(height(y.left), height(y.right)) + 1

	return y
}

// Insert function to add a node in AVL tree
func insert(node *Node, value int) *Node {
	if node == nil {
		return &Node{value: value, height: 1}
	}

	if value < node.value {
		node.left = insert(node.left, value)
	} else if value > node.value {
		node.right = insert(node.right, value)
	} else {
		return node // Duplicate values are not allowed
	}

	// Update height of the current node
	node.height = 1 + max(height(node.left), height(node.right))

	// Get Balance Factor to check for imbalance
	balance := getBalance(node)

	// Cases of imbalance and their corrections
	// LL Case (Right Rotation)
	if balance > 1 && value < node.left.value {
		return rightRotate(node)
	}

	// RR Case (Left Rotation)
	if balance < -1 && value > node.right.value {
		return leftRotate(node)
	}

	// LR Case (Left-Right Rotation)
	if balance > 1 && value > node.left.value {
		node.left = leftRotate(node.left)
		return rightRotate(node)
	}

	// RL Case (Right-Left Rotation)
	if balance < -1 && value < node.right.value {
		node.right = rightRotate(node.right)
		return leftRotate(node)
	}

	return node
}

// Helper function to find the maximum of two integers
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Print tree in InOrder traversal (for testing)
func inOrder(node *Node) {
	if node != nil {
		inOrder(node.left)
		fmt.Printf("%d ", node.value)
		inOrder(node.right)
	}
}

func main() {
	var root *Node
	values := []int{12, 65, 87, 42, 75, 2, 6, 8, 34, 76, 54, 10, 15, 16, 20}
	for _, v := range values {
		root = insert(root, v)
	}

	fmt.Println("AVL tree in InOrder traversal:")
	inOrder(root)
	fmt.Println()
	fmt.Println(root.value)
}
