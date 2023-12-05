package sendleaks

import (
	"errors"
	"testing"

	"go.uber.org/goleak"

	"github.com/VladSaioc/common-goroutine-leak-patterns/utils"
)

func TestEarlyReturn_HasLeak(t *testing.T) {
	defer utils.GoroutineLeakTest(t)
	go EarlyReturn(errors.New("error"))
}

func TestFixedEarlyReturn_NoLeak(t *testing.T) {
	defer goleak.VerifyNone(t)
	go FixedEarlyReturn(errors.New("error"))
	go FixedEarlyReturn(nil)
}
