package helpers

import (
	"encoding/base64"
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
	out, err := time.ParseInLocation("2006-01-02", strings.TrimSpace(c.FormValue(name)), time.Local)
	if err != nil {
		out = time.Time{}
	}
	return out
}

// FormValueTime returns the form field time value for the provided name.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/input/time
func (c *Ctx) FormValueTime(name string) time.Time {
	out, err := time.ParseInLocation("15:04", strings.TrimSpace(c.FormValue(name)), time.Local)
	if err != nil {
		out = time.Time{}
	}
	return out
}

// FormValueDateTime returns the form field datetime-local value for the provided name.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/input/datetime-local
func (c *Ctx) FormValueDateTime(name string) time.Time {
	out, err := time.ParseInLocation("2006-01-02T15:04", strings.TrimSpace(c.FormValue(name)), time.Local)
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
		v = string(de)
	}
	return v
}

// FormValueInt returns the form field value for the provided name, as int.
//
// If not found returns -1 and a non-nil error.
func (c *Ctx) FormValueInt(name string) (int, error) {
	v := c.FormValueTrim(name)
	if v == "" {
		return -1, fiber.ErrNotFound
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
// If not found returns -1 and a no-nil error.
func (c *Ctx) FormValueInt64(name string) (int64, error) {
	v := c.FormValueTrim(name)
	if v == "" {
		return -1, fiber.ErrNotFound
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
// If not found returns -1 and a non-nil error.
func (c *Ctx) FormValueFloat64(name string) (float64, error) {
	v := c.FormValueTrim(name)
	if v == "" {
		return -1, fiber.ErrNotFound
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
// If not found or value is false, then it returns false, otherwise true.
func (c *Ctx) FormValueBool(name string) (bool, error) {
	v := c.FormValueTrim(name)
	if v == "" {
		return false, fiber.ErrNotFound
	}

	return strconv.ParseBool(v)
}

// ParamDefault returns path parameter by name.
//
// Returns the "def" if not found.
func (c *Ctx) ParamDefault(name string, def string) string {
	if v := c.Params(name); len(v) > 0 {
		return v
	}
	return def
}

// ParamTrim returns path parameter by name, without trailing spaces.
func (c *Ctx) ParamTrim(name string) string {
	return strings.TrimSpace(c.Params(name))
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
// If not found returns -1 and a non-nil error.
func (c *Ctx) ParamInt(name string) (int, error) {
	v := c.ParamTrim(name)
	if v == "" {
		return -1, fiber.ErrNotFound
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
// If not found returns -1 and a no-nil error.
func (c *Ctx) ParamInt64(name string) (int64, error) {
	v := c.ParamTrim(name)
	if v == "" {
		return -1, fiber.ErrNotFound
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
// If not found returns -1 and a non-nil error.
func (c *Ctx) ParamFloat64(name string) (float64, error) {
	v := c.ParamTrim(name)
	if v == "" {
		return -1, fiber.ErrNotFound
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
// If not found or value is false, then it returns false, otherwise true.
func (c *Ctx) ParamBool(name string) (bool, error) {
	v := c.ParamTrim(name)
	if v == "" {
		return false, fiber.ErrNotFound
	}

	return strconv.ParseBool(v)
}
