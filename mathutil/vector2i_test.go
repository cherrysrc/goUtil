package mathutil_test

import (
	"math/rand"
	"testing"

	"github.com/cherrysrc/goUtil/mathutil"
)

func TestAddi(t *testing.T) {
	for i := 0; i < 1000; i++ {
		ax := int32(rand.Intn(100) - 50)
		ay := int32(rand.Intn(100) - 50)
		bx := int32(rand.Intn(100) - 50)
		by := int32(rand.Intn(100) - 50)

		va := mathutil.Vector2i{ax, ay}
		vb := mathutil.Vector2i{bx, by}

		va.Add(vb)

		if va.X != ax+bx || va.Y != ay+by {
			t.Errorf("Wanted (%d,%d) but got (%d,%d)\n", ax+bx, ay+by, va.X, va.Y)
		}
	}
}

func TestSubi(t *testing.T) {
	for i := 0; i < 1000; i++ {
		ax := int32(rand.Intn(100) - 50)
		ay := int32(rand.Intn(100) - 50)
		bx := int32(rand.Intn(100) - 50)
		by := int32(rand.Intn(100) - 50)

		va := mathutil.Vector2i{ax, ay}
		vb := mathutil.Vector2i{bx, by}

		va.Sub(vb)

		if va.X != ax-bx || va.Y != ay-by {
			t.Errorf("Wanted (%d,%d) but got (%d,%d)\n", ax-bx, ay-by, va.X, va.Y)
		}
	}
}

func TestScli(t *testing.T) {
	for i := 0; i < 1000; i++ {
		ax := int32(rand.Intn(100) - 50)
		ay := int32(rand.Intn(100) - 50)
		scl := rand.Float32()*100 - 50

		va := mathutil.Vector2i{ax, ay}

		va.Scl(scl)

		if va.X != ax*int32(scl) || va.Y != ay*int32(scl) {
			t.Errorf("Wanted (%d,%d) but got (%d,%d)\n", ax*int32(scl), ay*int32(scl), va.X, va.Y)
		}
	}
}
