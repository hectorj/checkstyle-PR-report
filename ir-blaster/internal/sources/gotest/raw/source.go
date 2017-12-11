package gotestraw

import (
	"github.com/jstemmer/go-junit-report/parser"
	"ir-blaster.com/ir-blaster/internal/sources/gotest"
	"ir-blaster.com/ir-blaster/internal/sources/raw"
)

func New(src raw.Source) (gotest.Source, error) {
	return gotest.SourceFunc(func() (*gotest.AllResults, error) {
		data, err := src.ProvideRawReader()
		if err != nil {
			return nil, err
		}
		if data == nil {
			return nil, nil
		}

		report, err := parser.Parse(data, "")
		if err != nil {
			return nil, err
		}

		return mapTestReportToResults(report)
	}), nil
}
