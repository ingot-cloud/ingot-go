package clientdetails

// NilClientdetails 空实现
type NilClientdetails struct{}

// LoadClientByClientID 根据 clientID 获取客户端详细信息
func (*NilClientdetails) LoadClientByClientID(string) (ClientDetails, error) {
	return nil, nil
}
