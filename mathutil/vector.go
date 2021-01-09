package mathutil

//IVector interface for all types of vectors
type IVector interface {
	//Add a different vector onto this one
	Add(IVector)
	//Sub substracts a different vector from this one
	Sub(IVector)
	//Scl multiplies this vectors components with a given scalar
	Scl(float32)
	//Equals returns true if another vector equals this one
	Equals(IVector) bool
	//Print the vector
	Print(IVector)
}
