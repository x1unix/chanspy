package chanspy

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsClosed(t *testing.T) {
	ch := make(chan int)
	require.False(t, IsClosed(ch))
	close(ch)
	require.True(t, IsClosed(ch))
}

func TestIsClosedConcurrent(t *testing.T) {
	ch := make(chan int)

	wait := make(chan struct{})
	go func() {
		require.False(t, IsClosed(ch))
		close(ch)
		wait <- struct{}{}
	}()

	<-wait
	require.True(t, IsClosed(ch))
}

func TestChan(t *testing.T) {
	cases := map[string]struct {
		dontClose bool

		newChan func() chan int
		prepare func(r *require.Assertions, ch chan int, c Chan)
		check   func(r *require.Assertions, c Chan)
	}{
		"nil": {
			newChan: func() chan int {
				return nil
			},
			check: func(r *require.Assertions, c Chan) {
				r.True(c.Closed())
				r.True(c.Empty())
				r.Equal(0, c.Len())
				r.Equal(0, c.Cap())
			},
		},
		"closed": {
			dontClose: true,
			prepare: func(r *require.Assertions, ch chan int, c Chan) {
				close(ch)
			},
			check: func(r *require.Assertions, c Chan) {
				r.True(c.Closed())
			},
		},
		"locked": {
			newChan: func() chan int {
				ch := make(chan int, 1)

				// Channel lock is initialized on initial write.
				ch <- 42
				return ch
			},
			check: func(r *require.Assertions, c Chan) {
				r.True(c.Locked())
			},
		},
		"unlocked": {
			check: func(r *require.Assertions, c Chan) {
				r.False(c.Locked())
			},
		},
		"buffered": {
			newChan: func() chan int {
				return make(chan int, 10)
			},
			prepare: func(r *require.Assertions, ch chan int, c Chan) {
				for i := 0; i < 5; i++ {
					ch <- i
				}
			},
			check: func(r *require.Assertions, c Chan) {
				r.Equal(5, c.Len())
				r.Equal(10, c.Cap())
				r.False(c.Empty())
				r.False(c.Closed())
			},
		},
	}

	for n, c := range cases {
		t.Run(n, func(t *testing.T) {
			r := require.New(t)

			var ch chan int
			if c.newChan != nil {
				ch = c.newChan()
			} else {
				ch = make(chan int)
			}

			view := ValueOf(ch, WithLock)
			if !c.dontClose && ch != nil {
				defer close(ch)
			}

			wait := make(chan struct{})
			defer close(wait)

			go func() {
				if c.prepare != nil {
					c.prepare(r, ch, view)
				}
				wait <- struct{}{}
			}()

			<-wait
			c.check(r, view)
		})
	}
}
