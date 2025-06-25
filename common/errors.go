package common

import "errors"

// Input validation errors
var (
	ErrInvalidInput    = errors.New("invalid input data")
	ErrMissingFullname = errors.New("full name is required")
	ErrMissingEmail    = errors.New("email is required")
	ErrInvalidEmail    = errors.New("invalid email format")
	ErrInvalidUserID   = errors.New("invalid user ID")
)

// ResponseData represents the standard API response structure
type ResponseData struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Error   string      `json:"error,omitempty"`
}
