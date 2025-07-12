package gostl

// Locker defines an interface for thread-safe operations
type Locker interface {
	Lock()
	Unlock()
	RLock()
	RUnlock()
}

// FakeLocker implements Locker interface but does nothing
type FakeLocker struct{}

// Lock implements Locker.Lock
func (locker FakeLocker) Lock() {}

// Unlock implements Locker.Unlock
func (locker FakeLocker) Unlock() {}

// RLock implements Locker.RLock
func (locker FakeLocker) RLock() {}

// RUnlock implements Locker.RUnlock
func (locker FakeLocker) RUnlock() {}

var (
	// DefaultLocker is a fake locker that doesn't provide synchronization
	DefaultLocker = FakeLocker{}
)