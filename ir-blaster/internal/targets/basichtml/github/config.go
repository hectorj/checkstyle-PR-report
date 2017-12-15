package github

type Config interface {
	GetAuth() Authentication
	GetRepoOwner() string
	GetRepoName() string
	GetPullRequestId() uint64
}

type ConfigStatic struct {
	Auth          Authentication
	RepoOwner     string
	RepoName      string
	PullRequestID uint64
}

var _ Config = ConfigStatic{}

func (cfg ConfigStatic) GetAuth() Authentication {
	return cfg.Auth
}

func (cfg ConfigStatic) GetRepoOwner() string {
	return cfg.RepoOwner
}

func (cfg ConfigStatic) GetRepoName() string {
	return cfg.RepoName
}

func (cfg ConfigStatic) GetPullRequestId() uint64 {
	return cfg.PullRequestID
}
