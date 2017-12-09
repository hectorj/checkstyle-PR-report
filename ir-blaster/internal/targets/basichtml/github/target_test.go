package github_test

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/dnaeon/go-vcr/cassette"
	"github.com/dnaeon/go-vcr/recorder"
	"ir-blaster.com/ir-blaster/internal/targets/basichtml/github"
)

const testRepoOwner = "hectorj"
const testRepoName = "ir-blaster"
const testPRID = 1
const testGithubTokenEnvKey = "GITHUB_TOKEN"

var githubOauthToken = os.Getenv(testGithubTokenEnvKey)

func createHTTPClientWithVCR(t *testing.T) (c *http.Client, stopRecording func() error) {
	_, filename, _, _ := runtime.Caller(1)
	cassettePath := filepath.Join(filepath.Dir(filename), "_testdata", "VCR", t.Name())

	r, err := recorder.New(cassettePath)
	if err != nil {
		t.Fatal(err)
	}
	r.SetMatcher(func(r *http.Request, i cassette.Request) bool {
		if !cassette.DefaultMatcher(r, i) {
			return false
		}

		// match requests' bodies
		if r.Body != nil {
			b := bytes.NewBuffer(nil)
			if _, err := b.ReadFrom(r.Body); err != nil {
				return false
			}
			r.Body = ioutil.NopCloser(b)

			return b.String() == i.Body
		}

		return i.Body == ""
	})

	stop := func() error {
		err := r.Stop()
		if err != nil {
			return err
		}

		if githubOauthToken == "" {
			return nil
		}

		// remove Github Oauth token from the cassette
		k7, err := ioutil.ReadFile(cassettePath + ".yaml")
		if err != nil {
			return err
		}

		return ioutil.WriteFile(cassettePath+".yaml", bytes.Replace(k7, []byte(githubOauthToken), []byte("censoredGithubOauthToken"), -1), 0666)
	}

	return &http.Client{
		Transport: r,
	}, stop
}

func TestTarget_DummyComment(t *testing.T) {
	c, stop := createHTTPClientWithVCR(t)
	defer stop()

	cfg := github.ConfigStatic{
		Auth:          github.NewOauth(githubOauthToken),
		RepoOwner:     testRepoOwner,
		RepoName:      testRepoName,
		PullRequestID: testPRID,
	}
	target, err := github.New(cfg, c)
	if err != nil {
		t.Fatal(err)
	}

	err = target.ProcessBasicHTML(bytes.NewBufferString("just a dummy test comment"))
	if err != nil {
		t.Fatal(err)
	}
}
