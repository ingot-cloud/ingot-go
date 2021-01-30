package request

// BaseRequest 基础请求信息
type BaseRequest interface {
	GetClientID() string
	GetScope() []string
	GetRequestParameters() map[string]string
}

// BaseRequestField 请求基础信息
type BaseRequestField struct {
	ClientID          string
	Scope             []string
	RequestParameters map[string]string
}

// GetClientID 获取ClientID
func (r *BaseRequestField) GetClientID() string {
	return r.ClientID
}

// GetScope 获取scope
func (r *BaseRequestField) GetScope() []string {
	return r.Scope
}

// GetRequestParameters 获取请求参数
func (r *BaseRequestField) GetRequestParameters() map[string]string {
	return r.RequestParameters
}
