package glue

import (
	"ir-blaster.com/sources/basichtml"
	"ir-blaster.com/sources/basichtml/checkstyle"
	"ir-blaster.com/sources/checkstyle"
	"ir-blaster.com/sources/checkstyle/raw"
	"ir-blaster.com/sources/raw"
	"ir-blaster.com/sources/raw/filereader"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func buildBasicHTMLCheckstyleSource(v *viper.Viper, flags *pflag.FlagSet) (basichtml.Source, error) {
	checkstyleSource, err := buildCheckstyleSource(v, flags)
	if err != nil {
		return nil, err
	}

	return basichtmlcheckstyle.New(checkstyleSource)
}

func buildCheckstyleSource(v *viper.Viper, flags *pflag.FlagSet) (checkstyle.Source, error) {
	rawSource, err := buildRawSourceForCheckstyle(v, flags)
	if err != nil {
		return nil, err
	}

	return checkstyleraw.New(rawSource)
}

func buildRawSourceForCheckstyle(v *viper.Viper, flags *pflag.FlagSet) (raw.Source, error) {
	flagkey := "checkstyle"
	flags.String(flagkey, "", "checkstyle XML source, like a filepath for example")

	return rawfilereader.New(rawConfig{v: v, flagkey: flagkey})
}
