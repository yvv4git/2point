package components

import (
	"sync"
	"time"
)

const maxCount = 10
const timeWaitMinuteCheck = 1

// StorrageSimple is a simple realization of Storager interface.
type StorageSimple struct {
	mx         sync.RWMutex
	counter    map[int]int
	startTimer bool
}

// Increment for incrementing value.
func (s *StorageSimple) Increment(userId int) {
	s.mx.Lock()
	defer s.mx.Unlock()
	s.counter[userId]++
	if s.startTimer {
		s.startTimer = false
		go func() {
			ticker := time.NewTicker(time.Minute * timeWaitMinuteCheck)
			for {
				select {
				case <-ticker.C:
					// re-initializing map
					s.counter = make(map[int]int)
				}
			}
		}()
	}
}

// GetValue method for return value.
func (s *StorageSimple) GetValue() int {
	s.mx.RLock()
	defer s.mx.RUnlock()

	resultCount := 0
	for _, v := range s.counter {
		if v >= maxCount {
			resultCount++
		}
	}

	return resultCount
}
