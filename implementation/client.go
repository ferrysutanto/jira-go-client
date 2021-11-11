package implementation

import (
	"context"
	"net/http"

	"github.com/ferrysutanto/jira-go-client/v3"
)

type clientImpl struct {
	apiHost  string
	apiToken string
	client   *http.Client
}

type ClientOption struct {
	APIHost  string
	APIToken string
}

func New(ctx context.Context, option *ClientOption) (jira.Client, error) {
	resp := clientImpl{
		apiHost:  option.APIHost,
		apiToken: option.APIToken,
		client:   &http.Client{},
	}

	return &resp, nil
}
