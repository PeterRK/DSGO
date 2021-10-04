package utils

import (
	"testing"
)

func Assert(t *testing.T, state bool) {
	if !state {
		t.FailNow()
	}
}

func FailInPanic(t *testing.T) {
	if err := recover(); err != nil {
		t.Fail()
	}
}
