package perfect

import (
	"fmt"
	"testing"
)

func assert(t *testing.T, state bool) {
	if !state {
		t.Fail()
	}
}
func guardUT(t *testing.T) {
	if err := recover(); err != nil {
		t.Fail()
	}
}

func Test_Nothing(t *testing.T) {
	fmt.Println(len(primes))
}
