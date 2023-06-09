package appErrors

import "encoding/json"

var (
	ErrNotFound = NewAppError(nil, "not found", "", "US-000003")
)

type AppError struct {
	Err              error  `json:"app-errors"`
	Message          string `json:"message,omitempty"`
	DeveloperMessage string `json:"developer_message,omitempty"`
	Code             string `json:"code,omitempty"`
}

func (e *AppError) Error() string {
	return e.Message
}

func (e *AppError) Marshal() []byte {
	marshal, err := json.Marshal(e)
	if err != nil {
		return nil
	}
	return marshal
}

func NewAppError(err error, message, developerMessage, code string) *AppError {
	return &AppError{
		Err:              err,
		Message:          message,
		DeveloperMessage: developerMessage,
		Code:             code,
	}
}

func SystemError(err error) *AppError {
	return NewAppError(err, "internal system error", err.Error(), "US-000000")
}
