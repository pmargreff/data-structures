package main

import "fmt"
import "unicode/utf8"

const GLOBAL_DEPTH int = 4
const INITIAL_SIZE int = 100

type node struct {
  key string
  next_node *node
}

type hash struct {
  buckets [INITIAL_SIZE]*node
  filledNodes int
  loadFactor float64
  size int
}

func main()  {
  hash := CreateHash()
  key := "abc"
  AddNode(hash, node{key: key})
  fmt.Println(*hash.buckets[3])
}

func CreateHash() *hash {
  return &hash{size: INITIAL_SIZE, filledNodes: 0, loadFactor: 0.0}
}

func AddNode(hash *hash, new_node node) {
  hash.filledNodes++;
  hash.loadFactor = (float64)(hash.size / hash.filledNodes)
  hash.buckets[HashFunction(new_node.key)] = &node{key: new_node.key}
  // first_
  // for ptr := ; ptr != nil; ptr =  {
  //
  // }
}

func HashFunction(key string) int {
  return utf8.RuneCountInString(key)
}
