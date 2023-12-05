package apileaks

import (
	"context"
	"testing"
	"time"

	"github.com/VladSaioc/common-goroutine-leak-patterns/utils"
	"go.uber.org/goleak"
)

func TestStateReporter_HasLeak(t *testing.T) {
	defer utils.GoroutineLeakTest(t)
	StateReporter()
}

func TestFixedStateReporter_NoLeak(t *testing.T) {
	defer goleak.VerifyNone(t)
	_, cancel := FixedWithContextStateReporter(context.Background())
	// Must cancel context here to avoid leaking a goroutine.
	cancel()
	<-time.After(1 * time.Second)
}
