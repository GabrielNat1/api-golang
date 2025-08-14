package errors

import "fmt"

type ErrorType string

const (
	NotFound      ErrorType = "NOT_FOUND"
	ValidationErr ErrorType = "VALIDATION_ERROR"
	DatabaseErr   ErrorType = "DATABASE_ERROR"
	InternalErr   ErrorType = "INTERNAL_ERROR"
	Unauthorized  ErrorType = "UNAUTHORIZED"
	BadRequest    ErrorType = "BAD_REQUEST"
)

type AppError struct {
	Type    ErrorType
	Message string
	Err     error
}

func (e AppError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %s (%v)", e.Type, e.Message, e.Err)
	}
	return fmt.Sprintf("%s: %s", e.Type, e.Message)
}

func NewNotFound(message string, err error) error {
	return AppError{
		Type:    NotFound,
		Message: message,
		Err:     err,
	}
}

func NewValidationError(message string, err error) error {
	return AppError{
		Type:    ValidationErr,
		Message: message,
		Err:     err,
	}
}

func NewDatabaseError(message string, err error) error {
	return AppError{
		Type:    DatabaseErr,
		Message: message,
		Err:     err,
	}
}

func NewInternalError(message string, err error) error {
	return AppError{
		Type:    InternalErr,
		Message: message,
		Err:     err,
	}
}
