package github

import (
	"context"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/google/go-github/github"
	"ir-blaster.com/ir-blaster/internal/targets/basichtml"
	"ir-blaster.com/ir-blaster/internal/targets/basichtml/minified"
)

type target struct {
	httpClient *http.Client
	cfg        Config
}

func (t target) ProcessBasicHTML(r io.Reader) error {
	httpClient, err := t.cfg.GetAuth().authenticateClient(t.httpClient)
	if err != nil {
		return err
	}
	client := github.NewClient(httpClient)

	bodyBytes, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}
	bodyStr := string(bodyBytes)

	comment := &github.IssueComment{
		Body: &bodyStr,
	}

	_, _, err = client.Issues.CreateComment(context.Background(), t.cfg.GetRepoOwner(), t.cfg.GetRepoName(), int(t.cfg.GetPullRequestID()), comment)
	return err
}

func New(cfg Config, client *http.Client) (basichtml.Target, error) {
	if client == nil {
		client = http.DefaultClient
	}

	s := target{
		httpClient: client,
		cfg:        cfg,
	}

	return minified.New(s)
}
