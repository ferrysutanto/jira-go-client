package implementation_test

import (
	"context"
	"net/http"
	"os"
	"testing"

	"github.com/ferrysutanto/jira-go-client/v3"
	"github.com/ferrysutanto/jira-go-client/v3/implementation"
	"github.com/joho/godotenv"
)

func clientImpl_SearchIssues_IntegrationTest(t *testing.T) {
	err := godotenv.Load()
	if err != nil {
		t.Skipf(".env file not provided")
		return
	}

	ctx := context.Background()

	apiHost := os.Getenv("API_HOST")
	apiToken := os.Getenv("API_TOKEN")
	project := ost.Getenv("PROJECT")

	opt := implementation.ClientOption{
		APIHost:  apiHost,
		APIToken: apiToken,
	}

	client, err := implementation.New(ctx, &opt)
	if err != nil {
		t.Fatalf("[Error] Instantiating JIRA client. Reason: %v", err)
	}

	type fields struct {
		apiHost string
		token   string
		client  *http.Client
	}
	type args struct {
		ctx   context.Context
		param *jira.ParamSearchIssues
	}
	tests := []struct {
		name    string
		args    args
		want    *jira.Issues
		wantErr bool
	}{
		{
			"success",
			args{
				ctx: context.Background(),
				param: &jira.ParamSearchIssues{
					Project:  func() *string { resp := project; return &resp }(),
					OrderBy:  func() *string { resp := "created"; return &resp }(),
					OrderDir: func() *jira.OrderDir { resp := jira.OrderDesc; return &resp }(),
				},
			},
			nil,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := client.SearchIssues(tt.args.ctx, tt.args.param)
			if (err != nil) != tt.wantErr {
				t.Errorf("clientImpl.SearchIssues() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			_ = got
		})
	}
}
