package helpers

type ResponseForm struct {
	Success    bool            `json:"success"`
	Result     interface{}     `json:"result,omitempty"`
	Data       interface{}     `json:"data,omitempty"`
	Messages   []string        `json:"messages,omitempty"`
	Errors     []ResponseError `json:"errors,omitempty"`
	ResultInfo *ResultInfo     `json:"result_info,omitempty"`
}

// ResponseError
// backward complatible Error
type ResponseError Error

type ResultInfo struct {
	Page      int `json:"page"`
	PerPage   int `json:"per_page"`
	Count     int `json:"count"`
	TotalCont int `json:"total_count"`
}

type AuthType string

const (
	BasicAuth   AuthType = "basic"
	BearerToken AuthType = "bearer"
)

type HttpAuth struct {
	Type     AuthType `json:"auth_type"`
	Token    string   `json:"token,omitempty"`
	Username string   `json:"username,omitempty"`
	Password string   `json:"password,omitempty"`
}
