package rest_err

import "net/http"

type RestErr struct {
	Message string `json:"message"`
	Err     string `json:"error"`
	Code    int    `json:"code"`
}

func (r *RestErr) Error() string {
	return r.Message
}

func NewRestErr(message string, err string, code int) *RestErr {
	return &RestErr{
		Message: message,
		Err:     err,
		Code:    code,
	}
}

func NewBadResquestErr(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "bad request",
		Code:    http.StatusBadRequest,
	}
}

func NewInternalServerErr() *RestErr {
	return &RestErr{Message: "Ocorreu um erro inesperado", Err: "internal server error", Code: http.StatusInternalServerError}
}
