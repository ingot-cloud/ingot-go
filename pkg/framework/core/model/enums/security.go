package enums

const (
	// AuthTypeUnique 唯一授权类型
	AuthTypeUnique = "unique"
	// AuthTypeStandard 标准授权类型
	AuthTypeStandard = "standard"
)

// TokenType token类型
type TokenType string

const (
	// BearerToken bearer
	BearerToken TokenType = "Bearer"
	// BasicToken basic
	BasicToken TokenType = "Basic"
)

const (
	// BearerWithSpace bearer带空格
	BearerWithSpace = BearerToken + " "
	// BasicWithSpace basic带空格
	BasicWithSpace = BasicToken + " "
)
