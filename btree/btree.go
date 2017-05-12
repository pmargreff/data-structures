package main

import "fmt"

const KEY_SIZE = 2
const CHILD_SIZE = KEY_SIZE + 1

func main() {
	fmt.Println("Hello GO");
}

type Node struct {
  key [KEY_SIZE] int64;
  child [CHILD_SIZE] *Node;
	root *Node;
	leaf bool;
}

type Tree struct {
  root *Node;
	depth int64;
}
