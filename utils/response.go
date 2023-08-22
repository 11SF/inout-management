package response

import "fmt"

type Response struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type Error struct {
	Code string
	Err  error
}

func (r *Error) Error() string {
	return fmt.Sprintf("status %v: err %v", r.Code, r.Err)
}

func NewResponse(httpStatusCode int, data interface{}) (int, Response) {
	return httpStatusCode, Response{
		Code:    "00000",
		Message: "",
		Data:    nil,
	}
}

func NewResponseError(httpStatusCode int, err Error) (int, Response) {
	return httpStatusCode, Response{
		Code:    err.Code,
		Message: err.Err.Error(),
		Data:    nil,
	}

}

func WriteError(code string, err error) *Error {
	return &Error{
		Code: code,
		Err:  err,
	}
}
