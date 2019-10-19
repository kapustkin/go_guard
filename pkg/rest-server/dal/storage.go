package storage

import (
	"time"
)

const (
	N = 1
	M = 2
	K = 3
)

// Bucket
type Bucket struct {
	Created time.Time
	Value   int
}

type Storage interface {
	FindOrCreateBucket(ident string) (Bucket, error)
	UpdateBucket(ident string, bucket *Bucket) error
	RemoveBucket(ident string) error
	//RemoveOldBuckets(expireTime time.Time) error
}
