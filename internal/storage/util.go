package storage

import (
	"errors"
)

var instance *Storage

func createStorage(fileName string) *Storage {
	s := new(Storage)
	s.fileName = fileName

	return s
}

// Init create the new instance of Storage
func Init(fileName string) error {
	instance = createStorage(fileName)
	return instance.Load()
}

// GetInstance return the instance of Storage
func GetInstance() (*Storage, error) {
	if instance == nil {
		return instance, errors.New("storage instance don't initialized")
	}
	return instance, nil
}

func findMaxKey[T interface{}](m map[int64]T) int64 {
	var maxKey int64
	for k := range m {
		maxKey = k
		break
	}
	for key := range m {
		if key > maxKey {
			maxKey = key
		}
	}
	return maxKey
}
