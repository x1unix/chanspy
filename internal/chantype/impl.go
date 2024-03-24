// Package chantype contains internal implementation of channel type reflection.
// Here be dragons.
package chantype

import "unsafe"

type Chan struct {
	// hChan holds original channel struct.
	hChan *hchan

	// sync controls whether channel operations should respect mutexes.
	sync bool
}

func (c Chan) Closed() bool {
	if c.hChan == nil {
		return true
	}

	c.lock()
	defer c.unlock()
	return c.hChan.closed != 0
}

func (c Chan) Data() unsafe.Pointer {
	if c.hChan == nil {
		return nil
	}

	c.lock()
	defer c.unlock()
	return c.hChan.buf
}

func (c Chan) Pointer() unsafe.Pointer {
	return unsafe.Pointer(c.hChan)
}

func (c Chan) Len() int {
	if c.hChan == nil {
		return 0
	}

	c.lock()
	defer c.unlock()
	return int(c.hChan.qcount)
}

func (c Chan) Cap() int {
	if c.hChan == nil {
		return 0
	}

	c.lock()
	defer c.unlock()
	return int(c.hChan.dataqsiz)
}

func (c Chan) Empty() bool {
	if c.hChan == nil {
		return true
	}

	return empty(unsafe.Pointer(c.hChan))
}

func (c Chan) lock() {
	if !c.sync {
		return
	}
	lock(unsafe.Pointer(&c.hChan.lock))
}

func (c Chan) unlock() {
	if !c.sync {
		return
	}
	unlock(unsafe.Pointer(&c.hChan.lock))
}

func FromPtr(ptr unsafe.Pointer, sync bool) Chan {
	hChanPtr := (**hchan)(ptr)
	return Chan{hChan: *hChanPtr, sync: sync}
}

//func ValueOf[T any](ch chan T) Chan {
//	hChan := (*hchan)(*(*unsafe.Pointer)(unsafe.Pointer(&ch)))
//	return Chan{hChan: hChan}
//}
