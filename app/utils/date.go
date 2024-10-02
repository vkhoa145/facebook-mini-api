package utils

import (
	"fmt"
	"strconv"
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
	if day == 0 || month == 0 || year == 0 {
		return false
	}

	if month == 2 && day > 29 && IsLeapYear(year) {
		return false
	}

	if month == 2 && day > 28 && !IsLeapYear(year) {
		return false
	}

	if month <= 7 && day > 30 {
		return false
	}

	if month > 7 && day > 31 {
		return false
	}

	return true
}

func IsLeapYear(year int) bool {
	return year%4 == 0
}
