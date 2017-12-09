package github

import (
	"context"
	"net/http"

	"golang.org/x/oauth2"
)

type Authentication interface {
	authenticateClient(*http.Client) (*http.Client, error)
}

func NewOauth(token string) Authentication {
	return oauth{
		token: token,
	}
}

type oauth struct {
	token string
}

func (auth oauth) authenticateClient(c *http.Client) (*http.Client, error) {
	ctx := context.WithValue(context.Background(), oauth2.HTTPClient, c)

	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: auth.token})

	return oauth2.NewClient(ctx, ts), nil
}
