package binarytree

type BinaryTree struct {
  Root *Node
}

func (tree *BinaryTree) Insert(node Node) {
  if tree.Root == nil {
    tree.Root = &node
  } else {
    tree.Root.insert(node)
  }
}

func (tree *BinaryTree) InOrder(handleNode func(node *Node)) {
  if tree.Root == nil { return }

  tree.Root.InOrder(handleNode)
}

func (tree *BinaryTree) Height() int {
  if tree.Root == nil { return 0 }

  return tree.Root.Height()
}

func (tree *BinaryTree) Find(value int) *Node {
  if tree.Root == nil { return nil }

  return tree.Root.find(value)
}

func (tree *BinaryTree) Min() *Node {
  if tree.Root == nil { return nil }

  return tree.Root.min()
}

func (tree *BinaryTree) Max() *Node {
  if tree.Root == nil { return nil }

  return tree.Root.max()
}
