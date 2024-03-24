package chantype

import "unsafe"

//go:linkname lock runtime.lock
func lock(mu unsafe.Pointer)

//go:linkname unlock runtime.unlock
func unlock(mu unsafe.Pointer)

//go:linkname empty runtime.empty
func empty(c unsafe.Pointer) bool
