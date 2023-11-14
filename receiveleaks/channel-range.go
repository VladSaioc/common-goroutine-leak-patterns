package receiveleaks

// Incoming list of items and the number of workers.
func NoCloseRange(list []any, workers int) {
	ch := make(chan any)

	// Create each worker
	for i := 0; i < workers; i++ {
		go func() {
			// Each worker waits for an item and processes it.
			for item := range ch {
				// Process each item
				_ = item
			}
		}()
	}

	// Send each item to one of the workers.
	for _, item := range list {
		// Sending can deadlock if workers == 0 or if one of the workers panics
		ch <- item
	}
	// The channel is never closed, so workers deadlock once there are no more
	// items left to process.
}

// Incoming list of items and the number of workers.
func FixedNoCloseRange(list []any, workers int) {
	// The channel can accept the require number of elements
	ch := make(chan any, len(list))

	// Create each worker (can assume workers > 0)
	for i := 0; i < workers; i++ {
		go func() {
			// Each worker waits for an item and processes it.
			for item := range ch {
				// Process each item
				_ = item
			}
		}()
	}

	// Send each item to one of the workers.
	for _, item := range list {
		// Sending no longer deadlocks, even if no workers are present
		ch <- item
	}
	// Close the channel once all items are sent.
	// This allows all workers to exit their range loop and terminate
	close(ch)
}
