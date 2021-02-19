package maputil

// CopyStringInterfaceMap 复制 map[string]interface{}
func CopyStringInterfaceMap(origin map[string]interface{}) map[string]interface{} {
	copy := make(map[string]interface{})
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
