package main

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

// test root creation and default values
func TestRoot(t *testing.T) {
  assert := assert.New(t);

  var root *Node = create_new_node();
  assert.Equal(root.total_keys, int8(0), "The total keys filled should be 0")

  for index := 0; index < CHILD_SIZE; index++ {
    assert.Nil(root.child[index], "The son {index} should be null")
  }

  for index := 0; index < KEY_SIZE; index++ {
    assert.Equal(root.key[index], int64(0), "The key {index} should be empty")
  }
}

// test insertion for fill only the first level (root) of tree
func TestIfIsLeaf(t *testing.T) {
  assert := assert.New(t);

  var tree *Node = create_new_node();
  var sub_tree *Node = create_new_node();

  assert.Equal(tree.is_leaf(), true, "The root should be a leaf")

  tree.child[0] = sub_tree

  assert.Equal(tree.is_leaf(), false, "The root shouldn't be a leaf")
  assert.Equal(tree.child[0].is_leaf(), true, "The sub tree should be a leaf")
}

func TestIfNodeHasSpace(t *testing.T) {
  assert := assert.New(t);

  var node *Node = create_new_node();

  assert.Equal(node.has_space(), true, "The node should have free spaces");

  for index := 0; index < KEY_SIZE; index++ {
    node.key[index] = int64(index);
    node.total_keys++;
  }

  assert.Equal(node.has_space(), false, "The node shouldn't have free spaces");

}

// test insertion for fill only the first level (root) of tree
func TestInsertionOneLevel(t *testing.T) {
  assert := assert.New(t);

  var tree *Node = create_new_node();

  for index := 0; index < KEY_SIZE; index++ {
    tree.insert(index)
  }

  for index := 0; index < KEY_SIZE; index++ {
    assert.Equal(tree.key[index], index, "The key {index} should be {index}")
  }
}

func TestInsertionMoreThanKeysSize(t *testing.T) {
}

// try to insert only in the left to have sure the tree stay balanced
func TestInsertionLeftOnly(t *testing.T) {
}

// try to insert only in the right to have sure the tree stay balanced
func TestInsertionRightOnly(t *testing.T) {
}

func TestSearch(t *testing.T) {
}

func TestRemove(t *testing.T) {
}

func TestBalance(t *testing.T) {
}
