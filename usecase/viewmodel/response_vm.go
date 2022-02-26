package viewmodel

//ResponseErrorVM return response error
type ResponseErrorVM struct {
	Messages interface{} `json:"messages"`
}

//ResponseSuccessVM return response success
type ResponseSuccessVM struct {
	Data interface{} `json:"data"`
	Meta interface{} `json:"meta"`
}
