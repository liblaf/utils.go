package errors

import (
	"fmt"

	"github.com/dsnet/try"
)

func O1[T any](a T, ok bool) T {
	if ok {
		return a
	} else {
		return try.E1(a, fmt.Errorf("no"))
	}
}
