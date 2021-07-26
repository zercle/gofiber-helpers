package helpers

import (
	"bytes"
	"fmt"
	"mime/multipart"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

type TestFormData struct {
	Subject  string
	Expect   interface{}
	MustFail bool
	Err      error
	T        *testing.T
}

func (d *TestFormData) TestFormCall(testFn string) {
	app := fiber.New()

	switch testFn {
	case "FormValueTrim":
		app.Post("/test", func(c *fiber.Ctx) (err error) {
			cc := Ctx{c}
			utils.AssertEqual(d.T, d.Expect, cc.FormValueTrim("subject"))
			return
		})
	case "FormValueDate":
		app.Post("/test", func(c *fiber.Ctx) (err error) {
			cc := Ctx{c}
			utils.AssertEqual(d.T, d.Expect, cc.FormValueDate("subject"))
			return
		})
	case "FormValueTime":
		app.Post("/test", func(c *fiber.Ctx) (err error) {
			cc := Ctx{c}
			utils.AssertEqual(d.T, d.Expect, cc.FormValueTime("subject"))
			return
		})
	case "FormValueDateTime":
		app.Post("/test", func(c *fiber.Ctx) (err error) {
			cc := Ctx{c}
			utils.AssertEqual(d.T, d.Expect, cc.FormValueDateTime("subject"))
			return
		})
	case "FormValueBase64":
		app.Post("/test", func(c *fiber.Ctx) (err error) {
			cc := Ctx{c}
			utils.AssertEqual(d.T, d.Expect, cc.FormValueBase64("subject"))
			return
		})
	case "FormValueInt":
		app.Post("/test", func(c *fiber.Ctx) (err error) {
			cc := Ctx{c}
			result, err := cc.FormValueInt("subject")
			utils.AssertEqual(d.T, d.Expect, result)
			return
		})
	case "FormValueIntDefault":
		app.Post("/test", func(c *fiber.Ctx) (err error) {
			cc := Ctx{c}
			result := cc.FormValueIntDefault("subject", d.Expect.(int))
			utils.AssertEqual(d.T, d.Expect, result)
			return
		})
	case "FormValueInt64":
		app.Post("/test", func(c *fiber.Ctx) (err error) {
			cc := Ctx{c}
			result, err := cc.FormValueInt64("subject")
			utils.AssertEqual(d.T, d.Expect, result)
			return
		})
	case "FormValueInt64Default":
		app.Post("/test", func(c *fiber.Ctx) (err error) {
			cc := Ctx{c}
			result := cc.FormValueInt64Default("subject", d.Expect.(int64))
			utils.AssertEqual(d.T, d.Expect, result)
			return
		})
	case "FormValueFloat64":
		app.Post("/test", func(c *fiber.Ctx) (err error) {
			cc := Ctx{c}
			result, err := cc.FormValueFloat64("subject")
			utils.AssertEqual(d.T, d.Expect, result)
			return
		})
	case "FormValueFloat64Default":
		app.Post("/test", func(c *fiber.Ctx) (err error) {
			cc := Ctx{c}
			result := cc.FormValueFloat64Default("subject", d.Expect.(float64))
			utils.AssertEqual(d.T, d.Expect, result)
			return
		})
	case "FormValueBool":
		app.Post("/test", func(c *fiber.Ctx) (err error) {
			cc := Ctx{c}
			utils.AssertEqual(d.T, d.Expect, cc.FormValueBool("subject"))
			return
		})
	}

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	utils.AssertEqual(d.T, nil, writer.WriteField("subject", d.Subject))

	writer.Close()
	req := httptest.NewRequest(fiber.MethodPost, "/test", body)
	req.Header.Set("Content-Type", fmt.Sprintf("multipart/form-data; boundary=%s", writer.Boundary()))
	req.Header.Set("Content-Length", strconv.Itoa(len(body.Bytes())))

	resp, err := app.Test(req)
	if d.MustFail {
		utils.AssertEqual(d.T, d.Err, err, "app.Test(req)")
	} else {
		utils.AssertEqual(d.T, nil, err, "app.Test(req)")
		utils.AssertEqual(d.T, fiber.StatusOK, resp.StatusCode, "Status code")
	}
}

func TestFormValueTrim(t *testing.T) {
	t.Parallel()

	testData := TestFormData{
		T:       t,
		Subject: "test ",
		Expect:  "test",
	}
	testData.TestFormCall("FormValueTrim")
}

func TestFormValueDate(t *testing.T) {
	t.Parallel()

	testData := TestFormData{
		T:       t,
		Subject: "1990-12-09",
		Expect:  time.Date(1990, 12, 9, 0, 0, 0, 0, time.Local),
	}
	testData.TestFormCall("FormValueDate")

	testData = TestFormData{
		T:        t,
		Subject:  "test",
		Expect:   time.Time{},
		MustFail: true,
	}
	testData.TestFormCall("FormValueDate")
}

func TestFormValueTime(t *testing.T) {
	t.Parallel()

	testData := TestFormData{
		T:       t,
		Subject: "15:04",
		Expect:  time.Date(0, 1, 1, 15, 4, 0, 0, time.Local),
	}
	testData.TestFormCall("FormValueTime")

	testData = TestFormData{
		T:        t,
		Subject:  "test",
		Expect:   time.Time{},
		MustFail: true,
	}
	testData.TestFormCall("FormValueTime")
}

func TestFormValueDateTime(t *testing.T) {
	t.Parallel()

	testData := TestFormData{
		T:       t,
		Subject: "1990-12-09T15:04",
		Expect:  time.Date(1990, 12, 9, 15, 4, 0, 0, time.Local),
	}
	testData.TestFormCall("FormValueDateTime")

	testData = TestFormData{
		T:        t,
		Subject:  "test",
		Expect:   time.Time{},
		MustFail: true,
	}
	testData.TestFormCall("FormValueDateTime")
}

func TestFormValueBase64(t *testing.T) {
	t.Parallel()

	testData := TestFormData{
		T:       t,
		Subject: "dGVzdCB1cmwgZW5jb2Rl",
		Expect:  "test url encode",
	}
	testData.TestFormCall("FormValueBase64")

	testData = TestFormData{
		T:       t,
		Subject: "dGVzdCBzdGQgZW5jb2Rl",
		Expect:  "test std encode",
	}
	testData.TestFormCall("FormValueBase64")

	testData = TestFormData{
		T:       t,
		Subject: "test pain text",
		Expect:  "test pain text",
	}
	testData.TestFormCall("FormValueBase64")
}

func TestFormValueInt(t *testing.T) {
	t.Parallel()

	testData := TestFormData{
		T:       t,
		Subject: "123",
		Expect:  123,
	}
	testData.TestFormCall("FormValueInt")

	testData = TestFormData{
		T:        t,
		Subject:  "",
		Expect:   0,
		MustFail: true,
	}
	testData.TestFormCall("FormValueInt")

	testData = TestFormData{
		T:        t,
		Subject:  "abc",
		Expect:   0,
		MustFail: true,
	}
	testData.TestFormCall("FormValueInt")
}

func TestFormValueIntDefault(t *testing.T) {
	t.Parallel()

	testData := TestFormData{
		T:       t,
		Subject: "123",
		Expect:  123,
	}
	testData.TestFormCall("FormValueIntDefault")

	testData = TestFormData{
		T:       t,
		Subject: "abc",
		Expect:  0,
	}
	testData.TestFormCall("FormValueIntDefault")
}

func TestFormValueInt64(t *testing.T) {
	t.Parallel()

	testData := TestFormData{
		T:       t,
		Subject: "123",
		Expect:  int64(123),
	}
	testData.TestFormCall("FormValueInt64")

	testData = TestFormData{
		T:        t,
		Subject:  "",
		Expect:   int64(0),
		MustFail: true,
	}
	testData.TestFormCall("FormValueInt64")

	testData = TestFormData{
		T:        t,
		Subject:  "abc",
		Expect:   int64(0),
		MustFail: true,
	}
	testData.TestFormCall("FormValueInt64")
}

func TestFormValueInt64Default(t *testing.T) {
	t.Parallel()

	testData := TestFormData{
		T:       t,
		Subject: "123",
		Expect:  int64(123),
	}
	testData.TestFormCall("FormValueInt64Default")

	testData = TestFormData{
		T:        t,
		Subject:  "",
		Expect:   int64(0),
		MustFail: true,
	}
	testData.TestFormCall("FormValueInt64Default")

	testData = TestFormData{
		T:       t,
		Subject: "abc",
		Expect:  int64(0),
	}
	testData.TestFormCall("FormValueInt64Default")
}

func TestFormValueFloat64(t *testing.T) {
	t.Parallel()

	testData := TestFormData{
		T:       t,
		Subject: "1.23",
		Expect:  1.23,
	}
	testData.TestFormCall("FormValueFloat64")

	testData = TestFormData{
		T:        t,
		Subject:  "",
		Expect:   0.0,
		MustFail: true,
	}
	testData.TestFormCall("FormValueFloat64")

	testData = TestFormData{
		T:        t,
		Subject:  "abc",
		Expect:   0.0,
		MustFail: true,
	}
	testData.TestFormCall("FormValueFloat64")
}

func TestFormValueFloat64Default(t *testing.T) {
	t.Parallel()

	testData := TestFormData{
		T:       t,
		Subject: "1.23",
		Expect:  1.23,
	}
	testData.TestFormCall("FormValueFloat64Default")

	testData = TestFormData{
		T:        t,
		Subject:  "",
		Expect:   0.0,
		MustFail: true,
	}
	testData.TestFormCall("FormValueFloat64Default")

	testData = TestFormData{
		T:        t,
		Subject:  "abc",
		Expect:   0.0,
		MustFail: true,
	}
	testData.TestFormCall("FormValueFloat64Default")
}

func TestFormValueBool(t *testing.T) {
	t.Parallel()

	testData := TestFormData{
		T:       t,
		Subject: "1",
		Expect:  true,
	}
	testData.TestFormCall("FormValueBool")

	testData = TestFormData{
		T:       t,
		Subject: "true",
		Expect:  true,
	}
	testData.TestFormCall("FormValueBool")

	testData = TestFormData{
		T:       t,
		Subject: "0",
		Expect:  false,
	}
	testData.TestFormCall("FormValueBool")

	testData = TestFormData{
		T:       t,
		Subject: "false",
		Expect:  false,
	}
	testData.TestFormCall("FormValueBool")

	testData = TestFormData{
		T:       t,
		Subject: "wth",
		Expect:  false,
	}
	testData.TestFormCall("FormValueBool")
}
