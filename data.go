package helpers

import (
	"fmt"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/segmentio/encoding/json"
)

func FileNameWithoutExtension(fileName string) string {
	_, fileName = path.Split(fileName)
	fileExtension := path.Ext(fileName)
	if pos := strings.LastIndex(fileName, fileExtension); pos != -1 {
		return fileName[:pos]
	}
	return fileName
}

func DateStrTotime(dateStr string) (result time.Time, err error) {
	time.Local, err = time.LoadLocation("Asia/Bangkok")
	if err != nil {
		err = fiber.NewError(http.StatusInternalServerError, err.Error())
		return
	}
	dateStr = strings.TrimSpace(dateStr)
	var (
		yearStr  string
		monthStr string
		dayStr   string
	)
	// convert / to -
	dateStr = strings.ReplaceAll(dateStr, "/", "-")
	currentTime := time.Now()

	if strings.Contains(dateStr, "-") {
		dateArr := strings.Split(dateStr, "-")
		if len(dateArr) != 3 {
			err = fiber.NewError(http.StatusBadRequest, "dateStr with /,- must in ISO date format")
			return
		}
		if len(dateArr[0]) == 4 {
			yearStr = dateArr[0]
			monthStr = dateArr[1]
			dayStr = dateArr[2]
		} else {
			yearStr = dateArr[2]
			monthStr = dateArr[1]
			dayStr = dateArr[0]
		}
		// try yyyy-mm-dd
		if monthStr == "00" {
			monthStr = "01"
		}
		if dayStr == "00" {
			dayStr = "01"
		}
		result, err = time.ParseInLocation("2006-01-02", fmt.Sprintf("%04s-%02s-%02s", yearStr, monthStr, dayStr), time.Local)
		if err == nil && !result.IsZero() {
			// convert from BE
			if result.Year() > (currentTime.Year() + 272) {
				result = result.AddDate(-543, 0, 0)
			}
			return
		}
		err = fiber.NewError(http.StatusBadRequest, err.Error())
	}

	if len(dateStr) == 8 {
		// try yyyymmdd
		yearStr = dateStr[:4]
		monthStr = dateStr[4:6]
		dayStr = dateStr[6:8]
		if monthStr == "00" {
			monthStr = "01"
		}
		if dayStr == "00" {
			dayStr = "01"
		}
		result, err = time.ParseInLocation("2006-01-02", fmt.Sprintf("%04s-%02s-%02s", yearStr, monthStr, dayStr), time.Local)
		if err == nil && !result.IsZero() {
			// convert from BE
			if result.Year() > (currentTime.Year() + 272) {
				result = result.AddDate(-543, 0, 0)
			}
			return
		}
		// try ddmmyyyy
		yearStr = dateStr[4:8]
		monthStr = dateStr[2:4]
		dayStr = dateStr[:2]
		result, err = time.ParseInLocation("2006-01-02", fmt.Sprintf("%04s-%02s-%02s", yearStr, monthStr, dayStr), time.Local)
		if err == nil && !result.IsZero() {
			// convert from BE
			if result.Year() > (currentTime.Year() + 272) {
				result = result.AddDate(-543, 0, 0)
			}
			return
		}
		err = fiber.NewError(http.StatusBadRequest, err.Error())
	}
	err = fiber.NewError(http.StatusBadRequest, "dateStr must be 8,10 chars")
	return
}

func TimeToDGADate(dateTime time.Time) (dgaDate string) {
	dgaDate = fmt.Sprintf("%04d%02d%02d", dateTime.Year()+543, dateTime.Month(), dateTime.Day())
	return
}

func ValidCID(cid string) (result bool, err error) {
	sum := 0
	if len(cid) == 13 {
		for i := 0; i < 12; i++ {
			n, _ := strconv.Atoi(string(cid[i]))
			sum = sum + ((13 - i) * n)
		}

		checkDigit := (11 - (sum % 11)) % 10

		if strconv.Itoa(checkDigit) != string(cid[12]) {
			err = fiber.NewError(http.StatusBadRequest, fmt.Sprintf("invalid cid: %+v", cid))
		} else {
			result = true
		}
	} else {
		err = fiber.NewError(http.StatusBadRequest, fmt.Sprintf("cid must be 13 digits: %+v", cid))
	}
	return
}

func ENVJSONArray(name string) (value []interface{}, err error) {
	jsonString := os.Getenv(name)
	err = json.Unmarshal([]byte(jsonString), &value)
	return
}

func ENVJSONObj(name string) (value map[string]interface{}, err error) {
	jsonString := os.Getenv(name)
	err = json.Unmarshal([]byte(jsonString), &value)
	return
}
