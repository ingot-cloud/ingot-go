package dto

// Pagination for request params
type Pagination struct {
	Size    int `json:"size" form:"size"`
	Current int `json:"current" form:"current"`
}
