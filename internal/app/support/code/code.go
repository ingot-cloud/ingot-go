package code

// SUCCESS code
const SUCCESS = 99990200

// Common code
const (
	Unknown             = 99990001
	IllegalArgument     = 99990002
	IllegalOperation    = 99990003
	Unauthorized        = 99990401
	Forbidden           = 99990403
	NoRoute             = 99990404
	NoMethod            = 99990405
	InternalServerError = 99990500
	IDClockBack         = 99991000
)

// Token Code
const (
	TokenInvalid = 99991000 + iota
	TokenExpired
)

// User Code
const (
	UserInvalid = 99992000 + iota
	UserDisabled
	UserAppForbidden
	UserEnterpriseDisabled
)
