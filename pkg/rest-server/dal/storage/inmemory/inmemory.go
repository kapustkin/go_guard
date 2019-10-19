package inmemory

import (
	"fmt"
	"sync"
	"time"

	storage "github.com/kapustkin/go_guard/pkg/rest-server/dal/storage"
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
	storage := make(map[string]storage.Bucket)

	//TODO start gourutine for remove old buckets

	return &Storage{db: &localStorage{data: storage}}
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
		s.db.data[ident] = *bucket
		return nil
	}
	return fmt.Errorf("record %s not found", ident)
}

// RemoveBucket edit event
func (s *Storage) RemoveBucket(ident string) error {
	s.db.Lock()
	defer s.db.Unlock()
	if ident == "" {
		return fmt.Errorf("ident must be not empty")
	}

	if _, ok := s.db.data[ident]; ok {
		delete(s.db.data, ident)
		return nil
	}

	return fmt.Errorf("record %s not found", ident)
}
