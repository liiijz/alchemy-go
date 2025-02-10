package alchemy

import "fmt"

// AlchemyError represents an error from the Alchemy API
type AlchemyError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

// Error implements the error interface
func (e *AlchemyError) Error() string {
	if e.Data != nil {
		return fmt.Sprintf("alchemy error %d: %s (data: %v)", e.Code, e.Message, e.Data)
	}
	return fmt.Sprintf("alchemy error %d: %s", e.Code, e.Message)
}

// Common error constructors
func newAPIError(code int, message string) *AlchemyError {
	return &AlchemyError{
		Code:    code,
		Message: message,
	}
}

// IsAlchemyError checks if an error is an AlchemyError
func IsAlchemyError(err error) bool {
	_, ok := err.(*AlchemyError)
	return ok
}
