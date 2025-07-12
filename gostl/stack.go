package gostl

import (
	"github.com/RajjjAryan/gostl/ds/container"
	"github.com/RajjjAryan/gostl/ds/deque"
	"github.com/RajjjAryan/gostl/utils/sync"
)

var (
	sDefaultLocker sync.FakeLocker
)

// SOptions holds the Stack's options
type SOptions[T any] struct {
	locker    sync.Locker
	container container.Container[T]
}

// SOption is a function type used to set SOptions
type SOption[T any] func(option *SOptions[T])

// Stack is a last-in-first-out data structure
type Stack[T any] struct {
	container container.Container[T]
	locker    sync.Locker
}

// New creates a new stack
func NewStack[T any](opts ...SOption[T]) *Stack[T] {
	option := SOptions[T]{
		locker:    sDefaultLocker,
		container: deque.New[T](),
	}
	for _, opt := range opts {
		opt(&option)
	}

	return &Stack[T]{
		container: option.container,
		locker:    option.locker,
	}
}

// Size returns the amount of elements in the stack
func (s *Stack[T]) Size() int {
	s.locker.RLock()
	defer s.locker.RUnlock()

	return s.container.Size()
}

// Empty returns true if the stack is empty, otherwise returns false
func (s *Stack[T]) Empty() bool {
	s.locker.RLock()
	defer s.locker.RUnlock()

	return s.container.Empty()
}

// Push pushes a value to the stack
func (s *Stack[T]) Push(value T) {
	s.locker.Lock()
	defer s.locker.Unlock()

	s.container.PushBack(value)
}

// Top returns the top value in the stack
func (s *Stack[T]) Top() T {
	s.locker.RLock()
	defer s.locker.RUnlock()

	return s.container.Back()
}

// Pop removes the top value in the stack and returns it
func (s *Stack[T]) Pop() T {
	s.locker.Lock()
	defer s.locker.Unlock()

	return s.container.PopBack()
}

// Clear clears all elements in the stack
func (s *Stack[T]) Clear() {
	s.locker.Lock()
	defer s.locker.Unlock()

	s.container.Clear()
}

// String returns a string representation of the stack
func (s *Stack[T]) String() string {
	s.locker.RLock()
	defer s.locker.RUnlock()

	return s.container.String()
}
