package static

import "github.com/gobuffalo/packr"

var TemplateBox = packr.NewBox("./templates")

func reallyMustString(str string, err error) string {
	if err != nil {
		panic(err)
	}

	return str
}

var CheckstyleBasicHTMLReportGoTpl = reallyMustString(TemplateBox.MustString("checkstyle/report.gohtml"))
var GoTestBasicHTMLReportGoTpl = reallyMustString(TemplateBox.MustString("go/test/report.gohtml"))
var GoCoverBasicHTMLReportGoTpl = reallyMustString(TemplateBox.MustString("go/cover/report.gohtml"))
