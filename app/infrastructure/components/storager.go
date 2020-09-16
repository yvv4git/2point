package components

import "errors"

// Storage is a interface for many storage types.
type Storager interface {
	Increment(userId int)
	GetValue() int
}

// NewStorage is a constructor for Storage.
func NewStorage(typeName string) (Storager, error) {
	switch typeName {
	case "simple":
		var storage StorageSimple
		storage.counter = make(map[int]int)
		storage.startTimer = true
		return &storage, nil
	default:
		return nil, errors.New("Unknown type")
	}
}
