package checker

import (
	"fmt"
	"testing"
	"time"

	storage "github.com/kapustkin/go_guard/pkg/rest-server/dal/storage"
	"github.com/kapustkin/go_guard/pkg/rest-server/dal/storage/inmemory"
)

type testPairCheck struct {
	backet *storage.Bucket
	limit  int
	result bool
}

// nolint: gochecknoglobals
var testsCheckBucket = []testPairCheck{
	{&storage.Bucket{Value: 2}, 3, true},
	{&storage.Bucket{Value: 3}, 3, true},
	{&storage.Bucket{Value: 3}, 0, true},
	{&storage.Bucket{Created: time.Now(), Value: 3}, 3, false},
	{&storage.Bucket{Created: time.Now().Add(time.Second * -59), Value: 3}, 3, false},
	{&storage.Bucket{Created: time.Now().Add(time.Second * -60), Value: 3}, 3, true},
	{&storage.Bucket{Created: time.Now().Add(time.Second * -30), Value: 2}, 3, true},
	{&storage.Bucket{Created: time.Now().Add(time.Second * -30), Value: 3}, 3, false},
}

type testPairProcess struct {
	limit     int
	ident     string
	result    bool
	resultErr error
}

// nolint: gochecknoglobals
var testsProcessBucket = []testPairProcess{
	{3, "test1", true, nil},
	{3, "test1", true, nil},
	{3, "test1", true, nil},
	{3, "test1", false, nil},
	{3, "test1", false, nil},
	{1, "test2", true, nil},
	{1, "test2", false, nil},
	{4, "test1", true, nil},
	{4, "", false, fmt.Errorf("ident must be not empty")},
}

func TestCheckBucket(t *testing.T) {
	for _, pair := range testsCheckBucket {
		_, v := checkBucket(pair.backet, pair.limit)
		if v != pair.result {
			t.Error(
				"For", pair.backet,
				"with limit", pair.limit,
				"expected", pair.result,
				"got", v,
			)
		}
	}
}

func TestProcessBucket(t *testing.T) {
	var db = inmemory.Init()

	for _, pair := range testsProcessBucket {
		res, err := ProcessBucket(db, pair.ident, pair.limit)
		if err == nil && res != pair.result {
			t.Error(
				"For", pair.ident,
				"with limit", pair.limit,
				"expected", pair.result,
				"got", res,
			)
		}
		if err != nil && err.Error() != pair.resultErr.Error() {
			t.Error(
				"Exceprition for", pair.ident,
				"with limit", pair.limit,
				"expected", pair.result,
				"got ", err.Error(),
			)
		}
	}
}
