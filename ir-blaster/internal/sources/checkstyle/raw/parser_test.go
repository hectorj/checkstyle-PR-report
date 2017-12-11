package checkstyleraw_test

// @FIXME
import (
	"bytes"
	"testing"

	"encoding/json"
	"io"

	"github.com/gobuffalo/packr"
	"github.com/stretchr/testify/require"
	"ir-blaster.com/ir-blaster/internal/_testdata"
	"ir-blaster.com/ir-blaster/internal/sources/checkstyle/raw"
	"ir-blaster.com/ir-blaster/internal/sources/raw"
)

func TestParser(t *testing.T) {
	goldenBox := packr.NewBox("./_testdata")
	type testCaseData struct {
		XML              []byte
		GoldenRecordPath string
	}
	testCases := map[string]testCaseData{
		"Empty": {
			XML:              _testdata.TestDataBox.Bytes("checkstyle_empty.xml"),
			GoldenRecordPath: "checkstyle_empty.json",
		},
		"Gometalinter": {
			XML:              _testdata.TestDataBox.Bytes("checkstyle_gometalinter.xml"),
			GoldenRecordPath: "checkstyle_gometalinter.json",
		},
	}

	for caseName, testCase := range testCases {
		t.Run(caseName, func(t *testing.T) {
			r := require.New(t)
			reader := bytes.NewReader(testCase.XML)

			parser, err := checkstyleraw.New(raw.SourceFunc(func() (io.Reader, error) {
				return reader, nil
			}))
			r.Nil(err)

			results, err := parser.ProvideCheckstyleResults()
			r.Nil(err)

			expected := goldenBox.String(testCase.GoldenRecordPath)

			actualBytes, err := json.MarshalIndent(results, "", "\t")
			r.Nil(err)

			//ioutil.WriteFile("_testdata/"+testCase.GoldenRecordPath, actualBytes, 0777)

			r.Equal(expected, string(actualBytes))
		})
	}
}
