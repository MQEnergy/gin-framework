package util

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type FormatTime time.Time

// Value implements the driver Valuer interface
func (t FormatTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	tlt := time.Time(t)
	//	判断给定的时间是否跟默认零时间的时间戳相同
	if tlt.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return tlt, nil
}

// Scan implements the Scanner interface
func (t *FormatTime) Scan(v interface{}) error {
	if value, ok := v.(time.Time); ok {
		*t = FormatTime(value)
		return nil
	}
	return fmt.Errorf("can't convert %v to timestamp", v)
}

// MarshalJSON 实现时间的json序列化
func (t FormatTime) MarshalJSON() ([]byte, error) {
	var stamp = fmt.Sprintf("\"%s\"", time.Time(t).Format("2006-01-02 15:04:05"))
	return []byte(stamp), nil
}

// UnmarshalJSON 反序列化
func (t *FormatTime) UnmarshalJSON(b []byte) error {
	now, err := time.ParseInLocation(`"2006-01-02 15:04:05"`, string(b), time.Local)
	*t = FormatTime(now)
	return err
}
