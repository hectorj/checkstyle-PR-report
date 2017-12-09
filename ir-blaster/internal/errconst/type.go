package errconst

type Error string

var _ error = Error("")

func (err Error) Error() string {
	return string(err)
}
