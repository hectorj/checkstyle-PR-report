package glue

import (
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"ir-blaster.com/ir-blaster/internal/sources/basichtml"
	"ir-blaster.com/ir-blaster/internal/sources/basichtml/gocover"
	"ir-blaster.com/ir-blaster/internal/sources/gocover"
	"ir-blaster.com/ir-blaster/internal/sources/gocover/raw"
	"ir-blaster.com/ir-blaster/internal/sources/raw"
	"ir-blaster.com/ir-blaster/internal/sources/raw/filereader"
)

func buildBasicHTMLGocoverSource(v *viper.Viper, flags *pflag.FlagSet) (basichtml.Source, error) {
	gocoverSource, err := buildGocoverSource(v, flags)
	if err != nil {
		return nil, err
	}

	return basichtmlgocover.New(gocoverSource)
}

func buildGocoverSource(v *viper.Viper, flags *pflag.FlagSet) (gocover.Source, error) {
	rawSource, err := buildRawSourceForGocover(v, flags)
	if err != nil {
		return nil, err
	}

	return gocoverraw.New(rawSource)
}

func buildRawSourceForGocover(v *viper.Viper, flags *pflag.FlagSet) (raw.Source, error) {
	flagkey := "gocover"
	flags.String(flagkey, "", "Go(lang) cover results' source, like a filepath for example")

	return rawfilereader.New(rawConfig{v: v, flagkey: flagkey})
}
