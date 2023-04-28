package api

// APIError is the response body for the OpenAI API when an error occurs
type APIError struct {
	Message string `json:"message"`
	Type    string `json:"type"`
	Param   string `json:"param"`
	Code    string `json:"code"`
}

// satisfy the error interface
func (e *APIError) Error() string {
	return e.Message
}
