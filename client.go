package jira

import "context"

type Client interface {
	SearchIssues(ctx context.Context, search *ParamSearchIssues) (*Issues, error)
}
