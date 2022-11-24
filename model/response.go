package model

type Response struct {
	Code    int     `json:"-"`
	Data    []Movie `json:"data"`
	Message string  `json:"message"`
}

type ErrorService struct {
	Code    int
	Message string
}

func (e ErrorService) Error() string {
	return e.Message
}

func NewErrorServiceBuilder() *ErrorService {
	return &ErrorService{
		Code:    500,
		Message: "internal service error",
	}
}

func (b *ErrorService) SetCode(code int) *ErrorService {
	b.Code = code
	return b
}

func (b *ErrorService) SetMessage(message string) *ErrorService {
	b.Message = message
	return b
}
