package types

import (
	"database/sql/driver"
	"fmt"
	"strconv"
)

// ID ID类型
type ID int64

// UnmarshalJSON 字符串转为int64  json.Unmarshaler interface
func (id *ID) UnmarshalJSON(data []byte) error {
	if string(data) == "" {
		return nil
	}

	val, err := strconv.ParseInt(string(data), 10, 64)
	if err != nil {
		return err
	}

	*id = ID(val)
	return nil
}

// MarshalJSON int64转字符串  json.MarshalJSON interface
func (id ID) MarshalJSON() ([]byte, error) {
	return []byte(id.String()), nil
}

func (id ID) String() string {
	return strconv.FormatInt(int64(id), 10)
}

// Scan sql.Scanner interface
func (id *ID) Scan(value interface{}) error {
	result, ok := value.(int64)
	if ok {
		*id = ID(result)
		return nil
	}

	return fmt.Errorf("Can not convert %v to ID", value)
}

// Value driver.Valuer interface
func (id ID) Value() (driver.Value, error) {
	return int64(id), nil
}
