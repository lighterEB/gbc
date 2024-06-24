package response

type Response[T any] struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data T      `json:"data"`
}

type Code struct {
	Code int
	Msg  string
}

const (
	SuccessCode       = 200
	NotFoundCode      = 404
	InternalErrorCode = 500
)

func NewResponseCode(code int, msg string) *Code {
	return &Code{Code: code, Msg: msg}
}

var (
	Success  = NewResponseCode(SuccessCode, "Success")
	NotFound = NewResponseCode(NotFoundCode, "Not Found")
	Internal = NewResponseCode(InternalErrorCode, "Internal Error")
)

func NewSuccessResponse[T any](data T) *Response[T] {
	return &Response[T]{
		Code: SuccessCode,
		Msg:  "Success",
		Data: data,
	}
}

func NewErrorResponse[T any](code int, msg string) *Response[T] {
	return &Response[T]{
		Code: code,
		Msg:  msg,
		Data: *new(T),
	}
}
