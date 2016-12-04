package jq

type Op interface {
	Apply([]byte) ([]byte, error)
}

type OpFunc func([]byte) ([]byte, error)

func (fn OpFunc) Apply(in []byte) ([]byte, error) {
	return fn(in)
}
