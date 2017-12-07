package gotestraw

import (
	"ir-blaster.com/sources/gotest"
	testparser "github.com/jstemmer/go-junit-report/parser"
)

func mapTestReportToResults(report *testparser.Report) (*gotest.AllResults, error) {
	result := &gotest.AllResults{
		Results: gotest.Results{
			Tests: make([]gotest.Test, 0, len(report.Packages)),
		},
		ByPackage: make(map[string]gotest.PackageResults, len(report.Packages)),
	}

	for i := range report.Packages {
		pkg := &report.Packages[i]

		pkgResult := gotest.PackageResults{
			Results: gotest.Results{
				Tests:     make([]gotest.Test, len(pkg.Tests)),
				TotalTime: uint(pkg.Time),
			},
			CoveragePercentage: pkg.CoveragePct,
		}

		for i2, testReport := range pkg.Tests {

			pkgResult.Tests[i2] = gotest.Test{
				Package: pkg.Name,
				Name:    testReport.Name,
				Time:    uint(testReport.Time),
				Output:  testReport.Output,
				Passed:  testReport.Result == testparser.PASS,
				Skipped: testReport.Result == testparser.SKIP,
				Failed:  testReport.Result == testparser.FAIL,
			}

			switch testReport.Result {
			case testparser.PASS:
				pkgResult.PassedTestsCount++
			case testparser.FAIL:
				pkgResult.FailedTestsCount++
			case testparser.SKIP:
				pkgResult.SkippedTestsCount++
			}
		}

		pkgResult.NotPassedTestsCount = pkgResult.FailedTestsCount + pkgResult.SkippedTestsCount

		result.ByPackage[pkg.Name] = pkgResult

		result.Tests = append(result.Tests, result.ByPackage[pkg.Name].Tests...)
		result.TotalTime += result.ByPackage[pkg.Name].TotalTime
		result.PassedTestsCount += result.ByPackage[pkg.Name].PassedTestsCount
		result.FailedTestsCount += result.ByPackage[pkg.Name].FailedTestsCount
		result.SkippedTestsCount += result.ByPackage[pkg.Name].SkippedTestsCount
		result.NotPassedTestsCount += result.ByPackage[pkg.Name].NotPassedTestsCount
	}

	return result, nil
}
