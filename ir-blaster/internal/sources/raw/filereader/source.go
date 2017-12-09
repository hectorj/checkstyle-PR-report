package rawfilereader

import (
	"io"

	"os"

	"ir-blaster.com/ir-blaster/internal/sources/raw"
)

func New(cfg Config) (raw.Source, error) {
	return raw.SourceFunc(func() (io.Reader, error) {
		filepath := cfg.GetFilePath()
		if filepath == "" {
			return nil, nil
		}
		return os.Open(filepath)
	}), nil
}
