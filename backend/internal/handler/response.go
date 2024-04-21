package handler

type ErrorCode struct{}

type Response struct {
	Payload interface{} `json:"payload"`
	Message string      `json:"message"`
}

func NewErrorResponse(err string) *Response {
	return &Response{
		Message: err,
	}
}

func NewResponse(payload interface{}) *Response {
	return &Response{
		Payload: payload,
		Message: "Operation completed successfully.",
	}
}
