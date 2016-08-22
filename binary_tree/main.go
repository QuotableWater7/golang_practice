package main

import "fmt"
import "binarytree"

func main() {
  root := binarytree.Node{nil, nil, nil, 30}
  tree := binarytree.BinaryTree{&root}

  fmt.Printf("Height is now %d\n", tree.Height())

  new_node_1 := binarytree.Node{nil, nil, nil, 45}
  tree.Add(new_node_1)

  fmt.Printf("Height is now %d\n", tree.Height())

  new_node_2 := binarytree.Node{nil, nil, nil, 15}
  tree.Add(new_node_2)

  fmt.Printf("Height is now %d\n", tree.Height())

  fmt.Println(tree.Find(30).Value)
  fmt.Println(tree.Min().Value)
  fmt.Println(tree.Max().Value)
}
