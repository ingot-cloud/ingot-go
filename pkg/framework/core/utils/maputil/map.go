package maputil

// CopyStringInterfaceMap 复制 map[string]any
func CopyStringInterfaceMap(origin map[string]any) map[string]any {
	copy := make(map[string]any)
	for k, v := range origin {
		copy[k] = v
	}
	return copy
}

// CopyStringStringMap 复制 map[string]string
func CopyStringStringMap(origin map[string]string) map[string]string {
	copy := make(map[string]string)
	for k, v := range origin {
		copy[k] = v
	}
	return copy
}
