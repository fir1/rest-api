package handlers

type ErrorResponse struct {
	ErrorMessage string `json:"error_message"`
}

func (e ErrorResponse) Error() string {
	return e.ErrorMessage
}
