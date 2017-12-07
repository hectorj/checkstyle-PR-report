package glue

import (
	"ir-blaster.com/sources/basichtml"
	"ir-blaster.com/sources/basichtml/gocover"
	"ir-blaster.com/sources/gocover"
	"ir-blaster.com/sources/gocover/raw"
	"ir-blaster.com/sources/raw"
	"ir-blaster.com/sources/raw/filereader"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
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
