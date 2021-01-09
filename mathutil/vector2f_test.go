package mathutil_test

import (
	"math/rand"
	"testing"

	"github.com/cherrysrc/goUtil/mathutil"
)

func TestAddf(t *testing.T) {
	for i := 0; i < 1000; i++ {
		ax := rand.Float32()*100 - 50
		ay := rand.Float32()*100 - 50
		bx := rand.Float32()*100 - 50
		by := rand.Float32()*100 - 50

		va := mathutil.Vector2f{ax, ay}
		vb := mathutil.Vector2f{bx, by}

		va.Add(vb)

		if va.X != ax+bx || va.Y != ay+by {
			t.Errorf("Wanted (%f,%f) but got (%f,%f)\n", ax+bx, ay+by, va.X, va.Y)
		}
	}
}

func TestSubf(t *testing.T) {
	for i := 0; i < 1000; i++ {
		ax := rand.Float32()*100 - 50
		ay := rand.Float32()*100 - 50
		bx := rand.Float32()*100 - 50
		by := rand.Float32()*100 - 50

		va := mathutil.Vector2f{ax, ay}
		vb := mathutil.Vector2f{bx, by}

		va.Sub(vb)

		if va.X != ax-bx || va.Y != ay-by {
			t.Errorf("Wanted (%f,%f) but got (%f,%f)\n", ax-bx, ay-by, va.X, va.Y)
		}
	}
}

func TestSclf(t *testing.T) {
	for i := 0; i < 1000; i++ {
		ax := rand.Float32()*100 - 50
		ay := rand.Float32()*100 - 50
		scl := rand.Float32() * 100

		va := mathutil.Vector2f{ax, ay}

		va.Scl(scl)

		if va.X != ax*scl || va.Y != ay*scl {
			t.Errorf("Wanted (%f,%f) but got (%f,%f)\n", ax*scl, ay*scl, va.X, va.Y)
		}
	}
}
