package main

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestCreateHash(t *testing.T) {
  hash := CreateHash()
  assert.Equal(t, hash.size, INITIAL_SIZE)
  assert.Equal(t, hash.filledNodes, 0)
  assert.Equal(t, hash.loadFactor, 0.0)
}

func TestAddNode(t *testing.T) {
  hash := CreateHash()
  key := "abc"
  AddNode(hash, node{key: key})
  assert.Equal(t, 1, hash.filledNodes)
  assert.Equal(t, "abc", hash.buckets[HashFunction(key)].key)
}

func TestAddTwiceOnSameBucket(t *testing.T) {
  hash := CreateHash()
  firstKey := "abc"
  secondKey := "def"
  AddNode(hash, node{key: firstKey})
  AddNode(hash, node{key: secondKey})
  assert.Equal(t, 2, hash.filledNodes)
  assert.Equal(t, "abc", hash.buckets[HashFunction(firstKey)].key)
  assert.Equal(t, "def", hash.buckets[HashFunction(firstKey)].key)
}

// hash function related spectations

func TestHashFunction(t *testing.T)  {
  assert.Equal(t, 1, HashFunction("a"))
  assert.Equal(t, 5, HashFunction("Pablo"))
  assert.Equal(t, 6, HashFunction("Random"))
  assert.Equal(t, 6, HashFunction("123456"))
}

// test colision

// test rehash
