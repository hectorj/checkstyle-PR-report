package basichtmlgotest_test

import (
	"testing"

	"io/ioutil"

	"github.com/gobuffalo/packr"
	"github.com/stretchr/testify/require"
	"ir-blaster.com/ir-blaster/internal/sources/basichtml/gotest"
	"ir-blaster.com/ir-blaster/internal/sources/gotest"
)

func TestSource(t *testing.T) {
	testdataBox := packr.NewBox("./_testdata")

	testCases := map[string]struct {
		Results          gotest.AllResults
		GoldenRecordPath string
	}{
		"empty": {
			Results:          gotest.AllResults{},
			GoldenRecordPath: "empty.html",
		},
		"one simple passed": {
			Results: gotest.AllResults{
				Results: gotest.Results{
					Tests: []gotest.Test{
						{
							Package: "whatever",
							Name:    "TestDoesNotMatter",
							Passed:  true,
						},
					},
					PassedTestsCount:    1,
					NotPassedTestsCount: 0,
					SkippedTestsCount:   0,
					FailedTestsCount:    0,
				},
			},
			GoldenRecordPath: "one_simple_passed.html",
		},
		"one simple failed": {
			Results: gotest.AllResults{
				Results: gotest.Results{
					Tests: []gotest.Test{
						{
							Package: "whatever",
							Name:    "TestDoesNotMatter",
							Passed:  false,
							Time:    42,
							Output: []string{
								"First line of output",
								"\tSecond line of output",
							},
						},
					},
					PassedTestsCount:    0,
					NotPassedTestsCount: 0,
					SkippedTestsCount:   0,
					FailedTestsCount:    1,
					TotalTime:           43,
				},
				ByPackage: map[string]gotest.PackageResults{
					"whatever": {
						Results: gotest.Results{
							Tests: []gotest.Test{
								{
									Package: "whatever",
									Name:    "TestDoesNotMatter",
									Passed:  false,
									Time:    41,
									Output: []string{
										"First line of output",
										"\tSecond line of output",
									},
								},
							},
							PassedTestsCount:    0,
							NotPassedTestsCount: 0,
							SkippedTestsCount:   0,
							FailedTestsCount:    1,
							TotalTime:           44,
						},
					},
				},
			},
			GoldenRecordPath: "one_simple_failed.html",
		},
	}

	for testName, testCase := range testCases {
		t.Run(testName, func(t *testing.T) {
			r := require.New(t)

			src, err := basichtmlgotest.New(gotest.SourceFunc(func() (*gotest.AllResults, error) {
				return &testCase.Results, nil
			}))
			r.Nil(err)
			r.NotNil(src)

			reader, err := src.ProvideBasicHTML()
			r.Nil(err)
			r.NotNil(reader)

			actualBytes, err := ioutil.ReadAll(reader)
			r.Nil(err)

			expected, err := testdataBox.MustString(testCase.GoldenRecordPath)
			r.Nil(err)

			//if expected != string(actualBytes) {
			//	ioutil.WriteFile("_testdata/"+testCase.GoldenRecordPath, actualBytes, 0777)
			//}

			r.Equal(expected, string(actualBytes))
		})
	}

}
