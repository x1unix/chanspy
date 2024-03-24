package chantype

import "unsafe"

// See: /src/runtime/chan.go:55
type waitq struct {
	first uintptr
	last  uintptr
}

// See: /src/runtime/runtime2.go:164
type mutex struct {
	// Empty struct if lock ranking is disabled, otherwise includes the lock rank
	lockRankStruct
	// Futex-based impl treats it as uint32 key,
	// while sema-based impl as M* waitm.
	// Used to be a union, but unions break precise GC.
	key uintptr
}

// Copy of hchan struct with first necessary fields.
// See: /src/runtime/chan.go
type hchan struct {
	qcount   uint           // total data in the queue
	dataqsiz uint           // size of the circular queue
	buf      unsafe.Pointer // points to an array of dataqsiz elements
	elemsize uint16
	closed   uint32
	timer    uintptr // timer feeding this chan
	elemtype uintptr // element type
	sendx    uint    // send index
	recvx    uint    // receive index
	recvq    waitq   // list of recv waiters
	sendq    waitq   // list of send waiters

	// lock protects all fields in hchan, as well as several
	// fields in sudogs blocked on this channel.
	//
	// Do not change another G's status while holding this lock
	// (in particular, do not ready a G), as this can deadlock
	// with stack shrinking.
	lock mutex
}
