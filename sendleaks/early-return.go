package sendleaks

import "github.com/VladSaioc/common-goroutine-leak-patterns/utils"

func EarlyReturn() {
	// Create a synchronous channel
	ch := make(chan any)

	go func() {
		// Send something to the channel.
		// Deadlocks if the parent goroutine terminates early.
		ch <- struct{}{}
	}()

	if utils.RandomError() != nil {
		// Interrupt evaluation of parent early in case of error.
		// Sender deadlocks.
		return
	}

	// Only receive if there is no error.
	<-ch
}

func FixedEarlyReturn() {
	// Create a synchronous channel
	ch := make(chan any, 1)

	go func() {
		// Send something to the channel.
		// Does not deadlock, as the channel can send one message without blocking.
		ch <- struct{}{}
	}()

	if utils.RandomError() != nil {
		// Interrupt evaluation of parent early in case of error.
		// Sender does not deadlock, because sending one item is non-blocking.
		return
	}

	// Only receive if there is no error.
	<-ch
}
