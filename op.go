package jq

// Op defines a single transformation to be applied to a []byte
type Op interface {
	Apply([]byte) ([]byte, error)
}

// OpFunc provides a convenient func type wrapper on Op
type OpFunc func([]byte) ([]byte, error)

// Apply executes the transformation defined by OpFunc
func (fn OpFunc) Apply(in []byte) ([]byte, error) {
	return fn(in)
}
