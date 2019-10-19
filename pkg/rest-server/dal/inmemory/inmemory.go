package inmemory

import (
	"fmt"
	"sync"
	"time"

	storage "github.com/kapustkin/go_guard/pkg/rest-server/dal"
)

// DB структура хранилища
type DB struct {
	db *database
}

type database struct {
	sync.RWMutex
	data map[string]storage.Bucket
}

// Init storage
func Init() *DB {
	storage := make(map[string]storage.Bucket)
	return &DB{db: &database{data: storage}}
}

// GetBucket return Bucket
func (d *DB) FindOrCreateBucket(ident string) (storage.Bucket, error) {
	d.db.RLock()
	defer d.db.RUnlock()

	// bucket exist
	if _, ok := d.db.data[ident]; ok {
		return d.db.data[ident], nil
	}
	// new bucket
	d.db.data[ident] = storage.Bucket{Created: time.Now()}
	return d.db.data[ident], nil
}

// UpdateBucket element to storage
func (d *DB) UpdateBucket(ident string, bucket *storage.Bucket) error {
	d.db.Lock()
	defer d.db.Unlock()
	// bucket exist
	if _, ok := d.db.data[ident]; ok {
		d.db.data[ident] = *bucket
		return nil
	}
	return fmt.Errorf("record %s not found", ident)
}

// RemoveBucket edit event
func (d *DB) RemoveBucket(ident string) error {
	d.db.Lock()
	defer d.db.Unlock()

	if _, ok := d.db.data[ident]; ok {
		delete(d.db.data, ident)
		return nil
	}

	return fmt.Errorf("record %s not found", ident)
}
