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
      node.left.add(new_node)
    }
  } else {
    if node.right == nil {
      node.right = &new_node
      new_node.parent = node
    } else {
      node.right.add(new_node)
    }
  }
}

func (tree *BinaryTree) height() int {
  if (*tree).root == nil { return 0 }

  return tree.root.height()
}

func (node *Node) height() int {
  if node.left == nil && node.right == nil {
    return 1
  }

  var height_left int = 0
  var height_right int = 0

  if node.left != nil {
    height_left = node.left.height()
  }

  if node.right != nil {
    height_right = node.right.height()
  }

  return 1 + max(height_left, height_right)
}

func max(a int, b int) int {
  if a > b { return a }

  return b
}

func main() {
  root := Node{nil, nil, nil, 30}
  tree := BinaryTree{&root}

  new_node_1 := Node{nil, nil, nil, 45}
  tree.add(new_node_1)

  new_node_2 := Node{nil, nil, nil, 15}
  tree.add(new_node_2)

  fmt.Println(tree.height())
}
