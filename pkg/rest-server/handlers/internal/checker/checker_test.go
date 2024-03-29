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
	value  int
}

//nolint: gochecknoglobals
var testsCheckBucket = []testPairCheck{
	{&storage.Bucket{Updated: time.Now(), QuotientUpdated: time.Now(), Value: 2}, 3, true, 3},
	{&storage.Bucket{Updated: time.Now(), QuotientUpdated: time.Now(), Value: 3}, 3, false, 3},

	{&storage.Bucket{Updated: time.Now(), QuotientUpdated: time.Now().Add(time.Second * -5), Value: 10}, 10, false, 10},
	{&storage.Bucket{Updated: time.Now(), QuotientUpdated: time.Now().Add(time.Second * -6), Value: 10}, 10, true, 10},

	{&storage.Bucket{Updated: time.Now(), QuotientUpdated: time.Now().Add(time.Second * -20), Value: 10}, 10, true, 8},

	{&storage.Bucket{Updated: time.Now(), QuotientUpdated: time.Now().Add(time.Second * -15), Value: 8}, 10, true, 7},
}

func TestCheckBucket(t *testing.T) {
	for _, pair := range testsCheckBucket {
		bucket, v := checkBucket(pair.backet, pair.limit)
		if v != pair.result {
			t.Error(
				"For", pair.backet,
				"with limit", pair.limit,
				"expected", pair.result,
				"got", v,
			)
		}

		if bucket.Value != pair.value {
			t.Error(
				"For", pair.backet,
				"with limit", pair.limit,
				"expected bucket value", pair.value,
				"got", bucket.Value,
			)
		}
	}
}

type testPairProcess struct {
	limit     int
	ident     string
	result    bool
	resultErr error
}

//nolint: gochecknoglobals
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
				"Exception for", pair.ident,
				"with limit", pair.limit,
				"expected", pair.result,
				"got ", err.Error(),
			)
		}
	}
}

type testPairAddress struct {
	network string
	address string
	result  bool
	err     error
}

//nolint: gochecknoglobals
var testsIsAddressInNewtork = []testPairAddress{
	{"192.168.1.0/24", "192.168.1.241", true, nil},
	{"192.168.1.0/24", "192.168.2.115", false, nil},
	{"192.168.1.0/30", "192.168.1.2", true, nil},
	{"192.168.1.0/16", "192.168.16.251", true, nil},
	{"192.168.1.0/16", "aaa", false, fmt.Errorf("ip address not corrected")},
	{"this is no ip", "some text", false, fmt.Errorf("network parse error invalid CIDR address: this is no ip")},
}

func TestIsAddressInNewtork(t *testing.T) {
	for _, pair := range testsIsAddressInNewtork {
		res, err := IsAddressInNewtork(pair.network, pair.address)
		if err != nil && err.Error() != pair.err.Error() {
			t.Error(
				"Exception for", pair.network,
				"with limit", pair.address,
				"expected", pair.result,
				"got ", err.Error(),
			)
		}

		if res != pair.result {
			t.Error(
				"For", pair.network,
				"with address", pair.address,
				"expected", pair.result,
				"got", res,
			)
		}
	}
}
