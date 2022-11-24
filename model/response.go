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
