package internal

import (
	"testing"
	"time"

	storage "github.com/kapustkin/go_guard/pkg/rest-server/dal"
	"github.com/kapustkin/go_guard/pkg/rest-server/dal/inmemory"
)

type testpair struct {
	backet *storage.Bucket
	limit  int
	ident  string
	result bool
}

// nolint: gochecknoglobals
var testsCheckBucket = []testpair{
	{&storage.Bucket{Value: 2}, 3, "", true},
	{&storage.Bucket{Value: 3}, 3, "", true},
	{&storage.Bucket{Value: 3}, 0, "", true},
	{&storage.Bucket{Created: time.Now(), Value: 3}, 3, "", false},
	{&storage.Bucket{Created: time.Now().Add(time.Second * -59), Value: 3}, 3, "", false},
	{&storage.Bucket{Created: time.Now().Add(time.Second * -60), Value: 3}, 3, "", true},
	{&storage.Bucket{Created: time.Now().Add(time.Second * -30), Value: 2}, 3, "", true},
	{&storage.Bucket{Created: time.Now().Add(time.Second * -30), Value: 3}, 3, "", false},
}

// nolint: gochecknoglobals
var testsProcessBucket = []testpair{
	{&storage.Bucket{}, 3, "test1", true},
	{&storage.Bucket{}, 3, "test1", true},
	{&storage.Bucket{}, 3, "test1", true},
	{&storage.Bucket{}, 3, "test1", false},
	{&storage.Bucket{}, 3, "test1", false},
	{&storage.Bucket{}, 1, "test2", true},
	{&storage.Bucket{}, 1, "test2", false},
	{&storage.Bucket{}, 4, "test1", true},
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
		res, _ := ProcessBucket(db, pair.ident, pair.limit)
		if res != pair.result {
			t.Error(
				"For", pair.backet,
				"with limit", pair.limit,
				"expected", pair.result,
				"got", res,
			)
		}
	}
}
