package glue

import "github.com/spf13/viper"

type rawConfig struct {
	v       *viper.Viper
	flagkey string
}

func (cfg rawConfig) GetFilePath() string {
	return cfg.v.GetString(cfg.flagkey)
}
