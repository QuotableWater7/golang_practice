package main

import "fmt"
import "binarytree"

func main() {
  root := binarytree.Node{nil, nil, nil, 30}
  tree := binarytree.BinaryTree{&root}

  new_node_1 := binarytree.Node{nil, nil, nil, 45}
  tree.Add(new_node_1)

  new_node_2 := binarytree.Node{nil, nil, nil, 15}
  tree.Add(new_node_2)

  fmt.Println(tree.Find(30).Value)
  fmt.Println(tree.Min().Value)
  fmt.Println(tree.Max().Value)
}
