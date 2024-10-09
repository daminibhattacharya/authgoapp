package models


type Response struct {
    ID       string `json:"_id"`
    Email    string `json:"email"`
    Verified bool   `json:"email_verified"`
	Name        string      `json:"name"`
	Message     string      `json:"message"`
	Code        string      `json:"code"`
	Description Description `json:"description"`
	Policy      string      `json:"policy"`
	StatusCode  int         `json:"statusCode"`
}
type Description struct {
	Rules    []Rule `json:"rules"`
	Verified bool   `json:"verified"`
}

type Rule struct {
	Message  string     `json:"message"`
	Format   []int      `json:"format,omitempty"`
	Code     string     `json:"code"`
	Verified bool       `json:"verified"`
	Items    []RuleItem `json:"items,omitempty"`
}

type RuleItem struct {
	Message  string `json:"message"`
	Code     string `json:"code"`
	Verified bool   `json:"verified"`
}
