package domain

// ErrorResponse struct is used to format error response
type ErrorResponse struct {
	Message string `json:"message"`
}

// SuccessResponse struct is used to format success response
type SuccessResponse struct {
	Message string `json:"message"`
}
