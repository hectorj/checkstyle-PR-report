package glue

import (
	"ir-blaster.com/targets/basichtml"
	"ir-blaster.com/targets/basichtml/github"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func BuildGithubCmd() (*cobra.Command, error) {
	basicHTMLFileV := viper.New()
	basicHTMLFileFlags := pflag.NewFlagSet("github", pflag.PanicOnError)
	basicHTMLFileFunc, err := buildGithubApp(basicHTMLFileV, basicHTMLFileFlags)
	if err != nil {
		return nil, err
	}
	err = basicHTMLFileV.BindPFlags(basicHTMLFileFlags)
	if err != nil {
		return nil, err
	}

	basicHTMLFileCmd := &cobra.Command{
		Use: "github",
		RunE: func(cmd *cobra.Command, args []string) error {
			return basicHTMLFileFunc()
		},
	}
	basicHTMLFileCmd.Flags().AddFlagSet(basicHTMLFileFlags)

	return basicHTMLFileCmd, nil
}

func buildGithubApp(v *viper.Viper, flags *pflag.FlagSet) (func() error, error) {
	trgt, err := buildGithubTarget(v, flags)
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

func buildGithubTarget(v *viper.Viper, flags *pflag.FlagSet) (basichtml.Target, error) {
	cfg := githubConfig{v: v}
	cfg.setupFlags(flags)

	return github.New(cfg, nil)
}
