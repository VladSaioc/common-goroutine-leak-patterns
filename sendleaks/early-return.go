package sendleaks

// EarlyReturn demonstrates a common pattern of goroutine leaks.
// A return statement interrupts the evaluation of the parent goroutine before it can consume a message.
// Incoming error simulates an error produced internally.
func EarlyReturn(err error) {
	// Create a synchronous channel
	ch := make(chan any)

	go func() {
		// Send something to the channel.
		// Deadlocks if the parent goroutine terminates early.
		ch <- struct{}{}
	}()

	if err != nil {
		// Interrupt evaluation of parent early in case of error.
		// Sender deadlocks.
		return
	}

	// Only receive if there is no error.
	<-ch
}

// FixedEarlyReturn demonstrates how to avoid the leak.
// A return statement interrupts the evaluation of the parent goroutine before it can consume a message.
// However, the send operation unblocks because the channel capacity is large enough.
// Incoming error simulates an error produced internally.
func FixedEarlyReturn(err error) {
	// Create a synchronous channel
	ch := make(chan any, 1)

	go func() {
		// Send something to the channel.
		// Does not deadlock, as the channel can send one message without blocking.
		ch <- struct{}{}
	}()

	if err != nil {
		// Interrupt evaluation of parent early in case of error.
		// Sender does not deadlock, because sending one item is non-blocking.
		return
	}

	// Only receive if there is no error.
	<-ch
}
