package checkstyleraw

import (
	"encoding/xml"
	"io"

	"ir-blaster.com/ir-blaster/internal/sources/checkstyle"
	"github.com/pkg/errors"
)

func parseCheckstyleXML(xmlReader io.Reader) (parsedXML result, err error) {
	if err := xml.NewDecoder(xmlReader).Decode(&parsedXML); err != nil {
		return result{}, errors.Wrap(err, "Failed to parse result XML")
	}

	return parsedXML, nil
}

func mapParsedXmlToViolations(parsedXML result) (*checkstyle.AllResults, error) {
	vs := checkstyle.AllResults{
		Results: checkstyle.Results{
			Violations: make([]checkstyle.Violation, 0, len(parsedXML.Files)),
		},
		ByFile: make(map[string]checkstyle.Results, len(parsedXML.Files)),
	}

	for i := range parsedXML.Files {
		file := &parsedXML.Files[i]
		vs.ByFile[file.Name] = checkstyle.Results{
			Violations: make([]checkstyle.Violation, len(file.Errors)),
		}

		for i2 := range file.Errors {
			e := &file.Errors[i2]
			v := checkstyle.Violation{
				File:     file.Name,
				Source:   e.Source,
				Severity: e.Severity,
				Message:  e.Message,
				Line:     e.Line,
				Column:   e.Column,
			}

			vs.Violations = append(vs.Violations, v)
			vs.ByFile[file.Name].Violations[i2] = v
		}
	}

	return &vs, nil
}

type result struct {
	XMLName xml.Name     `xml:"checkstyle"`
	Version string       `xml:"version,attr"`
	Files   []fileResult `xml:"file"`
}

type fileResult struct {
	XMLName xml.Name    `xml:"file"`
	Name    string      `xml:"name,attr"`
	Errors  []violation `xml:"error"`
}

type violation struct {
	XMLName  xml.Name `xml:"error"`
	Source   string   `xml:"sources,attr"`
	Line     uint     `xml:"line,attr"`
	Column   uint     `xml:"column,attr"`
	Severity string   `xml:"severity,attr"`
	Message  string   `xml:"message,attr"`
}
