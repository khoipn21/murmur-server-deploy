package model

type ErrorsResponse struct {
	Errors []FieldError `json:"errors"`
}

// FieldError is used to help extract validation errors
type FieldError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// ErrorResponse holds a custom error for the application
type ErrorResponse struct {
	Error HttpError `json:"error"`
}

// HttpError returns the Http error type and the specific message
type HttpError struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}
