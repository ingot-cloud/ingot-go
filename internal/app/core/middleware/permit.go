package middleware

import "github.com/gin-gonic/gin"

// PermitFunc permit function
type PermitFunc func(*gin.Context) bool

// Permit handler
func Permit(c *gin.Context, funcs ...PermitFunc) bool {
	for _, permit := range funcs {
		if permit(c) {
			return true
		}
	}
	return false
}

// NewPermitWithPrefix 创建 PermitFunc
func NewPermitWithPrefix(prefixes ...string) PermitFunc {
	return func(c *gin.Context) bool {
		path := c.Request.URL.Path
		pathLen := len(path)

		for _, p := range prefixes {
			if pl := len(p); pathLen >= pl && path[:pl] == p {
				return true
			}
		}
		return false
	}
}
