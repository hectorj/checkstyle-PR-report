package gotest

type Source interface {
	ProvideGoTestResults() (*AllResults, error)
}

type SourceFunc func() (*AllResults, error)

var _ Source = SourceFunc(nil)

func (fn SourceFunc) ProvideGoTestResults() (*AllResults, error) {
	return fn()
}
