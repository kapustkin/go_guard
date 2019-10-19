package checker

import (
	"time"

	storage "github.com/kapustkin/go_guard/pkg/rest-server/dal/storage"
)

// ProcessBucket полностью обрабатывает бакет
func ProcessBucket(db storage.Storage, ident string, limit int) (bool, error) {
	bucket, err := db.FindOrCreateBucket(ident)
	if err != nil {
		return false, err
	}
	bckt, result := checkBucket(&bucket, limit)
	err = db.UpdateBucket(ident, bckt)
	if err != nil {
		return false, err
	}
	return result, nil
}

func checkBucket(bucket *storage.Bucket, limit int) (*storage.Bucket, bool) {
	if bucket.Created.Add(time.Second * 60).Before(time.Now()) {
		bucket.Created = time.Now()
		bucket.Value = 1
		return bucket, true
	}

	if bucket.Value < limit {
		bucket.Value++
		return bucket, true
	}

	return bucket, false
}
