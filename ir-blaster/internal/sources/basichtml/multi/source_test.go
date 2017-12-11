package basichtmlmulti_test

import (
	"io"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
	"ir-blaster.com/ir-blaster/internal/sources/basichtml"
	"ir-blaster.com/ir-blaster/internal/sources/basichtml/multi"
)

func TestNew(t *testing.T) {
	testcases := map[string]struct {
		srcs           []basichtml.Source
		expectedResult string
	}{
		"single": {
			srcs: []basichtml.Source{
				newHardcodedSource("<p>test data</p>"),
			},
			expectedResult: "<p>test data</p>",
		},
		"double": {
			srcs: []basichtml.Source{
				newHardcodedSource("<p>test data</p>"),
				newHardcodedSource("<p>some other test data</p>"),
			},
			expectedResult: "<p>test data</p><hr/><p>some other test data</p>",
		},
		"ignores_nil": {
			srcs: []basichtml.Source{
				newNilSource(),
				newHardcodedSource("<p>test data</p>"),
				newNilSource(),
				newHardcodedSource("<p>some other test data</p>"),
				newNilSource(),
			},
			expectedResult: "<p>test data</p><hr/><p>some other test data</p>",
		},
	}

	for testname, testcase := range testcases {
		t.Run(testname, func(t *testing.T) {
			r := require.New(t)

			src, err := basichtmlmulti.New(testcase.srcs)
			r.Nil(err)
			r.NotNil(src)

			reader, err := src.ProvideBasicHTML()
			r.Nil(err)
			r.NotNil(reader)

			actualBytes, err := ioutil.ReadAll(reader)
			r.Nil(err)

			r.Equal(testcase.expectedResult, string(actualBytes))
		})
	}
}

func TestNew_Error(t *testing.T) {
	testcases := map[string]struct {
		srcs               []basichtml.Source
		expectedErrorCause error
	}{
		"nil_sources": {
			srcs:               nil,
			expectedErrorCause: basichtmlmulti.RequiresAtLeastOneSource,
		},
		"empty_sources": {
			srcs:               make([]basichtml.Source, 0, 1),
			expectedErrorCause: basichtmlmulti.RequiresAtLeastOneSource,
		},
	}

	for testname, testcase := range testcases {
		t.Run(testname, func(t *testing.T) {
			r := require.New(t)

			src, err := basichtmlmulti.New(testcase.srcs)
			r.NotNil(err)
			r.Nil(src)

			if testcase.expectedErrorCause != nil {
				r.Equal(testcase.expectedErrorCause, errors.Cause(err))
			}
		})
	}
}

func TestNew_ErrorOnProvide(t *testing.T) {
	var testErrorCause = errors.New("test error cause")

	testcases := map[string]struct {
		srcs               []basichtml.Source
		expectedErrorCause error
	}{
		"propagates": {
			srcs: []basichtml.Source{
				newHardcodedSource("<p>test data</p>"),
				newErrorOnProvideSource(testErrorCause),
				newHardcodedSource("<p>some other test data</p>"),
			},
			expectedErrorCause: testErrorCause,
		},
	}

	for testname, testcase := range testcases {
		t.Run(testname, func(t *testing.T) {
			r := require.New(t)

			src, err := basichtmlmulti.New(testcase.srcs)
			r.Nil(err)
			r.NotNil(src)

			reader, err := src.ProvideBasicHTML()
			r.NotNil(err)
			r.Nil(reader)

			if testcase.expectedErrorCause != nil {
				r.Equal(testcase.expectedErrorCause, errors.Cause(err))
			}
		})
	}
}

func TestNew_ProvidesNil(t *testing.T) {
	testcases := map[string]struct {
		srcs []basichtml.Source
	}{
		"single_nil": {
			srcs: []basichtml.Source{
				newNilSource(),
			},
		},
		"multiple_nil": {
			srcs: []basichtml.Source{
				newNilSource(),
				newNilSource(),
				newNilSource(),
			},
		},
	}

	for testname, testcase := range testcases {
		t.Run(testname, func(t *testing.T) {
			r := require.New(t)

			src, err := basichtmlmulti.New(testcase.srcs)
			r.Nil(err)
			r.NotNil(src)

			reader, err := src.ProvideBasicHTML()
			r.Nil(err)
			r.Nil(reader)
		})
	}
}

func TestNew_ErrorOnRead(t *testing.T) {
	var testErrorCause = errors.New("test error cause")

	testcases := map[string]struct {
		srcs               []basichtml.Source
		expectedErrorCause error
	}{
		"propagates": {
			srcs: []basichtml.Source{
				newHardcodedSource("<p>test data</p>"),
				newErrorOnReadSource(testErrorCause),
				newHardcodedSource("<p>some other test data</p>"),
			},
			expectedErrorCause: testErrorCause,
		},
	}

	for testname, testcase := range testcases {
		t.Run(testname, func(t *testing.T) {
			r := require.New(t)

			src, err := basichtmlmulti.New(testcase.srcs)
			r.Nil(err)
			r.NotNil(src)

			reader, err := src.ProvideBasicHTML()
			r.Nil(err)
			r.NotNil(reader)

			_, err = ioutil.ReadAll(reader)
			r.NotNil(err)

			if testcase.expectedErrorCause != nil {
				r.Equal(testcase.expectedErrorCause, errors.Cause(err))
			}
		})
	}
}

type erroringReader struct{ err error }

func (r erroringReader) Read(p []byte) (n int, err error) {
	return 0, r.err
}

func newErrorOnReadSource(err error) basichtml.Source {
	return basichtml.SourceFunc(func() (io.Reader, error) {
		return erroringReader{err: err}, nil
	})
}

func newErrorOnProvideSource(err error) basichtml.Source {
	return basichtml.SourceFunc(func() (io.Reader, error) {
		return nil, err
	})
}

func newNilSource() basichtml.Source {
	return basichtml.SourceFunc(func() (io.Reader, error) {
		return nil, nil
	})
}

func newHardcodedSource(data string) basichtml.Source {
	return basichtml.SourceFunc(func() (io.Reader, error) {
		return strings.NewReader(data), nil
	})
}
