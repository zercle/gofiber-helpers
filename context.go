package helpers

import (
	"bytes"
	"encoding/base64"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
)

// Ctx represents the context of the current HTTP request. It holds request and
// response objects, path, path parameters, data and registered handler.
type Ctx struct {
	*fiber.Ctx
}

// FormValueTrim returns the form field value for the provided name, without trailing spaces.
func (c *Ctx) FormValueTrim(name string) string {
	return strings.TrimSpace(c.FormValue(name))
}

// FormValueDate returns the form field date value for the provided name.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/input/date
func (c *Ctx) FormValueDate(name string) time.Time {
	out, err := time.ParseInLocation("2006-01-02", c.FormValueTrim(name), time.Local)
	if err != nil {
		out = time.Time{}
	}
	return out
}

// FormValueTime returns the form field time value for the provided name.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/input/time
func (c *Ctx) FormValueTime(name string) time.Time {
	out, err := time.ParseInLocation("15:04", c.FormValueTrim(name), time.Local)
	if err != nil {
		out = time.Time{}
	}
	return out
}

// FormValueDateTime returns the form field datetime-local value for the provided name.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/input/datetime-local
func (c *Ctx) FormValueDateTime(name string) time.Time {
	out, err := time.ParseInLocation("2006-01-02T15:04", c.FormValueTrim(name), time.Local)
	if err != nil {
		out = time.Time{}
	}
	return out
}

// FormValueBase64 returns the form field value for the provided name.
//
// If value encoded with base64 return will be decoded string.
func (c *Ctx) FormValueBase64(name string) string {
	v := c.FormValueTrim(name)
	if de, err := base64.URLEncoding.DecodeString(v); err == nil {
		return string(de)
	}
	if de, err := base64.StdEncoding.DecodeString(v); err == nil {
		return string(de)
	}
	return v
}

// FormValueInt returns the form field value for the provided name, as int.
//
// If not found returns 0 and a non-nil error.
func (c *Ctx) FormValueInt(name string) (int, error) {
	v := c.FormValueTrim(name)
	if v == "" {
		return 0, fiber.ErrNotFound
	}
	return strconv.Atoi(v)
}

// FormValueIntDefault returns the form field value for the provided name, as int.
//
// If not found returns or parse errors the "def".
func (c *Ctx) FormValueIntDefault(name string, def int) int {
	if v, err := c.FormValueInt(name); err == nil {
		return v
	}

	return def
}

// FormValueInt64 returns the form field value for the provided name, as float64.
//
// If not found returns 0 and a no-nil error.
func (c *Ctx) FormValueInt64(name string) (int64, error) {
	v := c.FormValueTrim(name)
	if v == "" {
		return 0, fiber.ErrNotFound
	}
	return strconv.ParseInt(v, 10, 64)
}

// FormValueInt64Default returns the form field value for the provided name, as int64.
//
// If not found or parse errors returns the "def".
func (c *Ctx) FormValueInt64Default(name string, def int64) int64 {
	if v, err := c.FormValueInt64(name); err == nil {
		return v
	}

	return def
}

// FormValueFloat64 returns the form field value for the provided name, as float64.
//
// If not found returns 0 and a non-nil error.
func (c *Ctx) FormValueFloat64(name string) (float64, error) {
	v := c.FormValueTrim(name)
	if v == "" {
		return 0, fiber.ErrNotFound
	}
	return strconv.ParseFloat(v, 64)
}

// FormValueFloat64Default returns the form field value for the provided name, as float64.
//
// If not found or parse errors returns the "def".
func (c *Ctx) FormValueFloat64Default(name string, def float64) float64 {
	if v, err := c.FormValueFloat64(name); err == nil {
		return v
	}

	return def
}

// FormValueBool returns the form field value for the provided name, as bool.
//
// If not found or value is false, then it returns true, otherwise false.
func (c *Ctx) FormValueBool(name string) bool {
	v, err := strconv.ParseBool(c.FormValueTrim(name))
	if err != nil {
		v = false
	}
	return v
}

// FormValueInt returns the form field value for the provided name, as string array.
//
// If not found returns empty array.
func (c *Ctx) FormValueArray(name string, sep ...string) (result []string) {
	v := c.FormValueTrim(name)
	if len(v) == 0 {
		return
	}
	if len(sep) == 0 {
		sep = append(sep, ",")
	}
	return strings.Split(v, sep[0])
}

// ParamTrim returns path parameter by name, without trailing spaces.
func (c *Ctx) ParamTrim(name string) string {
	v, err := url.PathUnescape(c.Params(name))
	if err != nil {
		return c.Params(name)
	}
	return strings.TrimSpace(v)
}

