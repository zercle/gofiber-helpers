package helpers

import "net/http"

// Error represents an error that occurred while handling a request.
type Error struct {
	Code     int         `json:"code"`
	Source   interface{} `json:"source,omitempty"`
	Title    string      `json:"title,omitempty"`
	Messages []string    `json:"messages"`
}

func (e *Error) Error() (errStr string) {
	if len(e.Messages) != 0 {
		errStr = e.Messages[0]
	}
	return
}

func NewError(code int, source string, message ...string) (err *Error) {
	if len(message) == 0 {
		message = append(message, http.StatusText(code))
	}
	err = &Error{
		Code:     code,
		Source:   source,
		Title:    http.StatusText(code),
		Messages: message,
	}
	return
}
