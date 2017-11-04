package main

import "fmt"

const GLOBAL_DEPTH int = 4
const INITIAL_SIZE int = 100
type bucket struct {
  key string
}

type hash struct {
  buckets [INITIAL_SIZE]*bucket
  filledBuckets int
  loadFactor float64
  size int
}

func main()  {
  fmt.Println(CreateHash())
}

func CreateHash() *hash{
  return &hash{size: INITIAL_SIZE, filledBuckets: 0, loadFactor: 0.0}
}

func AddBucket(hash *hash, new_bucket bucket) {
  hash.filledBuckets++;
  hash.loadFactor = (float64)(hash.size / hash.filledBuckets)
  hash.buckets[2] = &bucket{key: new_bucket.key}
}
