package main

import (
	"fmt"
	"strconv"
	"strings"
)

type node struct {
	value int
	left  *node
	right *node
}

func (n node) String() string {
	return strconv.Itoa(n.value)
}

type bst struct {
	root *node
	len  int
}

func (b bst) String() string {
	sb := strings.Builder{}
	b.inOrderTraversal(&sb)
	return sb.String()
}

func (b bst) inOrderTraversal(sb *strings.Builder) {
	b.inOrderTraversalByNode(b.root, sb)
}

func (b bst) inOrderTraversalByNode(root *node, sb *strings.Builder) {
	if root == nil {
		return
	}
	b.inOrderTraversalByNode(root.left, sb)
	sb.WriteString(fmt.Sprintf("%s ", root))
	b.inOrderTraversalByNode(root.right, sb)
}

func (b *bst) add(value int) {
	b.root = b.addByNode(b.root, value)
	b.len++
}

func (b *bst) addByNode(root *node, value int) *node {
	if root == nil {
		return &node{value: value}
	}
	if value < root.value {
		root.left = b.addByNode(root.left, value)
	} else {
		root.right = b.addByNode(root.right, value)
	}

	return root
}

func (b bst) search(value int) (*node, bool) {
	return b.searchByNode(b.root, value)
}

func (b bst) searchByNode(root *node, value int) (*node, bool) {
	if root == nil {
		return nil, false
	}

	if value < root.value {
		return b.searchByNode(root.left, value)
	} else if value > root.value {
		return b.searchByNode(root.right, value)
	} else if value == root.value {
		return root, true
	}

	return nil, false
}

func (b *bst) remove(value int) {
	b.removeByNode(b.root, value)
	b.len--
}

func (b *bst) removeByNode(root *node, value int) *node {
	if root == nil {
		return root
	}

	if value < root.value {
		root.left = b.removeByNode(root.left, value)
	} else if value > root.value {
		root.right = b.removeByNode(root.right, value)
	} else {
		if root.left == nil {
			return root.right
		} else {
			temp := root.left
			for temp != nil {
				temp = temp.right
			}

			root.value = temp.value
			root.left = b.removeByNode(root.left, value)
		}
	}
	return root
}

func main() {
	n := &node{1, nil, nil}
	n.left = &node{0, nil, nil}
	n.right = &node{2, nil, nil}

	b := &bst{
		root: n,
		len:  3,
	}
	fmt.Println(b)

	b.add(4)
	b.add(6)
	b.add(5)

	fmt.Println(b)
	fmt.Println(b.search(6))
	fmt.Println(b.search(3))

	b.remove(4)
	fmt.Println(b)
}
