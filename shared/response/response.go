package response

// type APIResponse struct {}
type SuccessResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
	Data    any    `json:"data,omitempty"`
}
type ErrorResponse struct {
	Success bool      `json:"success"`
	Message string    `json:"message,omitempty"`
	Error   ErrorBody `json:"error"`
}
type ErrorBody struct {
	Code string `json:"code"`
}
