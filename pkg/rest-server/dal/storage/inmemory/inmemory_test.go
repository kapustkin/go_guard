package inmemory

import (
	"fmt"
	"testing"
	"time"

	storage "github.com/kapustkin/go_guard/pkg/rest-server/dal/storage"
)

type testpair struct {
	ident     string
	backet    *storage.Bucket
	limit     int
	result    storage.Bucket
	resultErr error
}

//nolint: gochecknoglobals
var db = Init()

// nolint: gochecknoglobals
var testsFindOrCreateBucket = []testpair{
	{"test", &storage.Bucket{}, 0, storage.Bucket{Created: time.Now()}, nil},
	{"test", &storage.Bucket{}, 0, storage.Bucket{Created: time.Now()}, nil},
	{"", &storage.Bucket{}, 0, storage.Bucket{}, fmt.Errorf("ident must be not empty")},
}

// nolint: gochecknoglobals
var testsUpdateBucket = []testpair{
	{"test", &storage.Bucket{Value: 1}, 2, storage.Bucket{Value: 2}, nil},
	{"test2", &storage.Bucket{}, 0, storage.Bucket{}, fmt.Errorf("record test2 not found")},
	{"", &storage.Bucket{}, 0, storage.Bucket{}, fmt.Errorf("ident must be not empty")},
}

// nolint: gochecknoglobals
var testsRemoveBucket = []testpair{
	{"test", &storage.Bucket{}, 2, storage.Bucket{}, nil},
	{"", &storage.Bucket{}, 0, storage.Bucket{}, fmt.Errorf("ident must be not empty")},
}

func TestFindOrCreateBucket(t *testing.T) {
	for _, pair := range testsFindOrCreateBucket {
		backet, err := db.FindOrCreateBucket(pair.ident)
		if err == nil && backet.Created.Day() != pair.result.Created.Day() ||
			backet.Created.Minute() != pair.result.Created.Minute() ||
			backet.Value != pair.limit {
			t.Error(
				"For", pair.backet,
				"with limit", pair.limit,
				"expected", pair.result,
				"got", backet,
			)
		}
		if err != nil && err.Error() != pair.resultErr.Error() {
			t.Error(
				"For", pair.backet,
				"with limit", pair.limit,
				"expected", pair.resultErr,
				"got", err,
			)
		}
	}
}

func TestUpdateBucket(t *testing.T) {
	for _, pair := range testsUpdateBucket {
		err := db.UpdateBucket(pair.ident, pair.backet)
		if err == nil && pair.resultErr != nil {
			t.Error(
				"For", pair.backet,
				"with limit", pair.limit,
				"expected", pair.resultErr,
				"got", err,
			)
		}

		if pair.resultErr != nil && err != nil && err.Error() != pair.resultErr.Error() {
			t.Error(
				"For", pair.backet,
				"with limit", pair.limit,
				"expected", pair.resultErr,
				"got", err,
			)
		}
	}
}

func TestRemoveBucket(t *testing.T) {
	for _, pair := range testsRemoveBucket {
		err := db.RemoveBuckets(pair.ident)
		if err == nil && pair.resultErr != nil {
			t.Error(
				"For", pair.backet,
				"with limit", pair.limit,
				"expected", pair.resultErr,
				"got", err,
			)
		}

		if pair.resultErr != nil && err != nil && err.Error() != pair.resultErr.Error() {
			t.Error(
				"For", pair.backet,
				"with limit", pair.limit,
				"expected", pair.resultErr,
				"got", err,
			)
		}
	}
}
