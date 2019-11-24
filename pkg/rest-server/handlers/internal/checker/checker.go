package checker

import (
	"fmt"
	"net"
	"time"

	storage "github.com/kapustkin/go_guard/pkg/rest-server/dal/storage"
)

// IsAddressInNewtork
func IsAddressInNewtork(network, addr string) (bool, error) {
	_, subnet, err := net.ParseCIDR(network)
	if err != nil {
		return false, fmt.Errorf("network parse error %v", err)
	}

	ip := net.ParseIP(addr)
	if ip.IsUnspecified() {
		return false, fmt.Errorf("address parse error")
	}

	return subnet.Contains(ip), nil
}

// ProcessBucket полностью обрабатывает бакет
func ProcessBucket(db storage.Storage, ident string, limit int) (bool, error) {
	bucket, err := db.FindOrCreateBucket(ident)
	if err != nil {
		return false, err
	}

	bckt, result := checkBucket(&bucket, limit)
	if !result {
		return result, nil
	}

	// если доступ не разрешен, значит нет смысла обновлять бакет
	if result {
		err = db.UpdateBucket(ident, bckt)
		if err != nil {
			return false, err
		}
	}

	return result, nil
}

func checkBucket(bucket *storage.Bucket, limit int) (*storage.Bucket, bool) {
	var newLimit = int64(60 * 1000 / limit)

	var elapsedFromLastUpdate = time.Since(bucket.QuotientUpdated).Milliseconds()

	var quotient = int(elapsedFromLastUpdate / newLimit)
	if quotient > 0 {
		val := bucket.Value - quotient + 1
		if val < 1 {
			val = 1
		}

		bucket.Value = val
		bucket.QuotientUpdated = time.Now()

		return bucket, true
	}

	if bucket.Value < limit {
		bucket.Value++
		return bucket, true
	}

	return bucket, false
}
