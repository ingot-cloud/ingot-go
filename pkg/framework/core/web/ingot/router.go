package ingot

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Router gin.Router扩展
type Router struct {
	ginRouter *gin.RouterGroup
}

// NewRouter 实例Router
func NewRouter(routerGroup *gin.RouterGroup) *Router {
	return &Router{
		ginRouter: routerGroup,
	}
}

// Use adds middleware to the group, see example code in GitHub.
func (router *Router) Use(middleware ...any) gin.IRoutes {
	return router.ginRouter.Use(transformHandlers(middleware...)...)
}

// Group creates a new router group. You should add all the routes that have common middlewares or the same path prefix.
// For example, all the routes that use a common middleware for authorization could be grouped.
func (router *Router) Group(relativePath string, handlers ...any) *Router {
	return NewRouter(router.ginRouter.Group(relativePath, transformHandlers(handlers...)...))
}

// BasePath returns the base path of router group.
// For example, if v := router.Group("/rest/n/v1/api"), v.BasePath() is "/rest/n/v1/api".
func (router *Router) BasePath() string {
	return router.ginRouter.BasePath()
}

// Handle registers a new request handle and middleware with the given path and method.
// The last handler should be the real handler, the other ones should be middleware that can and should be shared among different routes.
// See the example code in GitHub.
//
// For GET, POST, PUT, PATCH and DELETE requests the respective shortcut
// functions can be used.
//
// This function is intended for bulk loading and to allow the usage of less
// frequently used, non-standardized or custom methods (e.g. for internal
// communication with a proxy).
func (router *Router) Handle(httpMethod, relativePath string, handlers ...any) gin.IRoutes {
	return router.ginRouter.Handle(httpMethod, relativePath, transformHandlers(handlers...)...)
}

// POST is a shortcut for router.Handle("POST", path, handle).
func (router *Router) POST(relativePath string, handlers ...any) gin.IRoutes {
	return router.ginRouter.POST(relativePath, transformHandlers(handlers...)...)
}

// GET is a shortcut for router.Handle("GET", path, handle).
func (router *Router) GET(relativePath string, handlers ...any) gin.IRoutes {
	return router.ginRouter.GET(relativePath, transformHandlers(handlers...)...)
}

// DELETE is a shortcut for router.Handle("DELETE", path, handle).
func (router *Router) DELETE(relativePath string, handlers ...any) gin.IRoutes {
	return router.ginRouter.DELETE(relativePath, transformHandlers(handlers...)...)
}

// PATCH is a shortcut for router.Handle("PATCH", path, handle).
func (router *Router) PATCH(relativePath string, handlers ...any) gin.IRoutes {
	return router.ginRouter.PATCH(relativePath, transformHandlers(handlers...)...)
}

// PUT is a shortcut for router.Handle("PUT", path, handle).
func (router *Router) PUT(relativePath string, handlers ...any) gin.IRoutes {
	return router.ginRouter.PUT(relativePath, transformHandlers(handlers...)...)
}

// OPTIONS is a shortcut for router.Handle("OPTIONS", path, handle).
func (router *Router) OPTIONS(relativePath string, handlers ...any) gin.IRoutes {
	return router.ginRouter.OPTIONS(relativePath, transformHandlers(handlers...)...)
}

// HEAD is a shortcut for router.Handle("HEAD", path, handle).
func (router *Router) HEAD(relativePath string, handlers ...any) gin.IRoutes {
	return router.ginRouter.HEAD(relativePath, transformHandlers(handlers...)...)
}

// Any registers a route that matches all the HTTP methods.
// GET, POST, PUT, PATCH, HEAD, OPTIONS, DELETE, CONNECT, TRACE.
func (router *Router) Any(relativePath string, handlers ...any) gin.IRoutes {
	return router.ginRouter.Any(relativePath, transformHandlers(handlers...)...)
}

// StaticFile registers a single route in order to serve a single file of the local filesystem.
// router.StaticFile("favicon.ico", "./resources/favicon.ico")
func (router *Router) StaticFile(relativePath, filepath string) gin.IRoutes {
	return router.ginRouter.StaticFile(relativePath, filepath)
}

// Static serves files from the given file system root.
// Internally a http.FileServer is used, therefore http.NotFound is used instead
// of the Router's NotFound handler.
// To use the operating system's file system implementation,
// use :
//
//	router.Static("/static", "/var/www")
func (router *Router) Static(relativePath, root string) gin.IRoutes {
	return router.ginRouter.Static(relativePath, root)
}

// StaticFS works just like `Static()` but a custom `http.FileSystem` can be used instead.
// Gin by default user: gin.Dir()
func (router *Router) StaticFS(relativePath string, fs http.FileSystem) gin.IRoutes {
	return router.ginRouter.StaticFS(relativePath, fs)
}
