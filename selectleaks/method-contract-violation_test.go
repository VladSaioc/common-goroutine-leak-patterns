package selectleaks

import (
	"testing"

	"github.com/VladSaioc/common-goroutine-leak-patterns/utils"
	"go.uber.org/goleak"
)

func TestWorkerLifecycle_HasLeak(t *testing.T) {
	defer utils.GoroutineLeakTest(t)
	go WorkerLifecycle([]any{1, 2, 3})
}

func TestFixedWorkerLifecycle_NoLeak(t *testing.T) {
	defer goleak.VerifyNone(t)
	go FixedWorkerLifecycle([]any{1, 2, 3})
}
