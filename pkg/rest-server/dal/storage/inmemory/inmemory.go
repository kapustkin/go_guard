package inmemory

import (
	"fmt"
	"sync"
	"time"

	storage "github.com/kapustkin/go_guard/pkg/rest-server/dal/storage"
	log "github.com/sirupsen/logrus"
)

//nolint
var (
	gcTimeout = 60 * time.Second
	validTime = 60 * time.Second
)

// DB структура хранилища
type Storage struct {
	db *localStorage
}

type localStorage struct {
	sync.RWMutex
	data map[string]storage.Bucket
}

// Init storage
func Init() *Storage {
	db := make(map[string]storage.Bucket)
	storage := &Storage{db: &localStorage{data: db}}

	go cleaner(storage)

	return storage
}

// GetBucket return Bucket
func (s *Storage) FindOrCreateBucket(ident string) (storage.Bucket, error) {
	s.db.RLock()
	defer s.db.RUnlock()

	if ident == "" {
		return storage.Bucket{}, fmt.Errorf("ident must be not empty")
	}

	// bucket exist
	if _, ok := s.db.data[ident]; ok {
		return s.db.data[ident], nil
	}
	// new bucket
	s.db.data[ident] = storage.Bucket{Created: time.Now()}

	return s.db.data[ident], nil
}

// UpdateBucket element to storage
func (s *Storage) UpdateBucket(ident string, bucket *storage.Bucket) error {
	s.db.Lock()
	defer s.db.Unlock()

	if ident == "" {
		return fmt.Errorf("ident must be not empty")
	}
	// bucket exist
	if _, ok := s.db.data[ident]; ok {
		bucket.Updated = time.Now()
		s.db.data[ident] = *bucket

		return nil
	}

	return fmt.Errorf("record %s not found", ident)
}

// RemoveBuckets
func (s *Storage) RemoveBuckets(idents ...string) error {
	s.db.Lock()
	defer s.db.Unlock()

	if len(idents) == 0 || len(idents) == 1 && idents[0] == "" {
		return fmt.Errorf("ident must be not empty")
	}

	for _, ident := range idents {
		delete(s.db.data, ident)
	}

	return nil
}

// cleaner remove old buckets
func cleaner(s *Storage) {
	for {
		time.Sleep(gcTimeout)
		s.db.Lock()
		count := 0

		for name, item := range s.db.data {
			if item.Created.After(time.Now().Add(-validTime)) {
				delete(s.db.data, name)
				count++
			}
		}

		if count > 0 {
			log.Infof("gc removed %d buckets", count)
		}
		s.db.Unlock()
	}
}
