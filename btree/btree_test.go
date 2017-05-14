package main

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestRoot(t *testing.T) {
  assert := assert.New(t);

  var root *Node = create_new_tree();
  assert.Equal(root.total_keys, int8(0), "The total keys filled should be 0")

  for index := 0; index < CHILD_SIZE; index++ {
    assert.Nil(root.child[index], "The son {index} should be null")
  }

  for index := 0; index < KEY_SIZE; index++ {
    assert.Equal(root.key[index], int64(0), "The key {index} should be empty")
  }
}

func TestInsert(t *testing.T) {
}

func TestSearch(t *testing.T) {
}

func TestRemove(t *testing.T) {
}

func TestBalance(t *testing.T) {
}
