package checkstyleraw

import (
	"ir-blaster.com/ir-blaster/internal/sources/checkstyle"
	"ir-blaster.com/ir-blaster/internal/sources/raw"
)

func New(src raw.Source, silentSuccess bool) (checkstyle.Source, error) {
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

		results, err := mapParsedXmlToViolations(parsedXml)
		if err != nil {
			return nil, err
		}

		if silentSuccess && len(results.Violations) == 0 {
			return nil, nil
		}

		return results, nil
	}), nil
}
