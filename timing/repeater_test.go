package timing_test

import (
	"testing"
	"time"

	"github.com/cherrysrc/goUtil/timing"
)

func TestRepeat(t *testing.T) {
	counter := 0

	stop := timing.Repeat(1*time.Millisecond, func() {
		if counter < 10 {
			counter++
		}
	})

	time.Sleep(11 * time.Millisecond)
	stop <- true

	if counter != 10 {
		t.Errorf("Expected counter to be 10 but it was %d\n", counter)
	}
}
