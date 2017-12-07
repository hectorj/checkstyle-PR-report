package basichtmlcheckstyle

import (
	"bytes"
	"html/template"
	"io"

	"ir-blaster.com/sources/basichtml"
	"ir-blaster.com/sources/checkstyle"
	"ir-blaster.com/static"
)

var tpl = template.Must(template.New("main").Parse(static.CheckstyleBasicHTMLReportGoTpl))

func New(src checkstyle.Source) (basichtml.Source, error) {
	return basichtml.SourceFunc(func() (io.Reader, error) {
		results, err := src.ProvideCheckstyleResults()
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
