package helpers

type ResponseForm struct {
	Success    bool            `json:"success"`
	Result     interface{}     `json:"result"`
	Messages   []string        `json:"messages"`
	Errors     []*ResposeError `json:"errors"`
	ResultInfo *ResultInfo     `json:"result_info,omitempty"`
}

type ResposeError struct {
	Code    int         `json:"code"`
	Source  interface{} `json:"source,omitempty"`
	Title   string      `json:"title,omitempty"`
	Message string      `json:"message"`
}

type ResultInfo struct {
	Page      int `json:"page"`
	PerPage   int `json:"per_page"`
	Count     int `json:"count"`
	TotalCont int `json:"total_count"`
}
