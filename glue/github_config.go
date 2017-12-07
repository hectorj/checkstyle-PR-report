package glue

import (
	"ir-blaster.com/targets/basichtml/github"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const (
	githubOauthTokenKey = "github-oauth-token"
	githubRepoOwnerKey  = "github-repo-owner"
	githubRepoNameKey   = "github-repo-name"
	githubPRIDKey       = "github-pr-id"
)

type githubConfig struct {
	v *viper.Viper
}

func (cfg githubConfig) setupFlags(flags *pflag.FlagSet) {
	flags.String(githubOauthTokenKey, "", "")
	flags.String(githubRepoOwnerKey, "", "")
	flags.String(githubRepoNameKey, "", "")
	flags.String(githubPRIDKey, "", "")
}

func (cfg githubConfig) GetAuth() github.Authentication {
	oauthToken := viper.GetString(githubOauthTokenKey)

	return github.NewOauth(oauthToken)
}

func (cfg githubConfig) GetRepoOwner() string {
	return viper.GetString(githubRepoOwnerKey)
}

func (cfg githubConfig) GetRepoName() string {
	return viper.GetString(githubRepoNameKey)
}

func (cfg githubConfig) GetPullRequestID() uint64 {
	return uint64(viper.GetInt64(githubPRIDKey))
}
