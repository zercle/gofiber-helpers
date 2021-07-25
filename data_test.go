package helpers

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

func TestFileNameWithoutExtension(t *testing.T) {
	t.Parallel()
	var (
		// err     error
		subject string
		expect  string
	)

	subject = "/dir0/dir1/test.ext"
	expect = "test"
	result := FileNameWithoutExtension(subject)

	utils.AssertEqual(t, expect, result)
}

func TestDateStrTotime(t *testing.T) {
	t.Parallel()
	var (
		err     error
		subject string
		expect  time.Time
	)

	subject = "1990-12-09"
	expect = time.Date(1990, 12, 9, 0, 0, 0, 0, time.Local)
	result, err := DateStrTotime(subject)
	if err != nil {
		utils.AssertEqual(t, nil, err)
	}
	utils.AssertEqual(t, expect.Unix(), result.Unix())

	subject = "09-12-1990"
	result, err = DateStrTotime(subject)
	if err != nil {
		utils.AssertEqual(t, nil, err)
	}
	utils.AssertEqual(t, expect.Unix(), result.Unix())

	subject = "19901209"
	result, err = DateStrTotime(subject)
	if err != nil {
		utils.AssertEqual(t, nil, err)
	}
	utils.AssertEqual(t, expect.Unix(), result.Unix())

	subject = "09121990"
	result, err = DateStrTotime(subject)
	if err != nil {
		utils.AssertEqual(t, nil, err)
	}
	utils.AssertEqual(t, expect.Unix(), result.Unix())
}

func TestTimeToDGADate(t *testing.T) {
	t.Parallel()
	var (
		// err     error
		subject time.Time
		expect  string
	)

	subject = time.Date(1990, 12, 9, 0, 0, 0, 0, time.Local)
	expect = "25331209"
	result := TimeToDGADate(subject)

	utils.AssertEqual(t, expect, result)
}

func TestValidCID(t *testing.T) {
	t.Parallel()
	var (
		err     error
		subject string
		expect  bool
	)

	subject = "1111111111119"
	expect = true
	result, err := ValidCID(subject)
	if err != nil {
		utils.AssertEqual(t, nil, err)
	}
	utils.AssertEqual(t, expect, result)

	subject = "1111111111110"
	expect = true
	result, err = ValidCID(subject)
	if err != nil {
		utils.AssertEqual(t, fiber.NewError(http.StatusBadRequest, fmt.Sprintf("invalid cid: %+v", subject)), err)
	}
	utils.AssertEqual(t, false, result)

	subject = "111111111111"
	expect = true
	result, err = ValidCID(subject)
	if err != nil {
		utils.AssertEqual(t, fiber.NewError(http.StatusBadRequest, fmt.Sprintf("cid must be 13 digits: %+v", subject)), err)
	}
	utils.AssertEqual(t, false, result)
}
