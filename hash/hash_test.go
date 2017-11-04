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

func TestAddBucket(t *testing.T)  {
  hash := CreateHash()
  AddBucket(hash, bucket{"abc"})
  assert.Equal(t, 1, hash.filledBuckets)
  assert.Equal(t, "abc", hash.buckets[2].key)
}
