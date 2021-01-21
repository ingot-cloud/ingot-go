package id

// Generator id生成器
type Generator interface {
	// 生成ID
	NextID() (int64, error)
}
