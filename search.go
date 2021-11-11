package jira

type ParamSearchIssues struct {
	JQL      *string
	Fields   []string
	Limit    *uint
	OrderBy  *string
	OrderDir *OrderDir
}

type Issues struct {
	Expand     string  `json:"expand"`
	StartAt    uint    `json:"startAt"`
	MaxResults uint    `json:"maxResults"`
	Total      uint    `json:"total"`
	Issues     []Issue `json:"issues"`
}

type Issue struct {
	ID     string                 `json:"id"`
	Expand string                 `json:"expand"`
	Self   string                 `json:"self"`
	Key    string                 `json:"key"`
	Fields map[string]interface{} `json:"fields"`
}

type IssueType struct {
	Self           string `json:"self"`
	ID             string `json:"id"`
	Description    string `json:"description"`
	IconURL        string `json:"iconUrl"`
	Name           string `json:"name"`
	Subtask        bool   `json:"subtask"`
	AvatarID       uint   `json:"avatarId"`
	HierarchyLevel uint   `json:"hierarchyLevel"`
}

type IssueStatus struct {
	Self           string
	Description    *string
	IconURL        string
	Name           string
	ID             string
	StatusCategory StatusCategory
}
