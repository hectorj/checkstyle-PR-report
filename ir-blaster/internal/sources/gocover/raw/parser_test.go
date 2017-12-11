package gocoverraw_test

import (
	"bytes"
	"testing"

	"io"

	"github.com/stretchr/testify/require"
	"ir-blaster.com/ir-blaster/internal/_testdata"
	"ir-blaster.com/ir-blaster/internal/sources/gocover"
	"ir-blaster.com/ir-blaster/internal/sources/gocover/raw"
	"ir-blaster.com/ir-blaster/internal/sources/raw"
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
