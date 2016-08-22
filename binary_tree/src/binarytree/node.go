package binarytree

type Node struct {
  Left *Node
  Right *Node
  Parent *Node
  Value int
}

func (node *Node) find(value int) *Node {
  if node.Value == value {
    return node
  }

  if value < node.Value && node.Left != nil {
    return node.Left.find(value)
  } else if value > node.Value && node.Right != nil {
    return node.Right.find(value)
  } else {
    return nil
  }
}

func (node *Node) InOrder(handleNode func(node *Node)) {
  if node.Left != nil { node.Left.InOrder(handleNode) }
  handleNode(node)
  if node.Right != nil { node.Right.InOrder(handleNode) }
}

func (node *Node) Height() int {
  if node.Left == nil && node.Right == nil {
    return 1
  }

  var height_Left int = 0
  var height_Right int = 0

  if node.Left != nil {
    height_Left = node.Left.Height()
  }

  if node.Right != nil {
    height_Right = node.Right.Height()
  }

  return 1 + max(height_Left, height_Right)
}

func (node *Node) add(new_node Node) {
  if new_node.Value <= node.Value {
    if node.Left == nil {
      node.Left = &new_node
      new_node.Parent = node
    } else {
      node.Left.add(new_node)
    }
  } else {
    if node.Right == nil {
      node.Right = &new_node
      new_node.Parent = node
    } else {
      node.Right.add(new_node)
    }
  }
}

func (node *Node) min() *Node {
  if node.Left == nil { return node }

  return node.Left.min()
}

func (node *Node) max() *Node {
  if node.Right == nil { return node }

  return node.Right.max();
}

func max(a int, b int) int {
  if a > b { return a }

  return b
}
