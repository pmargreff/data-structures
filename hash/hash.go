package main

import "fmt"

const LOAD_FACTOR int = 4
const INITIAL_SIZE int = 100
type bucket struct {
  key string
}

type hash struct {
  size int
  buckets [INITIAL_SIZE]*bucket
}

func main()  {
  fmt.Println(CreateHash())
}

func CreateHash() *hash{
  return &hash{size: INITIAL_SIZE}
}
