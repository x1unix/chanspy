package chanspy_test

import (
	"fmt"

	"github.com/x1unix/chanspy"
)

func ExampleIsClosed() {
	ch := make(chan int)
	close(ch)
	fmt.Println(chanspy.IsClosed(ch)) // Output: true
}

func ExampleValueOf() {
	// Simple case with immediate access.
	// May cause race condition if the channel is used for read/write atm.
	ch := make(chan int, 10)
	ch <- 1
	ch <- 2
	ch <- 3
	ch <- 4
	defer close(ch)

	c := chanspy.ValueOf(ch)
	fmt.Printf(
		"Regular: IsClosed=%t Len=%d Cap=%d\n",
		c.Closed(), c.Len(), c.Cap(),
	)

	// Thread-safe access.
	// Will try to lock a channel before performing any operations.
	// Useful for concurrent access.
	done := make(chan struct{})
	go func() {
		c := chanspy.ValueOf(ch, chanspy.WithLock)
		_ = <-ch
		_ = <-ch
		fmt.Printf(
			"WithLock: IsClosed=%t Len=%d Cap=%d\n",
			c.Closed(), c.Len(), c.Cap(),
		)
		done <- struct{}{}
	}()
	<-done
	// Output:
	// Regular: IsClosed=false Len=4 Cap=10
	// WithLock: IsClosed=false Len=2 Cap=10
}
