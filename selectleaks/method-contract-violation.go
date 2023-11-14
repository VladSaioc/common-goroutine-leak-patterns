package selectleaks

// A worker processes items pushed to `ch` one by one in the background.
// When the worker is no longer needed, it must be closed with `Stop`.
//
// Specifications:
//
//	A worker may be started any number of times, but must be stopped only once.
//		Stopping a worker multiple times will lead to a close panic.
//	Any worker that is started must eventually be stopped.
//		Failing to stop a worker results in a goroutine leak
type Worker struct {
	ch   chan any
	done chan any
}

// Start spawns a background goroutine that extracts items pushed to the queue.
func (w Worker) Start() {
	go func() {
		for {
			select {
			case <-w.ch: // Normal workflow
			case <-w.done:
				return // Shut down
			}
		}
	}()
}

func (w Worker) Stop() {
	// Allows goroutine created by Start to terminate
	close(w.done)
}

func (w Worker) AddToQueue(item any) {
	w.ch <- item
}

// Worker limited in scope by WorkerLifecycle
func WorkerLifecycle(items []any) {
	// Create a new worker
	w := Worker{
		ch:   make(chan any),
		done: make(chan any),
	}
	// Start worker
	w.Start()

	// Operate on worker
	for _, item := range items {
		w.AddToQueue(item)
	}

	// Exits without calling ’Stop’. Goroutine created by `Start` eventually deadlocks.
	return
}

// Worker limited in scope by WorkerLifecycle
func FixedWorkerLifecycle(items []any) {
	// Create a new worker
	w := Worker{
		ch:   make(chan any),
		done: make(chan any),
	}
	// Start worker
	w.Start()
	// Stop worker when it is no longer needed.
	defer w.Stop()

	// Operate on worker
	for _, item := range items {
		w.AddToQueue(item)
	}

	// Exits without calling ’Stop’. Goroutine created by `Start` eventually deadlocks.
	return
}
