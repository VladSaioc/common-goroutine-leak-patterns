package sendleaks

import (
	"context"
)

// A context is provided to short-circuit evaluation.
func Timeout(ctx context.Context) {
	ch := make(chan any)

	go func() {
		// Perform some work and send it to the channel
		ch <- struct{}{}
	}()

	select {
	case <-ctx.Done(): // Context was cancelled or timed out
		// Sender is stuck
	case <-ch: // Receive message
		// Sender is released
	}
}

func FixedTimeout(ctx context.Context) {
	// One message may be sent over the channel without deadlocking.
	ch := make(chan any, 1)

	go func() {
		// Sending no longer deadlocks if the context is cancelled.
		ch <- struct{}{}
	}()

	select {
	case <-ctx.Done():
	case <-ch:
	}
}
