package helpers

import (
	"log"
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

func (e *Error) Log() {
	log.Printf("source: %+s \nerr: %+s \n", e.Source, e.Message)
}

func NewError(code int, message ...string) (err *Error) {
	if len(message) == 0 {
		message = append(message, http.StatusText(code))
	}
	err = &Error{
		Code:    code,
		Source:  WhereAmI(2),
		Title:   http.StatusText(code),
		Message: strings.Join(message, " \n"),
	}
	return
}

func NewErrorSource(code int, source string, message ...string) (err *Error) {
	if len(message) == 0 {
		message = append(message, http.StatusText(code))
	}
	err = &Error{
		Code:    code,
		Source:  source,
		Title:   http.StatusText(code),
		Message: strings.Join(message, " \n"),
	}
	return
}