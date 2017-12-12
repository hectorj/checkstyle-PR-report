package basichtmlgotest

import (
	"bytes"
	"html/template"
	"io"

	"ir-blaster.com/ir-blaster/internal/sources/basichtml"
	"ir-blaster.com/ir-blaster/internal/sources/gotest"
	"ir-blaster.com/ir-blaster/internal/static"
)

var tpl = template.Must(template.New("main").Parse(static.GoTestBasicHTMLReportGoTpl()))

func New(src gotest.Source) (basichtml.Source, error) {
	return basichtml.SourceFunc(func() (io.Reader, error) {
		results, err := src.ProvideGoTestResults()
		if err != nil {
			return nil, err
		}
		if results == nil {
			return nil, nil
		}

		buf := bytes.NewBuffer(nil)

		err = tpl.Execute(buf, results)
		if err != nil {
			return nil, err
		}

		return buf, nil
	}), nil
}
