package binarytree

import (
  "testing"
  "binarytree"
  "github.com/stretchr/testify/assert"
)

func TestHeight(t *testing.T) {
  node := binarytree.Node{nil, nil, nil, 1}

  assert.Equal(t, node.Height(), 1)

  node.Left = &binarytree.Node{nil, nil, nil, 10}
  assert.Equal(t, node.Height(), 2)

  node.Left.Left = &binarytree.Node{nil, nil, nil, 10}
  assert.Equal(t, node.Height(), 3)

  node.Left.Right = &binarytree.Node{nil, nil, nil, 10}
  assert.Equal(t, node.Height(), 3)

  node.Right = &binarytree.Node{nil, nil, nil, 10}
  assert.Equal(t, node.Height(), 3)

  node.Right.Left = &binarytree.Node{nil, nil, nil, 10}
  assert.Equal(t, node.Height(), 3)

  node.Right.Left.Right = &binarytree.Node{nil, nil, nil, 10}
  assert.Equal(t, node.Height(), 4)

  node.Right = nil
  node.Left = nil
  assert.Equal(t, node.Height(), 1)
}
