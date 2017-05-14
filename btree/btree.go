package main

import "fmt"

const KEY_SIZE = 2;
const CHILD_SIZE = KEY_SIZE + 1;

func main() {
	var tree *Node = create_new_node();
	fmt.Println(tree);
}

type Node struct {
	total_keys int8;
	key [KEY_SIZE] int64;
	child [CHILD_SIZE] *Node;
}

func (node Node) is_leaf () bool {
	if node.child[0] == nil {
		return true;
	}
	return false;
}

func (node Node) has_space () bool {
	if node.total_keys < KEY_SIZE {
		return true;
	}
	return false;
}

func (node Node) insert (key int){
	if node.is_leaf() {
	}
}

func create_new_node () *Node{
	root := new(Node)
	return root;
}
