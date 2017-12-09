package glue

import (
	"ir-blaster.com/ir-blaster/internal/sources/basichtml"
	"ir-blaster.com/ir-blaster/internal/sources/basichtml/multi"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func buildBasicHTMLMultiSource(v *viper.Viper, flags *pflag.FlagSet) (basichtml.Source, error) {
	srcs, err := buildBasicHTMLSources(v, flags)
	if err != nil {
		return nil, err
	}

	return baischtmlmulti.New(srcs)
}

func buildBasicHTMLSources(v *viper.Viper, flags *pflag.FlagSet) ([]basichtml.Source, error) {
	builders := []func(v *viper.Viper, flags *pflag.FlagSet) (basichtml.Source, error){
		buildBasicHTMLCheckstyleSource,
		buildBasicHTMLGotestSource,
		buildBasicHTMLGocoverSource,
	}
	srcs := make([]basichtml.Source, len(builders))
	for i, builder := range builders {
		var err error
		srcs[i], err = builder(v, flags)
		if err != nil {
			return nil, err
		}
	}

	return srcs, nil
}
