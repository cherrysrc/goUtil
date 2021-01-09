package mathutil

import "fmt"

type Vector2f struct {
	X float32
	Y float32
}

func (vector *Vector2f) Add(other Vector2f) {
	vector.X += other.X
	vector.Y += other.Y
}

func (vector *Vector2f) Sub(other Vector2f) {
	vector.X -= other.X
	vector.Y -= other.Y
}

func (vector *Vector2f) Scl(scalar float32) {
	vector.X *= scalar
	vector.Y *= scalar
}

func (vector *Vector2f) Equals(other Vector2f) bool {
	return !(vector.X != other.X || vector.Y != other.Y)
}

func (vector *Vector2f) Print() {
	fmt.Printf("(%f, %f)\n", vector.X, vector.Y)
}
