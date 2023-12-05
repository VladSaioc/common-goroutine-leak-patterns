package receiveleaks

import (
	"testing"

	"github.com/VladSaioc/common-goroutine-leak-patterns/utils"
	"go.uber.org/goleak"
)

func TestNoCloseRange0Workers_HasLeak(t *testing.T) {
	defer utils.GoroutineLeakTest(t)
	go NoCloseRange([]any{1, 2, 3}, 0)
}

func TestNoCloseRange_HasLeak(t *testing.T) {
	defer utils.GoroutineLeakTest(t)
	go NoCloseRange([]any{1, 2, 3}, 2)
}

func TestFixedNoCloseRange_NoLeak(t *testing.T) {
	defer goleak.VerifyNone(t)
	go FixedNoCloseRange([]any{1, 2, 3}, 2)
}
