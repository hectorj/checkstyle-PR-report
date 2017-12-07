package raw

import "io"

type Source interface {
	ProvideRawReader() (io.Reader, error)
}

type SourceFunc func() (io.Reader, error)

var _ Source = SourceFunc(nil)

func (fn SourceFunc) ProvideRawReader() (io.Reader, error) {
	return fn()
}
