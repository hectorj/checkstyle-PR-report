package gocoverraw_test

import (
	"bytes"
	"testing"

	"io"

	"ir-blaster.com/_testdata"
	"ir-blaster.com/sources/gocover"
	"ir-blaster.com/sources/gocover/raw"
	"ir-blaster.com/sources/raw"
	"github.com/stretchr/testify/require"
)

func TestParser(t *testing.T) {
	r := require.New(t)

	reader := bytes.NewReader(_testdata.TestDataBox.Bytes("gocover.txt"))

	parser, err := gocoverraw.New(raw.SourceFunc(func() (io.Reader, error) {
		return reader, nil
	}))
	r.Nil(err)

	results, err := parser.ProvideGoCoverResults()
	r.Nil(err)

	r.Equal(gocover.Coverage{
		StatementsCovered: 18,
		Statements:        20,
		PercentageCovered: "90",
	}, *results)
}
