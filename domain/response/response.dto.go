package response

type response struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
	Error  interface{} `json:"error"`
}

func NewResponse(status int, data interface{}) response {
	return response{
		Status: status,
		Data:   data,
		Error:  nil,
	}
}
func NewError(status int, err interface{}) response {
	return response{
		Status: status,
		Data:   nil,
		Error:  err,
	}
}
