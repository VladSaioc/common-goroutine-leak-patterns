package utils

import (
	"fmt"
	"math/rand"
)

func RandomError() error {
	if rand.Int31n(2) == 0 {
		return nil
	}
	return fmt.Errorf("Random error")
}

func RandomSlice(max int32) []struct{} {
	return make([]struct{}, rand.Int31n(max))
}

func RandomBool() bool {
	return rand.Int31n(2) == 1
}

func RandomValue[T any](ts ...T) T {
	i := rand.Int31n(int32(len(ts)))
	return ts[i]
}
