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

func TestFormValueTrim(t *testing.T) {
	t.Parallel()
	var (
		// err     error
		subject string
		expect  string
	)

	subject = "test "
	expect = "test"

	app := fiber.New()

	app.Post("/test", func(c *fiber.Ctx) error {
		cc := Ctx{c}
		utils.AssertEqual(t, expect, cc.FormValueTrim("name"))
		return nil
	})
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	utils.AssertEqual(t, nil, writer.WriteField("name", subject))

	writer.Close()
	req := httptest.NewRequest(fiber.MethodPost, "/test", body)
	req.Header.Set("Content-Type", fmt.Sprintf("multipart/form-data; boundary=%s", writer.Boundary()))
	req.Header.Set("Content-Length", strconv.Itoa(len(body.Bytes())))

	resp, err := app.Test(req)
	utils.AssertEqual(t, nil, err, "app.Test(req)")
	utils.AssertEqual(t, fiber.StatusOK, resp.StatusCode, "Status code")
}

func TestFormValueDate(t *testing.T) {
	t.Parallel()
	var (
		// err     error
		subject string
		expect  time.Time
	)

	subject = "1990-12-09"
	expect = time.Date(1990, 12, 9, 0, 0, 0, 0, time.Local)

	app := fiber.New()

	app.Post("/test", func(c *fiber.Ctx) error {
		cc := Ctx{c}
		utils.AssertEqual(t, expect, cc.FormValueDate("name"))
		return nil
	})
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	utils.AssertEqual(t, nil, writer.WriteField("name", subject))

	writer.Close()
	req := httptest.NewRequest(fiber.MethodPost, "/test", body)
	req.Header.Set("Content-Type", fmt.Sprintf("multipart/form-data; boundary=%s", writer.Boundary()))
	req.Header.Set("Content-Length", strconv.Itoa(len(body.Bytes())))

	resp, err := app.Test(req)
	utils.AssertEqual(t, nil, err, "app.Test(req)")
	utils.AssertEqual(t, fiber.StatusOK, resp.StatusCode, "Status code")
}

func TestFormValueTime(t *testing.T) {
	t.Parallel()
	var (
		// err     error
		subject string
		expect  time.Time
	)

	subject = "15:04"
	expect = time.Date(0, 1, 1, 15, 4, 0, 0, time.Local)

	app := fiber.New()

	app.Post("/test", func(c *fiber.Ctx) error {
		cc := Ctx{c}
		utils.AssertEqual(t, expect, cc.FormValueTime("name"))
		return nil
	})
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	utils.AssertEqual(t, nil, writer.WriteField("name", subject))

	writer.Close()
	req := httptest.NewRequest(fiber.MethodPost, "/test", body)
	req.Header.Set("Content-Type", fmt.Sprintf("multipart/form-data; boundary=%s", writer.Boundary()))
	req.Header.Set("Content-Length", strconv.Itoa(len(body.Bytes())))

	resp, err := app.Test(req)
	utils.AssertEqual(t, nil, err, "app.Test(req)")
	utils.AssertEqual(t, fiber.StatusOK, resp.StatusCode, "Status code")
}

func TestFormValueDateTime(t *testing.T) {
	t.Parallel()
	var (
		// err     error
		subject string
		expect  time.Time
	)

	subject = "1990-12-09T15:04"
	expect = time.Date(1990, 12, 9, 15, 4, 0, 0, time.Local)

	app := fiber.New()

	app.Post("/test", func(c *fiber.Ctx) error {
		cc := Ctx{c}
		utils.AssertEqual(t, expect, cc.FormValueDateTime("name"))
		return nil
	})
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	utils.AssertEqual(t, nil, writer.WriteField("name", subject))

	writer.Close()
	req := httptest.NewRequest(fiber.MethodPost, "/test", body)
	req.Header.Set("Content-Type", fmt.Sprintf("multipart/form-data; boundary=%s", writer.Boundary()))
	req.Header.Set("Content-Length", strconv.Itoa(len(body.Bytes())))

	resp, err := app.Test(req)
	utils.AssertEqual(t, nil, err, "app.Test(req)")
	utils.AssertEqual(t, fiber.StatusOK, resp.StatusCode, "Status code")
}
