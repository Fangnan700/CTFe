package models

type Response struct {
	Code int         `json:"code"`
	Body interface{} `json:"body"`
}

func NewResponse(code int, body interface{}) Response {
	return Response{
		Code: code,
		Body: body,
	}
}
