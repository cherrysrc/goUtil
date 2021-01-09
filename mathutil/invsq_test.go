package mathutil_test

import (
	"math"
	"math/rand"
	"testing"

	"github.com/cherrysrc/goUtil/mathutil"
)

func TestInverseSquareRoot(t *testing.T) {
	delta := 0.005
	for i := 0; i < 1000; i++ {
		x := rand.Float32() * 100

		std := 1 / math.Sqrt(float64(x))
		fast := mathutil.InverseSquareRoot(x)

		diff := math.Abs(std - float64(fast))
		if diff > delta {
			t.Errorf("Difference exceeds delta: %f\n", diff)
		}
	}
}
