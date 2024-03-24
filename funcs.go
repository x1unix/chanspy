package chanspy

import (
	"unsafe"

	"github.com/x1unix/chanspy/internal/chantype"
)

type opts struct {
	lock bool
}

// Option is a functional option for ValueOf function.
type Option func(*opts)

// WithLock locks a channel before performing any operations.
//
// Please note that locking a channel may cause a deadlock if the channel is used.
var WithLock Option = func(o *opts) {
	o.lock = true
}

// ValueOf returns a Chan instance for a given channel.
//
// Second argument sync controls whether channel operations should lock a channel.
// Please note that locking a channel may cause a deadlock if the channel is used.
func ValueOf[T any](ch chan T, options ...Option) Chan {
	o := opts{}
	for _, opt := range options {
		opt(&o)
	}

	return chantype.FromPtr(unsafe.Pointer(&ch), o.lock)
}

// IsClosed reports whether a channel is closed.
func IsClosed[T any](ch chan T) bool {
	return ValueOf(ch).Closed()
}
