package main

import "fmt"
import "unicode/utf8"

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
  hash.buckets[HashFunction(new_bucket.key)] = &bucket{key: new_bucket.key}
}

func HashFunction(key string) int{
  return utf8.RuneCountInString(key)
}
