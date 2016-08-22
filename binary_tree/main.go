package main

import "fmt"

type BinaryTree struct {
  root *Node
}

func (tree *BinaryTree) add(node Node) {
  if tree.root == nil {
    tree.root = &node
  } else {
    tree.root.add(node)
  }
}

func (tree *BinaryTree) height() int {
  if (*tree).root == nil { return 0 }

  return tree.root.height()
}

func (tree *BinaryTree) find(value int) *Node {
  if tree.root == nil { return nil }

  return tree.root.find(value)
}

func (tree *BinaryTree) min() *Node {
  if tree.root == nil { return nil }

  return tree.root.min()
}

func (tree *BinaryTree) max() *Node {
  if tree.root == nil { return nil }

  return tree.root.max()
}

type Node struct {
  left *Node
  right *Node
  parent *Node
  value int
}

func (node *Node) find(value int) *Node {
  if node.value == value {
    return node
  }

  if value < node.value && node.left != nil {
    return node.left.find(value)
  } else if value > node.value && node.right != nil {
    return node.right.find(value)
  } else {
    return nil
  }
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

func (node *Node) min() *Node {
  if node.left == nil { return node }

  return node.left.min()
}

func (node *Node) max() *Node {
  if node.right == nil { return node }

  return node.right.max();
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

  fmt.Println(tree.find(30).value)
  fmt.Println(tree.min().value)
  fmt.Println(tree.max().value)
}
