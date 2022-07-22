package services

import (
	"time"

	"github.com/tobiwild/holidays"
)

func IsWeekend(t time.Time) bool {
	t = t.Local()

	switch t.Weekday() {
	case time.Saturday:
		return true
	case time.Sunday:
		return true
	}
	return false
}

func FormatTime(timestr string) (string, time.Time, error) {
	if timestr != "" {
		estLocation, err := time.LoadLocation("America/Sao_Paulo")
		if err != nil {
			return "", time.Time{}, err
		}

		layout := "2006-01-02"

		t, erro := time.ParseInLocation(layout, timestr, estLocation)
		if erro != nil {
			return "", time.Time{}, err
		}

		return t.Format(layout), t, nil
	}
	return "", time.Time{}, nil
}

func CheckHoliday(t time.Time) bool {
	holidays.SetHolidaysFunction(holidays.HolidaysBR)
	return holidays.IsHoliday(t)
}
