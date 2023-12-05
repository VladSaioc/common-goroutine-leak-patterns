package sendleaks

import (
	"testing"

	"go.uber.org/goleak"

	"github.com/VladSaioc/common-goroutine-leak-patterns/utils"
)

func TestNCastLeakEmpty_HasLeak(t *testing.T) {
	defer utils.GoroutineLeakTest(t)
	go NCastLeak(nil)
}

func TestNCastLeakManyItems_HasLeak(t *testing.T) {
	defer utils.GoroutineLeakTest(t)
	NCastLeak([]any{1, 2, 3})
}

func TestFixedNCastLeakManyItems_NoLeak(t *testing.T) {
	defer goleak.VerifyNone(t)
	FixedNCastLeak(nil)
	FixedNCastLeak([]any{1})
	FixedNCastLeak([]any{1, 2, 3})
}
