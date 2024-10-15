package utils

import (
	"slices"
	"sync"
	"time"
)

type Queue[T any] []T

// Push an item onto the queue
func (q *Queue[T]) Push(item T) {
	*q = append(*q, item)
}

// Consume the next item in line from the queue
func (q *Queue[T]) Next() (T, bool) {
	var zeroValue T
	if len(*q) == 0 {
		return zeroValue, false
	}

	targetVal := (*q)[0]
	*q = slices.Delete(*q, 0, 1)
	return targetVal, true
}

// Queue consumer with a callback function
func (q *Queue[T]) Watch(wg *sync.WaitGroup, callback func(T)) {
	defer wg.Done()

	for {
		// Attempt to dequeue a message
		msg, ok := q.Next()
		if ok {
			// Give the message to caller via callback
			callback(msg)
		} else {
			// Sleep for a short duration to prevent busy waiting
			time.Sleep(100 * time.Millisecond)
		}
	}
}
