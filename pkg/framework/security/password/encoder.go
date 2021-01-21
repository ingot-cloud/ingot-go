package password

// Encoder Service interface for encoding passwords.
type Encoder interface {
	// Encode the raw password
	Encode(raw string) (string, error)
	// Verify the encoded password obtained from storage matches the submitted raw
	// password after it too is encoded. Returns true if the passwords match, false if
	// they do not
	Matches(raw string, encodedPassword string) bool
}
