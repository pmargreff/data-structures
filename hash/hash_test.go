package main

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestCreateHash(t *testing.T) {
  hash := CreateHash()
  assert.Equal(t, hash.size, INITIAL_SIZE)
}
