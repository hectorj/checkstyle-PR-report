package filewriter

type Config interface {
	GetFilePath() string
}

type ConfigStatic struct {
	FilePath string
}

var _ Config = ConfigStatic{}

func (cfg ConfigStatic) GetFilePath() string {
	return cfg.FilePath
}
