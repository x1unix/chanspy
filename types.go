package chanspy

import "unsafe"

type Chan interface {
	// Empty reports whether a read from a channel would block (that is, the channel is empty).
	Empty() bool

	// Closed returns true if a channel is closed.
	Closed() bool

	// Data returns a pointer to a channel buffer.
	Data() unsafe.Pointer

	// Pointer returns a raw pointer to underlying channel structure.
	Pointer() unsafe.Pointer

	// Len returns a number of elements in a channel buffer.
	Len() int

	// Cap returns a size of a channel buffer.
	Cap() int
}
