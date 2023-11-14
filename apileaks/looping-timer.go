package apileaks

import (
	"context"
	"math/rand"
	"time"
)

// Some task to perform in the background.
func LogMetrics() {}

// StateReporter periodically performs some task (logging/monitoring)
func StateReporter() {
	period := time.Duration(rand.Intn(1000) * 1000 * 1000)
	go func() {
		// This for loop has no exit clause.
		// This is not a problem if only a limited number of state reporters are created throughout a program.
		// However, whenever StateReporter is a transitive dependency to some other package, clients may invoke
		// StateReporter without realizing any number of times, creating slowdown in the system.
		for {
			<-time.After(period)
			LogMetrics()
		}
	}()
}

// No simple idiomatic fix presents itself. Either the goroutine created by StateReporter should have
// its lifecycle tied to a context or structure, or the clients must ensure they only invoke StateReport
// a limited number of times
func FixedWithContextStateReporter(ctx context.Context) (context.Context, context.CancelFunc) {
	reporterContext, cancel := context.WithCancel(ctx)
	period := time.Duration(rand.Intn(1000) * 1000 * 1000)
	go func() {
		for {
			select {
			case <-time.After(period):
				LogMetrics()
			case <-ctx.Done():
			}
		}
	}()

	// Context and cancellation function exposed to client
	return reporterContext, cancel
}
