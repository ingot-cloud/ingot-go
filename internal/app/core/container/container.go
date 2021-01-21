package container

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

// ContainerSet inject
var ContainerSet = wire.NewSet(wire.Struct(new(Container), "*"))

// Container for app
type Container struct {
	Engine *gin.Engine
}
