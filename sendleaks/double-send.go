package sendleaks

// DoubleSend incoming channel must send a message (incoming error simulates an error generated internally).
func DoubleSend(ch chan any, err error) {
	if err != nil {
		// In case of an error, send nil.
		ch <- nil
		// Return is missing here.
	}
	// Otherwise, continue with normal behaviour
	// This send is still executed in the error case, which may lead to deadlock.
	ch <- struct{}{}
}

// FixedDoubleSend incoming channel must send a message (incoming error simulates an error generated internally).
func FixedDoubleSend(ch chan any, err error) {
	if err != nil {
		ch <- nil
		return // Return interrupts control flow here.
	}
	// This send is no longer executed in the error case, avoiding a potential deadlock.
	ch <- struct{}{}
}
