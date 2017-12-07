package glue

import (
	"ir-blaster.com/sources/basichtml"
	"ir-blaster.com/sources/basichtml/gotest"
	"ir-blaster.com/sources/gotest"
	"ir-blaster.com/sources/gotest/raw"
	"ir-blaster.com/sources/raw"
	"ir-blaster.com/sources/raw/filereader"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func buildBasicHTMLGotestSource(v *viper.Viper, flags *pflag.FlagSet) (basichtml.Source, error) {
	gotestSource, err := buildGotestSource(v, flags)
	if err != nil {
		return nil, err
	}

	return basichtmlgotest.New(gotestSource)
}

func buildGotestSource(v *viper.Viper, flags *pflag.FlagSet) (gotest.Source, error) {
	rawSource, err := buildRawSourceForGotest(v, flags)
	if err != nil {
		return nil, err
	}

	return gotestraw.New(rawSource)
}

func buildRawSourceForGotest(v *viper.Viper, flags *pflag.FlagSet) (raw.Source, error) {
	flagkey := "gotest"
	flags.String(flagkey, "", "Go(lang) test results' source, like a filepath for example")

	return rawfilereader.New(rawConfig{v: v, flagkey: flagkey})
}
