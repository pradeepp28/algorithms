package main

import (
	"fmt"
)

type Tree struct {
	root *Node
}

type Node struct {
	key   byte
	left  *Node
	right *Node
}

// String prints a visual representation of the tree
func (bst *Tree) String() {
	fmt.Println("------------------------------------------------")
	stringify(bst.root, 0)
	fmt.Println("------------------------------------------------")
}

// internal recursive function to print a tree
func stringify(n *Node, level int) {
	if n != nil {
		format := ""
		for i := 0; i < level; i++ {
			format += "       "
		}
		format += "---[ "
		level++
		stringify(n.right, level)
		fmt.Printf(format+"%c\n", n.key)
		stringify(n.left, level)
	}
}

func (t *Tree) insert(data byte) {
	if t.root == nil {
		t.root = &Node{key: data}
	} else {
		t.root.insert(data)
	}
}

func (n *Node) insert(data byte) {
	if data <= n.key {
		if n.left == nil {
			n.left = &Node{key: data}
		} else {
			n.left.insert(data)
		}
	} else {
		if n.right == nil {
			n.right = &Node{key: data}
		} else {
			n.right.insert(data)
		}
	}
}

func printPreOrder(n *Node) {
	if n == nil {
		return
	} else {
		fmt.Printf("%c ", n.key)
		printPreOrder(n.left)
		printPreOrder(n.right)
	}
}

func printPostOrder(n *Node) {
	if n == nil {
		return
	} else {
		printPostOrder(n.left)
		printPostOrder(n.right)
		fmt.Printf("%c ", n.key)
	}
}

func printInOrder(n *Node) {
	if n == nil {
		return
	} else {
		printInOrder(n.left)
		fmt.Printf("%c ", n.key)
		printInOrder(n.right)
	}
}

func bfs(n *Node) {
	if n == nil {
		return
	}
	// fmt.Printf("%c ", n.key)
	if n.left != nil {
		fmt.Printf("%c ", n.left.key)
	}
	if n.right != nil {
		fmt.Printf("%c ", n.right.key)
	}
	bfs(n.left)
	bfs(n.right)
}
func main() {
	var t Tree

	t.insert('F')
	t.insert('D')
	t.insert('B')
	t.insert('K')
	t.insert('E')
	t.insert('C')
	t.insert('A')
	t.insert('L')
	t.insert('H')
	t.insert('I')
	t.insert('G')

	fmt.Printf("Pre Order: ")
	printPreOrder(t.root)
	fmt.Println()
	fmt.Printf("Post Order: ")
	printPostOrder(t.root)
	fmt.Println()
	fmt.Printf("In Order: ")
	printInOrder(t.root)
	fmt.Println()
	t.String()
	bfs(t.root)
}
