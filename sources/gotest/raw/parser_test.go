package gotestraw_test

import (
	"bytes"
	"encoding/json"
	"io"
	"testing"

	"github.com/gobuffalo/packr"
	"github.com/stretchr/testify/require"
	"ir-blaster.com/_testdata"
	"ir-blaster.com/sources/gotest/raw"
	"ir-blaster.com/sources/raw"
)

func TestParser(t *testing.T) {
	r := require.New(t)

	reader := bytes.NewReader(_testdata.TestDataBox.Bytes("gotest.txt"))

	parser, err := gotestraw.New(raw.SourceFunc(func() (io.Reader, error) {
		return reader, nil
	}))
	r.Nil(err)

	results, err := parser.ProvideGoTestResults()
	r.Nil(err)

	actualBytes, err := json.MarshalIndent(results, "", "\t")
	r.Nil(err)

	goldenBytes := packr.NewBox("./_testdata").Bytes("golden_master.json")

	//r.Nil(ioutil.WriteFile("./_testdata/golden_master.json", actualBytes, 0666))
	r.Equal(string(goldenBytes), string(actualBytes))

}
