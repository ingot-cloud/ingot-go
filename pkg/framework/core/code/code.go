package code

// SUCCESS code
const SUCCESS = "0200"

// Common code
const (
	BadRequest          = "0400"
	Unauthorized        = "0401"
	Forbidden           = "0403"
	NoRoute             = "0404"
	NoMethod            = "0405"
	InternalServerError = "0500"

	RequestFallback  = "0001"
	IllegalOperation = "0002"
	IllegalArgument  = "0003"

	IDClockBack = "1000"
)

// User Code
const (
	UserInvalid      = "U0001"
	UserDisabled     = "U0002"
	UserAppForbidden = "U0003"
)
