package checkstyle

type Source interface {
	ProvideCheckstyleResults() (*AllResults, error)
}

type SourceFunc func() (*AllResults, error)

var _ Source = SourceFunc(nil)

func (fn SourceFunc) ProvideCheckstyleResults() (*AllResults, error) {
	return fn()
}
