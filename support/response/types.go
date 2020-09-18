package response

// D is a shortcut for map[string]interface{}
type D map[string]interface{}

// R is the Response struct
type R struct {
	Code    int    `json:"code"`
	Data    D      `json:"data"`
	Message string `json:"message"`
}
