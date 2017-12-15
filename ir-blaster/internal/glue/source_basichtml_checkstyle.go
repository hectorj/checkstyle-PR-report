package glue

import (
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"ir-blaster.com/ir-blaster/internal/sources/basichtml"
	"ir-blaster.com/ir-blaster/internal/sources/basichtml/checkstyle"
	"ir-blaster.com/ir-blaster/internal/sources/checkstyle"
	"ir-blaster.com/ir-blaster/internal/sources/checkstyle/raw"
	"ir-blaster.com/ir-blaster/internal/sources/raw"
	"ir-blaster.com/ir-blaster/internal/sources/raw/filereader"
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

	return checkstyleraw.New(rawSource, true)
}

func buildRawSourceForCheckstyle(v *viper.Viper, flags *pflag.FlagSet) (raw.Source, error) {
	flagkey := "checkstyle"
	flags.String(flagkey, "", "checkstyle XML source, like a filepath for example")

	return rawfilereader.New(rawConfig{v: v, flagkey: flagkey})
}
