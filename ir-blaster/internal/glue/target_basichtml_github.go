package glue

import (
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"ir-blaster.com/ir-blaster/internal/targets/basichtml"
	"ir-blaster.com/ir-blaster/internal/targets/basichtml/github"
)

func BuildGithubCmd() (*cobra.Command, error) {
	githubV := viper.New()
	githubFlags := pflag.NewFlagSet("github", pflag.PanicOnError)
	githubFunc, err := buildGithubApp(githubV, githubFlags)
	if err != nil {
		return nil, err
	}
	err = githubV.BindPFlags(githubFlags)
	if err != nil {
		return nil, err
	}

	basicHTMLFileCmd := &cobra.Command{
		Use: "github",
		RunE: func(cmd *cobra.Command, args []string) error {
			return githubFunc()
		},
	}
	basicHTMLFileCmd.Flags().AddFlagSet(githubFlags)

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
