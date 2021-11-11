package implementation_test

import (
	"testing"
)

func Test_clientImpl_SearchIssues(t *testing.T) {
	if !testing.Short() {
		clientImpl_SearchIssues_IntegrationTest(t)
	}
}
