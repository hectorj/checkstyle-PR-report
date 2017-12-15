package glue

import (
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"ir-blaster.com/ir-blaster/internal/targets/basichtml/github"
)

/* #nosec */
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
	oauthToken := cfg.v.GetString(githubOauthTokenKey)

	return github.NewOauth(oauthToken)
}

func (cfg githubConfig) GetRepoOwner() string {
	return cfg.v.GetString(githubRepoOwnerKey)
}

func (cfg githubConfig) GetRepoName() string {
	return cfg.v.GetString(githubRepoNameKey)
}

func (cfg githubConfig) GetPullRequestId() uint64 {
	return uint64(cfg.v.GetInt64(githubPRIDKey))
}
