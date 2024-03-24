package chanspy

func init() {
	// Check if package is working as expected and panic if Go runtime is changed.
	if !sanityCheckTest() {
		panic("chanspy: unsupported Go version, please update the package.")
	}
}

func sanityCheckTest() (passed bool) {
	defer func() {
		if r := recover(); r != nil {
			passed = false
		}
	}()

	ch := make(chan struct{}, 10)
	ch <- struct{}{}
	c := ValueOf(ch)
	return c.Len() == len(ch) && c.Cap() == cap(ch) && !c.Empty() && !c.Closed()
}
