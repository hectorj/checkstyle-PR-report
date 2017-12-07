package gocover

type Source interface {
	ProvideGoCoverResults() (*Coverage, error)
}

type SourceFunc func() (*Coverage, error)

var _ Source = SourceFunc(nil)

func (fn SourceFunc) ProvideGoCoverResults() (*Coverage, error) {
	return fn()
}
