package gotestraw

import (
	"github.com/jstemmer/go-junit-report/parser"
	"ir-blaster.com/ir-blaster/internal/sources/gotest"
	"ir-blaster.com/ir-blaster/internal/sources/raw"
)

func New(src raw.Source, silentSuccess bool) (gotest.Source, error) {
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

		results, err := mapTestReportToResults(report)
		if err != nil {
			return nil, err
		}

		if silentSuccess && results.NotPassedTestsCount == 0 {
			return nil, nil
		}

		return results, nil
	}), nil
}
