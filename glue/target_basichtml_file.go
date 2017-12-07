package glue

import (
	"ir-blaster.com/targets/basichtml"
	"ir-blaster.com/targets/basichtml/filewriter"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func BuildBasicHTMLFileCmd() (*cobra.Command, error) {
	basicHTMLFileV := viper.New()
	basicHTMLFileFlags := pflag.NewFlagSet("htmlfile", pflag.PanicOnError)
	basicHTMLFileFunc, err := buildBasicHTMLFileApp(basicHTMLFileV, basicHTMLFileFlags)
	if err != nil {
		return nil, err
	}
	err = basicHTMLFileV.BindPFlags(basicHTMLFileFlags)
	if err != nil {
		return nil, err
	}

	basicHTMLFileCmd := &cobra.Command{
		Use: "htmlfile",
		RunE: func(cmd *cobra.Command, args []string) error {
			return basicHTMLFileFunc()
		},
	}
	basicHTMLFileCmd.Flags().AddFlagSet(basicHTMLFileFlags)

	return basicHTMLFileCmd, nil
}

func buildBasicHTMLFileApp(v *viper.Viper, flags *pflag.FlagSet) (func() error, error) {
	trgt, err := buildBasicHTMLFileTarget(v, flags)
	if err != nil {
		return nil, err
	}

	src, err := buildBasicHTMLMultiSource(v, flags)
	if err != nil {
		return nil, err
	}

	return func() error {
		html, err := src.ProvideBasicHTML()
		if err != nil {
			return err
		}

		return trgt.ProcessBasicHTML(html)
	}, nil
}

func buildBasicHTMLFileTarget(v *viper.Viper, flags *pflag.FlagSet) (basichtml.Target, error) {
	flagkey := "output"
	flags.String("output", "", "output filepath. '.html' extension will be added if necessary")

	return filewriter.New(rawConfig{v: v, flagkey: flagkey})
}
