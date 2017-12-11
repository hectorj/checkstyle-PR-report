package filewriter

import (
	"io"
	"os"

	"strings"

	"github.com/pkg/errors"
	"ir-blaster.com/ir-blaster/internal/targets/basichtml"
	"ir-blaster.com/ir-blaster/internal/targets/basichtml/minified"
)

func New(cfg Config) (basichtml.Target, error) {
	return minified.New(basichtml.TargetFunc(func(reader io.Reader) error {
		filePath := cfg.GetFilePath()
		if filePath == "" {
			return errors.New("missing basic html output filepath")
		}
		if !strings.HasSuffix(filePath, ".html") {
			filePath += ".html"
		}

		file, err := os.Create(filePath)
		if err != nil {
			return err
		}

		_, err = io.Copy(file, reader)
		if err != nil {
			return err
		}

		return file.Close()
	}))
}
