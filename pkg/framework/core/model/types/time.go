package types

import (
	"database/sql/driver"
	"fmt"
	"strconv"
	"time"
)

// TimeStamp 时间戳
type TimeStamp time.Time

// UnmarshalJSON implements the json.Unmarshaler interface.
// The time is expected to be a quoted string in RFC 3339 format.
func (ts *TimeStamp) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}
	millis, err := strconv.ParseInt(string(data), 10, 64)

	*ts = TimeStamp(time.Unix(0, millis*int64(time.Millisecond)))
	return err
}

// MarshalJSON implements the json.MarshalJSON interface.
func (ts TimeStamp) MarshalJSON() ([]byte, error) {
	origin := time.Time(ts)
	return []byte(strconv.FormatInt(origin.UnixNano()/1000000, 10)), nil
}

// Value 实现 driver.Valuer 接口
func (ts TimeStamp) Value() (driver.Value, error) {
	var zeroTime time.Time
	var ti = time.Time(ts)
	if ti.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return ti, nil
}

// Scan 实现 sql.Scanner 接口
func (ts *TimeStamp) Scan(v any) error {
	value, ok := v.(time.Time)
	if ok {
		*ts = TimeStamp(value)
		return nil
	}

	return fmt.Errorf("Can not convert %v to timestamp", v)
}

func (ts TimeStamp) String() string {
	return ts.ToTime().Format("2006-01-02 15:04:05")
}

// ToTime to time.Time
func (ts *TimeStamp) ToTime() time.Time {
	return time.Time(*ts)
}
