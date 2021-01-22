package types

import (
	"strconv"
)

// ID ID类型
type ID int64

// UnmarshalJSON 字符串转为int64
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

// MarshalJSON int64转字符串
func (id ID) MarshalJSON() ([]byte, error) {
	val := strconv.FormatInt(int64(id), 10)
	return []byte(val), nil
}
