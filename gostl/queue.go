package gostl

import (
	"github.com/RajjjAryan/gostl/ds/container"
	"github.com/RajjjAryan/gostl/ds/deque"
	"github.com/RajjjAryan/gostl/ds/list/bidlist"
)

var (
	qDefaultLocker FakeLocker
)

// QOptions holds Queue's options
type QOptions[T any] struct {
	locker    Locker
	container container.Container[T]
}

// Option is a function type used to set QOptions
type QOption[T any] func(option *QOptions[T])

// WithContainer is used to set a Queue's underlying container
func WithContainer[T any](c container.Container[T]) QOption[T] {
	return func(option *QOptions[T]) {
		option.container = c
	}
}

// WithListContainer is used to set List as a Queue's underlying container
func WithListContainer[T any]() QOption[T] {
	return func(option *QOptions[T]) {
		option.container = bidlist.New[T]()
	}
}

// Queue is a first-in-first-out data structure
type Queue[T any] struct {
	container container.Container[T]
	locker    Locker
}

// New creates a new queue
func NewQueue[T any](opts ...QOption[T]) *Queue[T] {
	option := QOptions[T]{
		locker:    qDefaultLocker,
		container: deque.New[T](),
	}
	for _, opt := range opts {
		opt(&option)
	}

	return &Queue[T]{
		container: option.container,
		locker:    option.locker,
	}
}

// Size returns the amount of elements in the queue
func (q *Queue[T]) Size() int {
	q.locker.RLock()
	defer q.locker.RUnlock()

	return q.container.Size()
}

// Empty returns true if the queue is empty, otherwise returns false
func (q *Queue[T]) Empty() bool {
	q.locker.RLock()
	defer q.locker.RUnlock()

	return q.container.Empty()
}

// Push pushes a value to the end of the queue
func (q *Queue[T]) Push(value T) {
	q.locker.Lock()
	defer q.locker.Unlock()

	q.container.PushBack(value)
}

// Front returns the front value in the queue
func (q *Queue[T]) Front() T {
	q.locker.RLock()
	defer q.locker.RUnlock()

	return q.container.Front()
}

// Back returns the back value in the queue
func (q *Queue[T]) Back() T {
	q.locker.RLock()
	defer q.locker.RUnlock()

	return q.container.Back()
}

// Pop removes the the front element in the queue, and returns its value
func (q *Queue[T]) Pop() T {
	q.locker.Lock()
	defer q.locker.Unlock()

	return q.container.PopFront()
}

// Clear clears all elements in the queue
func (q *Queue[T]) Clear() {
	q.locker.Lock()
	defer q.locker.Unlock()

	q.container.Clear()
}

// String returns a string representation of the queue
func (q *Queue[T]) String() string {
	q.locker.RLock()
	defer q.locker.RUnlock()

	return q.container.String()
}
