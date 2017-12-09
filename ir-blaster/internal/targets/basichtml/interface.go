package basichtml

import "io"

type Target interface {
	ProcessBasicHTML(reader io.Reader) error
}

type TargetFunc func(reader io.Reader) error

var _ Target = TargetFunc(nil)

func (fn TargetFunc) ProcessBasicHTML(reader io.Reader) error {
	return fn(reader)
}
