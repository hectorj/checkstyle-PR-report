package static

import (
	"github.com/gobuffalo/packr"
)

var TemplateBox = packr.NewBox("templates")

func reallyMustString(str string, err error) string {
	if err != nil {
		panic(err)
	}

	return str
}

func CheckstyleBasicHTMLReportGoTpl() string {
	return reallyMustString(TemplateBox.MustString("checkstyle/report.gohtml"))
}

func GoTestBasicHTMLReportGoTpl() string {
	return reallyMustString(TemplateBox.MustString("go/test/report.gohtml"))
}

func GoCoverBasicHTMLReportGoTpl() string {
	return reallyMustString(TemplateBox.MustString("go/cover/report.gohtml"))
}
