package checkstyleraw

import (
	"ir-blaster.com/sources/checkstyle"
	"ir-blaster.com/sources/raw"
)

func New(src raw.Source) (checkstyle.Source, error) {
	return checkstyle.SourceFunc(func() (*checkstyle.AllResults, error) {
		reader, err := src.ProvideRawReader()
		if err != nil {
			return nil, err
		}
		if reader == nil {
			return nil, nil
		}

		parsedXml, err := parseCheckstyleXML(reader)
		if err != nil {
			return nil, err
		}

		return mapParsedXmlToViolations(parsedXml)
	}), nil
}
