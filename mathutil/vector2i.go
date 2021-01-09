package mathutil

import "fmt"

//Vector2f implements IVector for float32
type Vector2i struct {
	X int32
	Y int32
}

func (vector *Vector2i) Add(other Vector2i) {
	vector.X += other.X
	vector.Y += other.Y
}

func (vector *Vector2i) Sub(other Vector2i) {
	vector.X -= other.X
	vector.Y -= other.Y
}

func (vector *Vector2i) Scl(scalar float32) {
	vector.X *= int32(scalar)
	vector.Y *= int32(scalar)
}

func (vector *Vector2i) Equals(other Vector2i) bool {
	return !(vector.X != other.X || vector.Y != other.Y)
}

func (vector *Vector2i) Print() {
	fmt.Printf("(%d, %d)\n", vector.X, vector.Y)
}
