package helpers

import (
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

type TestParamData struct {
	Subject  string
	Expect   interface{}
	MustFail bool
	Err      error
	T        *testing.T
}

func (d *TestParamData) TestFormCall(testFn string) {
	app := fiber.New()

	switch testFn {
	case "ParamTrim":
		app.Get("/test/:subject", func(c *fiber.Ctx) (err error) {
			cc := Ctx{c}
			utils.AssertEqual(d.T, d.Expect, cc.ParamTrim("subject"))
			return
		})
	case "ParamDate":
		app.Get("/test/:subject", func(c *fiber.Ctx) (err error) {
			cc := Ctx{c}
			utils.AssertEqual(d.T, d.Expect, cc.ParamDate("subject"))
			return
		})
	case "ParamTime":
		app.Get("/test/:subject", func(c *fiber.Ctx) (err error) {
			cc := Ctx{c}
			utils.AssertEqual(d.T, d.Expect, cc.ParamTime("subject"))
			return
		})
	case "ParamDateTime":
		app.Get("/test/:subject", func(c *fiber.Ctx) (err error) {
			cc := Ctx{c}
			utils.AssertEqual(d.T, d.Expect, cc.ParamDateTime("subject"))
			return
		})
	case "ParamBase64":
		app.Get("/test/:subject", func(c *fiber.Ctx) (err error) {
			cc := Ctx{c}
			utils.AssertEqual(d.T, d.Expect, cc.ParamBase64("subject"))
			return
		})
	case "ParamInt":
		app.Get("/test/:subject", func(c *fiber.Ctx) (err error) {
			cc := Ctx{c}
			result, err := cc.ParamInt("subject")
			utils.AssertEqual(d.T, d.Expect, result)
			return
		})
	case "ParamIntDefault":
		app.Get("/test/:subject", func(c *fiber.Ctx) (err error) {
			cc := Ctx{c}
			result := cc.ParamIntDefault("subject", d.Expect.(int))
			utils.AssertEqual(d.T, d.Expect, result)
			return
		})
	case "ParamInt64":
		app.Get("/test/:subject", func(c *fiber.Ctx) (err error) {
			cc := Ctx{c}
			result, err := cc.ParamInt64("subject")
			utils.AssertEqual(d.T, d.Expect, result)
			return
		})
	case "ParamInt64Default":
		app.Get("/test/:subject", func(c *fiber.Ctx) (err error) {
			cc := Ctx{c}
			result := cc.ParamInt64Default("subject", d.Expect.(int64))
			utils.AssertEqual(d.T, d.Expect, result)
			return
		})
	case "ParamFloat64":
		app.Get("/test/:subject", func(c *fiber.Ctx) (err error) {
			cc := Ctx{c}
			result, err := cc.ParamFloat64("subject")
			utils.AssertEqual(d.T, d.Expect, result)
			return
		})
	case "ParamFloat64Default":
		app.Get("/test/:subject", func(c *fiber.Ctx) (err error) {
			cc := Ctx{c}
			result := cc.ParamFloat64Default("subject", d.Expect.(float64))
			utils.AssertEqual(d.T, d.Expect, result)
			return
		})
	case "ParamBool":
		app.Get("/test/:subject", func(c *fiber.Ctx) (err error) {
			cc := Ctx{c}
			utils.AssertEqual(d.T, d.Expect, cc.ParamBool("subject"))
			return
		})
	}

	req := httptest.NewRequest(fiber.MethodGet, "/test/"+d.Subject, nil)
	req.Header.Set(fiber.HeaderContentType, fiber.MIMEApplicationForm)

	resp, err := app.Test(req)
	if d.MustFail {
		utils.AssertEqual(d.T, d.Err, err, "app.Test(req)")
	} else {
		utils.AssertEqual(d.T, nil, err, "app.Test(req)")
		utils.AssertEqual(d.T, fiber.StatusOK, resp.StatusCode, "Status code")
	}
}

func TestParamTrim(t *testing.T) {
	t.Parallel()

	testData := TestParamData{
		T:       t,
		Subject: "test%20",
		Expect:  "test",
	}
	testData.TestFormCall("ParamTrim")
}

func TestParamDate(t *testing.T) {
	t.Parallel()

	testData := TestParamData{
		T:       t,
		Subject: "1990-12-09",
		Expect:  time.Date(1990, 12, 9, 0, 0, 0, 0, time.Local),
	}
	testData.TestFormCall("ParamDate")

	testData = TestParamData{
		T:        t,
		Subject:  "test",
		Expect:   time.Time{},
		MustFail: true,
	}
	testData.TestFormCall("ParamDate")
}

func TestParamTime(t *testing.T) {
	t.Parallel()

	testData := TestParamData{
		T:       t,
		Subject: "15:04",
		Expect:  time.Date(0, 1, 1, 15, 4, 0, 0, time.Local),
	}
	testData.TestFormCall("ParamTime")

	testData = TestParamData{
		T:        t,
		Subject:  "test",
		Expect:   time.Time{},
		MustFail: true,
	}
	testData.TestFormCall("ParamTime")
}

func TestParamDateTime(t *testing.T) {
	t.Parallel()

	testData := TestParamData{
		T:       t,
		Subject: "1990-12-09T15:04",
		Expect:  time.Date(1990, 12, 9, 15, 4, 0, 0, time.Local),
	}
	testData.TestFormCall("ParamDateTime")

	testData = TestParamData{
		T:        t,
		Subject:  "test",
		Expect:   time.Time{},
		MustFail: true,
	}
	testData.TestFormCall("ParamDateTime")
}

func TestParamBase64(t *testing.T) {
	t.Parallel()

	testData := TestParamData{
		T:       t,
		Subject: "dGVzdCB1cmwgZW5jb2Rl",
		Expect:  "test url encode",
	}
	testData.TestFormCall("ParamBase64")

	testData = TestParamData{
		T:       t,
		Subject: "dGVzdCBzdGQgZW5jb2Rl",
		Expect:  "test std encode",
	}
	testData.TestFormCall("ParamBase64")

	testData = TestParamData{
		T:       t,
		Subject: "test%20pain%20text",
		Expect:  "test pain text",
	}
	testData.TestFormCall("ParamBase64")
}

func TestParamInt(t *testing.T) {
	t.Parallel()

	testData := TestParamData{
		T:       t,
		Subject: "123",
		Expect:  123,
	}
	testData.TestFormCall("ParamInt")

	testData = TestParamData{
		T:        t,
		Subject:  "",
		Expect:   0,
		MustFail: true,
	}
	testData.TestFormCall("ParamInt")

	testData = TestParamData{
		T:        t,
		Subject:  "abc",
		Expect:   0,
		MustFail: true,
	}
	testData.TestFormCall("ParamInt")
}

func TestParamIntDefault(t *testing.T) {
	t.Parallel()

	testData := TestParamData{
		T:       t,
		Subject: "123",
		Expect:  123,
	}
	testData.TestFormCall("ParamIntDefault")

	testData = TestParamData{
		T:       t,
		Subject: "abc",
		Expect:  0,
	}
	testData.TestFormCall("ParamIntDefault")
}

func TestParamInt64(t *testing.T) {
	t.Parallel()

	testData := TestParamData{
		T:       t,
		Subject: "123",
		Expect:  int64(123),
	}
	testData.TestFormCall("ParamInt64")

	testData = TestParamData{
		T:        t,
		Subject:  "",
		Expect:   int64(0),
		MustFail: true,
	}
	testData.TestFormCall("ParamInt64")

	testData = TestParamData{
		T:        t,
		Subject:  "abc",
		Expect:   int64(0),
		MustFail: true,
	}
	testData.TestFormCall("ParamInt64")
}

func TestParamInt64Default(t *testing.T) {
	t.Parallel()

	testData := TestParamData{
		T:       t,
		Subject: "123",
		Expect:  int64(123),
	}
	testData.TestFormCall("ParamInt64Default")

	testData = TestParamData{
		T:        t,
		Subject:  "",
		Expect:   int64(0),
		MustFail: true,
	}
	testData.TestFormCall("ParamInt64Default")

	testData = TestParamData{
		T:       t,
		Subject: "abc",
		Expect:  int64(0),
	}
	testData.TestFormCall("ParamInt64Default")
}

func TestParamFloat64(t *testing.T) {
	t.Parallel()

	testData := TestParamData{
		T:       t,
		Subject: "1.23",
		Expect:  1.23,
	}
	testData.TestFormCall("ParamFloat64")

	testData = TestParamData{
		T:        t,
		Subject:  "",
		Expect:   0.0,
		MustFail: true,
	}
	testData.TestFormCall("ParamFloat64")

	testData = TestParamData{
		T:        t,
		Subject:  "abc",
		Expect:   0.0,
		MustFail: true,
	}
	testData.TestFormCall("ParamFloat64")
}

func TestParamFloat64Default(t *testing.T) {
	t.Parallel()

	testData := TestParamData{
		T:       t,
		Subject: "1.23",
		Expect:  1.23,
	}
	testData.TestFormCall("ParamFloat64Default")

	testData = TestParamData{
		T:        t,
		Subject:  "",
		Expect:   0.0,
		MustFail: true,
	}
	testData.TestFormCall("ParamFloat64Default")

	testData = TestParamData{
		T:        t,
		Subject:  "abc",
		Expect:   0.0,
		MustFail: true,
	}
	testData.TestFormCall("ParamFloat64Default")
}

func TestParamBool(t *testing.T) {
	t.Parallel()

	testData := TestParamData{
		T:       t,
		Subject: "1",
		Expect:  true,
	}
	testData.TestFormCall("ParamBool")

	testData = TestParamData{
		T:       t,
		Subject: "true",
		Expect:  true,
	}
	testData.TestFormCall("ParamBool")

	testData = TestParamData{
		T:       t,
		Subject: "0",
		Expect:  false,
	}
	testData.TestFormCall("ParamBool")

	testData = TestParamData{
		T:       t,
		Subject: "false",
		Expect:  false,
	}
	testData.TestFormCall("ParamBool")

	testData = TestParamData{
		T:       t,
		Subject: "wth",
		Expect:  false,
	}
	testData.TestFormCall("ParamBool")
}
