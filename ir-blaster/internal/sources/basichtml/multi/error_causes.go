package basichtmlmulti


type constantError string

func (err constantError) Error() string {
	return string(err)
}

const RequiresAtLeastOneSource = constantError("at least one source is required as input, none provided")