// ParamDate returns the form field date value for the provided name.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/input/date
func (c *Ctx) ParamDate(name string) time.Time {
	out, err := time.ParseInLocation("2006-01-02", c.ParamTrim(name), time.Local)
	if err != nil {
		out = time.Time{}
	}
	return out
}

// ParamTime returns the form field time value for the provided name.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/input/time
func (c *Ctx) ParamTime(name string) time.Time {
	out, err := time.ParseInLocation("15:04", c.ParamTrim(name), time.Local)
	if err != nil {
		out = time.Time{}
	}
	return out
}

// ParamDateTime returns the form field datetime-local value for the provided name.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/input/datetime-local
func (c *Ctx) ParamDateTime(name string) time.Time {
	out, err := time.ParseInLocation("2006-01-02T15:04", c.ParamTrim(name), time.Local)
	if err != nil {
		out = time.Time{}
	}
	return out
}

// ParamBase64 returns path parameter by name.
//
// If value encoded with base64 return will be decoded string.
func (c *Ctx) ParamBase64(name string) string {
	v := c.ParamTrim(name)
	if de, err := base64.URLEncoding.DecodeString(v); err == nil {
		v = string(de)
	}
	return v
}

// ParamInt returns path parameter by name, as int.
//
// If not found returns 0 and a non-nil error.
func (c *Ctx) ParamInt(name string) (int, error) {
	v := c.ParamTrim(name)
	if v == "" {
		return 0, fiber.ErrNotFound
	}
	return strconv.Atoi(v)
}

// ParamIntDefault returns path parameter by name, as int.
//
// If not found returns or parse errors the "def".
func (c *Ctx) ParamIntDefault(name string, def int) int {
	if v, err := c.ParamInt(name); err == nil {
		return v
	}

	return def
}

// ParamInt64 returns path parameter by name, as float64.
//
// If not found returns 0 and a no-nil error.
func (c *Ctx) ParamInt64(name string) (int64, error) {
	v := c.ParamTrim(name)
	if v == "" {
		return 0, fiber.ErrNotFound
	}
	return strconv.ParseInt(v, 10, 64)
}

// ParamInt64Default returns path parameter by name, as int64.
//
// If not found or parse errors returns the "def".
func (c *Ctx) ParamInt64Default(name string, def int64) int64 {
	if v, err := c.ParamInt64(name); err == nil {
		return v
	}

	return def
}

// ParamFloat64 returns path parameter by name, as float64.
//
// If not found returns 0 and a non-nil error.
func (c *Ctx) ParamFloat64(name string) (float64, error) {
	v := c.ParamTrim(name)
	if v == "" {
		return 0, fiber.ErrNotFound
	}
	return strconv.ParseFloat(v, 64)
}

// ParamFloat64Default returns path parameter by name, as float64.
//
// If not found or parse errors returns the "def".
func (c *Ctx) ParamFloat64Default(name string, def float64) float64 {
	if v, err := c.ParamFloat64(name); err == nil {
		return v
	}

	return def
}

// ParamBool returns path parameter by name, as bool.
//
// If not found or value is false, then it returns true, otherwise false.
func (c *Ctx) ParamBool(name string) bool {
	v, err := strconv.ParseBool(c.ParamTrim(name))
	if err != nil {
		v = false
	}
	return v
}

func (c *Ctx) BasicAuth(user, passwd string) error {
	// Get authorization header
	authStr := c.Get(fiber.HeaderAuthorization)

	authSlices := strings.Split(authStr, " ")
	if len(authSlices) != 2 {
		return fiber.NewError(http.StatusUnauthorized, "invalid authorization format")
	}

	// Check if the header contains content besides "basic".
	if strings.ToLower(authSlices[0]) != "basic" {
		return fiber.NewError(http.StatusUnauthorized, "invalid basic auth format")
	}

	// Decode the header contents
	raw, err := base64.StdEncoding.DecodeString(authSlices[1])
	if err != nil {
		return fiber.NewError(http.StatusUnauthorized, err.Error())
	}

	// Check if the credentials are in the correct form
	// which is "username:password".
	credentials := bytes.Split(raw, []byte(":"))
	if len(credentials) != 2 {
		return fiber.NewError(http.StatusUnauthorized, "invalid basic auth format")
	}

	if bytes.Equal([]byte(user), credentials[0]) && bytes.Equal([]byte(passwd), credentials[1]) {
		return nil
	}

	// Authentication failed
	return fiber.NewError(http.StatusUnauthorized, "invalid user/passwd")
}
