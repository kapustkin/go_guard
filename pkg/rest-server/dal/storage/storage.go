package storage

import (
	"time"
)

// Bucket
type Bucket struct {
	Created time.Time
	Updated time.Time
	Value   int
}

type Storage interface {
	FindOrCreateBucket(ident string) (Bucket, error)
	UpdateBucket(ident string, bucket *Bucket) error
	RemoveBuckets(idents ...string) error
}
