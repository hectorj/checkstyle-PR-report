package checkstyle

type AllResults struct {
	Results
	ByFile map[string]Results
}

type Results struct {
	Violations []Violation
}

func (vs Results) Count() int {
	return len(vs.Violations)
}

type Violation struct {
	File     string
	Source   string
	Line     uint
	Column   uint
	Severity string
	Message  string
}
