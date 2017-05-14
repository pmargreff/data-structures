package main

import "fmt"

const KEY_SIZE = 2
const CHILD_SIZE = KEY_SIZE + 1
func main() {
	fmt.Println("Hello GO");
	var root *Node = create_new_tree();
	fmt.Println(root);
}

type Node struct {
	total_keys int8;
  key [KEY_SIZE] int64;
  child [CHILD_SIZE] *Node;
}

func create_new_tree () *Node{
	structure := new(Node)

	return structure;
}
