package utils

import (
	"fmt"
	"strconv"
	"time"
)

func ModifyBirthday(BirthDay int, BirthMonth int, BirthYear int) string {
	var birthday string
	var dayOfBirth string
	var monthOfBirth string

	if BirthDay < 10 {
		dayOfBirth = fmt.Sprintf("0%d", BirthDay)
	} else {
		dayOfBirth = strconv.Itoa(BirthDay)
	}

	if BirthMonth < 10 {
		monthOfBirth = fmt.Sprintf("0%d", BirthMonth)
	} else {
		monthOfBirth = strconv.Itoa(BirthMonth)
	}

	birthday = fmt.Sprintf("%s/%s/%d", dayOfBirth, monthOfBirth, BirthYear)
	return birthday
}

func IsValidDay(day int, month int, year int) bool {
	if day == 0 || month == 0 || month > 12 || year == 0 {
		return false
	}

	if month == 2 && day > 29 && IsLeapYear(year) {
		return false
	}

	if month == 2 && day > 28 && !IsLeapYear(year) {
		return false
	}

	if month != 2 && day > getLastDayOfMonth(int64(month)) {
		return false
	}

	return true
}

func IsLeapYear(year int) bool {
	return year%4 == 0
}

func getLastDayOfMonth(month int64) int {
	monthWith30days := [4]int64{4, 6, 9, 11}
	monthWith31days := [7]int64{1, 3, 5, 7, 8, 10, 12}

	var lastDayOfMonth int
	if IsInsideArray(month, monthWith30days) {
		return 30
	}

	if IsInsideArray(month, monthWith31days) {
		return 31
	}

	return lastDayOfMonth
}

func FormatDateTime(datetime time.Time) string {
	return datetime.Format("2006-01-02 15:04")
}