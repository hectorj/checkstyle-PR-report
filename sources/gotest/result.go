package gotest

type AllResults struct {
	Results
	ByPackage map[string]PackageResults
}

type PackageResults struct {
	Results
	CoveragePercentage string
}

type Results struct {
	Tests               []Test
	PassedTestsCount    uint
	NotPassedTestsCount uint
	SkippedTestsCount   uint
	FailedTestsCount    uint
	TotalTime           uint
}

type Test struct {
	Package string
	Name    string
	Time    uint
	Passed  bool
	Failed  bool
	Skipped bool
	Output  []string
}
