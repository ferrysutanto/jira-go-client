package implementation

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/ferrysutanto/jira-go-client/v3"
)

func (this *clientImpl) SearchIssues(ctx context.Context, param *jira.ParamSearchIssues) (*jira.Issues, error) {
	// Declare method for search issues with "GET"
	method := "POST"

	// Declare endpoint for search issues (host + "/search")
	ep := fmt.Sprintf("%s/%s", this.apiHost, "search")

	req, err := http.NewRequestWithContext(ctx, method, ep, nil)
	if err != nil {
		return nil, err
	}

	// Setting authentication (using basic auth b64 auth)
	headerKey := "Authorization"
	headerVal := fmt.Sprintf("Basic %s", this.apiToken)
	req.Header.Set(headerKey, headerVal)

	// Execute API call to JIRA server
	callResp, err := this.client.Do(req)
	if err != nil {
		return nil, err
	}

	// Validate returned status code
	if callResp.StatusCode != 200 {
		return nil, fmt.Errorf("Error %d", callResp.StatusCode)
	}

	// Parse to []byte array before it can get mapped to struct "Issues"
	encResp, err := ioutil.ReadAll(callResp.Body)
	if err != nil {
		return nil, err
	}

	// Declare response
	resp := jira.Issues{}

	// Map API call parsed body to response
	if err := json.Unmarshal(encResp, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
