package storage

import (
	"time"
)

// Bucket
type Bucket struct {
	QuotientUpdated time.Time // last qoute update
	Updated         time.Time // last update
	Value           int
}

type Storage interface {
	FindOrCreateBucket(ident string) (Bucket, error)
	UpdateBucket(ident string, bucket *Bucket) error
	RemoveBuckets(idents ...string) error
}
