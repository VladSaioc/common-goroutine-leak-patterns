package sendleaks

import (
	"errors"
	"testing"

	"go.uber.org/goleak"

	"github.com/VladSaioc/common-goroutine-leak-patterns/utils"
)

func TestDoubleSendUnbuffered_HasLeak(t *testing.T) {
	defer utils.GoroutineLeakTest(t)
	// Unbuffered channel leaks
	ch := make(chan any)
	go DoubleSend(ch, errors.New("error"))
	<-ch
}

func TestDoubleOneUnbuffered_HasLeak(t *testing.T) {
	defer utils.GoroutineLeakTest(t)
	// Buffer-1 channel leaks
	ch := make(chan any, 1)
	go DoubleSend(ch, errors.New("error"))
}

func TestFixedDoubleSendUnbuffered_NoLeak(t *testing.T) {
	defer goleak.VerifyNone(t)
	ch := make(chan any)
	go FixedDoubleSend(ch, errors.New("error"))
	<-ch
	ch = make(chan any)
	go FixedDoubleSend(ch, nil)
	<-ch
}

func TestFixedDoubleOneUnbuffered_NoLeak(t *testing.T) {
	defer goleak.VerifyNone(t)
	ch := make(chan any, 1)
	go FixedDoubleSend(ch, errors.New("error"))
	ch = make(chan any, 1)
	go FixedDoubleSend(ch, nil)
}
