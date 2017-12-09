package basichtmlmulti

import (
	"io"
	"strings"

	"ir-blaster.com/ir-blaster/internal/sources/basichtml"
	"github.com/pkg/errors"
)

func New(sources []basichtml.Source) (basichtml.Source, error) {
	if len(sources) == 0 {
		return nil, errors.Wrap(RequiresAtLeastOneSource, "failed to create a basichtml multi source")
	}

	return basichtml.SourceFunc(func() (io.Reader, error) {
		readers := make([]io.Reader, 0, (2*len(sources))-1)
		for _, src := range sources {
			reader, err := src.ProvideBasicHTML()
			if err != nil {
				return nil, errors.Wrap(err, "one of the sources given failed to provide its reader")
			}
			if reader == nil {
				continue
			}

			if len(readers) > 0 {
				readers = append(readers, strings.NewReader("<hr/>"))
			}

			readers = append(readers, reader)
		}

		if len(readers) == 0 {
			return nil, nil
		}

		return io.MultiReader(readers...), nil
	}), nil
}
