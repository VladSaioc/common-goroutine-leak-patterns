package utils

import (
	"reflect"
	"testing"
	"unsafe"

	"go.uber.org/goleak"
)

func RecoverTestFailure(t *testing.T) {
	rPtrT := reflect.ValueOf(t)
	rT := rPtrT.Elem()
	rCommon := rT.FieldByName("common")

	var pointerOffset uintptr = 0
	for i, field := range reflect.VisibleFields(rCommon.Type()) {
		if field.Name == "failed" {
			break
		}
		pointerOffset += rCommon.Field(i).Type().Size()
	}

	ptrToT := unsafe.Pointer(t)

	*(*bool)(unsafe.Pointer(uintptr(ptrToT) + pointerOffset)) = false
}

func GoroutineLeakTest(t *testing.T) {
	defer func() {
		if !t.Failed() {
			t.Error("Test did not fail with Goroutine leak.")
		} else {
			RecoverTestFailure(t)
		}
	}()
	defer goleak.VerifyNone(t)
}
