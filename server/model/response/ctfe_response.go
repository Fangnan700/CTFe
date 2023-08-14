package response

/*
	响应代码对照：
	900:	操作成功
	901：	数据为空
	902：	未授权
	903：	系统异常
*/

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
