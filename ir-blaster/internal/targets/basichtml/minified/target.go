package minified

import (
	"io"

	"github.com/tdewolff/minify"
	"github.com/tdewolff/minify/html"
	"ir-blaster.com/ir-blaster/internal/targets/basichtml"
)

func New(wrapped basichtml.Target) (basichtml.Target, error) {
	return basichtml.TargetFunc(func(reader io.Reader) error {
		m := minify.New()

		m.Add("text/html", &html.Minifier{
			KeepDefaultAttrVals: true,
			KeepEndTags:         true,
		})
		miniReader := m.Reader("text/html", reader)

		return wrapped.ProcessBasicHTML(miniReader)
	}), nil
}
