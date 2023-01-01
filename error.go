package helpers

import (
	"net/http"
	"strings"
)

// Error represents an error that occurred while handling a request.
type Error struct {
	Code    int         `json:"code"`
	Source  interface{} `json:"source,omitempty"`
	Title   string      `json:"title,omitempty"`
	Message string      `json:"message,omitempty"`
}

func (e *Error) Error() (errStr string) {
	return e.Message
}

func NewError(code int, source string, message ...string) (err *Error) {
	if len(message) == 0 {
		message = append(message, http.StatusText(code))
	}
	err = &Error{
		Code:    code,
		Source:  source,
		Title:   http.StatusText(code),
		Message: strings.Join(message, "\n "),
	}
	return
}
