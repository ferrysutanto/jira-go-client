package jira

type OrderDir int

const (
	OrderAsc OrderDir = iota
	OrderDesc
)

type User struct {
	Self         string            `json:"self"`
	AccountID    string            `json:"accountId"`
	EmailAddress string            `json:"emailAddress"`
	AvatarURLs   map[string]string `json:"avatarUrls"`
	DisplayName  string            `json:"displayName"`
	Active       bool              `json:"active"`
	TimeZone     string            `json:"timeZone"`
	AccountType  string            `json:"accountType"`
}

type StatusCategory struct {
	Self      string
	ID        uint
	Key       string
	ColorName string
	Name      string
}

type Project struct {
	Self           string
	ID             string
	Key            string
	Name           string
	ProjectTypeKey string
	Simplified     bool
	AvatarUrls     map[string]string
}
