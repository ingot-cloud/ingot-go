package core

import "sync"

// Loader interface
type Loader interface {
	Init() (interface{}, error)
}

var coreLoader *CoreLoader
var once sync.Once

// GetCoreLoader 获取核心加载器
func GetCoreLoader() *CoreLoader {
	once.Do(func() {
		coreLoader = new(CoreLoader)
	})
	return coreLoader
}

type CoreLoader struct {
	Loaders []*Loader
}

func (c *CoreLoader) Init() (interface{}, error) {
	for _, loader := range c.Loaders {
		var a Loader = *loader
		a.Init()
	}
	return c, nil
}

func (c *CoreLoader) Register(loader *Loader) {
	c.Loaders = append(c.Loaders, loader)
}
