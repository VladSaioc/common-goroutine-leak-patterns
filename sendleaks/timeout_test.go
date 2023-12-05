package sendleaks

import (
	"context"
	"testing"

	"go.uber.org/goleak"

	"github.com/VladSaioc/common-goroutine-leak-patterns/utils"
)

func TestTimeout_HasLeak(t *testing.T) {
	defer utils.GoroutineLeakTest(t)
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	// May only fail occasionally due to select non-determinism.
	Timeout(ctx)
}

func TestFixedTimeout_NoLeak(t *testing.T) {
	defer goleak.VerifyNone(t)
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	FixedTimeout(ctx)
}
