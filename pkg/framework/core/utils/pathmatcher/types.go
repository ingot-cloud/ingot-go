package pathmatcher

// PathMatcher 路径匹配器
type PathMatcher interface {
	// 匹配
	Match(pattern string, path string) bool
}
