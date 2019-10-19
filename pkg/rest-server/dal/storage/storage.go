package storage

import (
	"time"
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
}
