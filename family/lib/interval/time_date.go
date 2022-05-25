package interval

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type TimeDate struct {
	time.Time
}

func (t *TimeDate) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}
	// Fractional seconds are handled implicitly by Parse.
	var err error
	t.Time, err = time.ParseInLocation(`"2006-01-02"`, string(data), time.Local)

	return err
}

func (t TimeDate) MarshalJSON() ([]byte, error) {
	tune := t.Format(`"2006-01-02"`)
	return []byte(tune), nil
}

func (t TimeDate) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}

func (t *TimeDate) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = TimeDate{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}
