package gocoverraw

import (
	"io"
	"io/ioutil"
	"os"
	"strconv"

	"ir-blaster.com/ir-blaster/internal/sources/gocover"
	"ir-blaster.com/ir-blaster/internal/sources/raw"
	"golang.org/x/tools/cover"
)

func New(src raw.Source) (gocover.Source, error) {
	return gocover.SourceFunc(func() (*gocover.Coverage, error) {
		data, err := src.ProvideRawReader()
		if err != nil {
			return nil, err
		}
		if data == nil {
			return nil, nil
		}

		profiles, err := parseGoCoverage(data)
		if err != nil {
			return nil, err
		}

		return mapCoverProfilesToCoverage(profiles)
	}), nil
}

func parseGoCoverage(data io.Reader) ([]*cover.Profile, error) {
	// create a tmp file, because golang.org/x/tools/cover::ParseProfiles only takes a filename...
	tmp, err := ioutil.TempFile("", "ir-blaster_gocover")
	if err != nil {
		return nil, err
	}
	tmpFilePath := tmp.Name()
	defer os.Remove(tmpFilePath)

	// write the data to this tmp file
	_, err = io.Copy(tmp, data)
	if err != nil {
		return nil, err
	}
	err = tmp.Close()
	if err != nil {
		return nil, err
	}

	return cover.ParseProfiles(tmpFilePath)
}

func mapCoverProfilesToCoverage(profiles []*cover.Profile) (c *gocover.Coverage, err error) {
	c = new(gocover.Coverage)
	for _, profile := range profiles {
		for _, block := range profile.Blocks {
			c.Statements += uint(block.NumStmt)
			if block.Count > 0 {
				c.StatementsCovered += uint(block.NumStmt)
			}
		}
	}

	floatPercentage := (float64(c.StatementsCovered) / float64(c.Statements)) * 100.0
	c.PercentageCovered = strconv.FormatFloat(floatPercentage, 'f', -1, 64)

	return c, nil
}
