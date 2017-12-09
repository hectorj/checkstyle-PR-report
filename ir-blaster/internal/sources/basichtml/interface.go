package basichtml

import "io"

type Source interface {
	ProvideBasicHTML() (io.Reader, error)
}

type SourceFunc func() (io.Reader, error)

var _ Source = SourceFunc(nil)

func (fn SourceFunc) ProvideBasicHTML() (io.Reader, error) {
	return fn()
}
