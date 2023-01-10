package oidc

import "fmt"

const (
	ErrCodeInvalidRequest          = "invalid_request"
	ErrCodeUnauthorizedClient      = "unauthorized_client"
	ErrCodeAccessDenied            = "access_denied"
	ErrCodeUnsupportedResponseType = "unsupported_response_type"
	ErrCodeInvalidScope            = "invalid_scope"
	ErrCodeServerError             = "server_error"
	ErrCodeTemporarilyUnavailable  = "temporarily_unavailable"
)

// Error represents error defined in RFC6749.
type Error struct {
	// Code is the error code. REQUIRED.
	Code string `json:"error,omitempty"`
	// Message is the human-readable description of the error. OPTIONAL, but RECOMMENDED.
	Message string `json:"error_description,omitempty"`
	// URI is the URL referencing a web page displaying additional details about the error. It is usually
	// assigned by the HTTP layer. Internal components can ignore this field. OPTIONAL.
	URI string `json:"error_uri,omitempty"`
}

func (e Error) Error() string {
	if len(e.Message) == 0 {
		return e.Code
	}
	return fmt.Sprintf("%s[%s]", e.Code, e.Message)
}
