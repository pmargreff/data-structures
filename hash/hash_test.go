package main

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestCreateHash(t *testing.T) {
  hash := CreateHash()
  assert.Equal(t, hash.size, INITIAL_SIZE)
  assert.Equal(t, hash.filledBuckets, 0)
  assert.Equal(t, hash.loadFactor, 0.0)
}

func TestAddBucket(t *testing.T) {
  hash := CreateHash()
  key := "abc"
  AddBucket(hash, bucket{key: key})
  assert.Equal(t, 1, hash.filledBuckets)
  assert.Equal(t, "abc", hash.buckets[HashFunction(key)].key)
}

// hash function related spectations

func TestHashFunction(t *testing.T)  {
  assert.Equal(t, 1, HashFunction("a"))
  assert.Equal(t, 5, HashFunction("Pablo"))
  assert.Equal(t, 6, HashFunction("Random"))
  assert.Equal(t, 6, HashFunction("123456"))
}
