package main

import "fmt"

type BinaryTree struct {
  root *Node
}

type Node struct {
  left *Node
  right *Node
  parent *Node
  value int
}

func (tree *BinaryTree) add(node Node) {
  if tree.root == nil {
    tree.root = &node
  } else {
    tree.root.add(node)
  }
}

func (node *Node) add(new_node Node) {
  if new_node.value <= node.value {
    if node.left == nil {
      node.left = &new_node
      new_node.parent = node
    } else {
      new_node.left.add(new_node)
    }
  } else {
    if node.right == nil {
      node.right = &new_node
      new_node.parent = node
    } else {
      new_node.right.add(new_node)
    }
  }
}

func main() {
  root := Node{nil, nil, nil, 10}
  tree := BinaryTree{&root}

  new_node := Node{nil, nil, nil, 60}
  tree.add(new_node)

  fmt.Println(tree.root.right.parent.value)
}
