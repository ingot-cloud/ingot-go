package code

// SUCCESS code
const SUCCESS = 99990200

// Common code
const (
	Unknown             = 99990001
	Unauthorized        = 99990401
	Forbidden           = 99990403
	NoRoute             = 99990404
	NoMethod            = 99990405
	InternalServerError = 99990500
)

// Token Code
const (
	TokenInvalid = 99991000 + iota
	TokenExpired
)
