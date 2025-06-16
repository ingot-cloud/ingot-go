package types

import (
	"strconv"
)

// ID ID类型
type ID int64

// Zero 0
const Zero ID = 0

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

// NewIDFrom 实例化
func NewIDFrom(id any) ID {
	switch value := id.(type) {
	case ID:
		return value
	case string:
		return NewIDFromString(value)
	case int:
		return ID(value)
	case int64:
		return ID(value)
	}

	return Zero
}

// NewIDFromString 实例ID
func NewIDFromString(id string) ID {
	i, err := strconv.Atoi(id)
	if err != nil {
		return Zero
	}
	return ID(i)
}
