package wallet

// RPC response structure
type JsonResponse struct {
	Status  bool        `json:"status"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

// The Error function sets response status false
// and renders an error message
func (response *JsonResponse) Error(err error) {
	response.Status = false
	response.Message = err.Error()
}

// The render function renders response data
// and sets response status true
func (response *JsonResponse) Render(data interface{}) {
	response.Status = true
	response.Data = data
}
