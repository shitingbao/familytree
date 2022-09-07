package interval

import (
	"database/sql/driver"
	"fmt"
	"regexp"
	"strings"
	"time"
)

type TimeNormal struct {
	time.Time
}

const (
	DateTimeFormat              = "2006-01-02 15:04:05"
	DateFormat                  = "2006-01-02"
	ShortDateTimeFormat         = "20060102150405"
	ShortDateFormat             = "20060102"
	DateTimeFormatWithOutSecond = "2006-01-02 15:04"
)

const (
	PatternDateTimeFormat              = `^\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}$`
	PatternDateFormat                  = `^\d{4}-\d{2}-\d{2}$`
	PatternShortDateTimeFormat         = `^\d{4}\d{2}\d{2}\d{2}\d{2}\d{2}$`
	PatternShortDateFormat             = `^\d{4}\d{2}\d{2}$`
	PatternDateTimeFormatWithOutSecond = `^\d{4}-\d{2}-\d{2} \d{2}:\d{2}$`
)

func patternMatch(timeStr string, pattern string) bool {
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(timeStr)
}

var (
	patternFormats = map[string]string{
		PatternDateTimeFormat:              DateTimeFormat,
		PatternDateFormat:                  DateFormat,
		PatternShortDateTimeFormat:         ShortDateTimeFormat,
		PatternShortDateFormat:             ShortDateFormat,
		PatternDateTimeFormatWithOutSecond: DateTimeFormatWithOutSecond,
	}
)

func (t *TimeNormal) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}
	// Fractional seconds are handled implicitly by Parse.
	var err error

	timeStr := string(data)
	timeStr = strings.Trim(timeStr, "\"")
	for pattern, format := range patternFormats {
		if patternMatch(timeStr, pattern) {
			t.Time, err = time.ParseInLocation(format, timeStr, time.Local)
		}
	}

	return err
}

func (t TimeNormal) MarshalJSON() ([]byte, error) {
	tune := t.Format(`"2006-01-02 15:04:05"`)
	return []byte(tune), nil
}

func (t TimeNormal) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}

func (t *TimeNormal) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = TimeNormal{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}
