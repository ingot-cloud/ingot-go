package password

import (
	"strings"
)

// NoopEncoder impl
type NoopEncoder struct {
}

// Encode the raw password
func (e *NoopEncoder) Encode(raw string) (string, error) {
	return raw, nil
}

// Matches Verify the encoded password obtained from storage matches the submitted raw
// password after it too is encoded. Returns true if the passwords match, false if
// they do not
func (e *NoopEncoder) Matches(raw string, encodedPassword string) (bool, error) {
	return strings.Compare(raw, encodedPassword) == 0, nil
}
