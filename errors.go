package okta

import "fmt"

type ErrorResponse struct {
	ErrorCode    string `json:"errorCode"`
	ErrorSummary string `json:"errorSummary"`
	ErrorLink    string `json:"errorLink"`
	ErrorID      string `json:"errorId"`
	ErrorCauses  []struct {
		ErrorSummary string `json:"errorSummary"`
	} `json:"errorCauses"`
}

type errorString struct {
	s string
}

type errorResponse struct {
	HTTPCode int
	Response ErrorResponse
	Endpoint string
}

func (e *errorResponse) Error() string {
	return fmt.Sprintf("Error hitting api endpoint %s %s", e.Endpoint, e.Response.ErrorCode)
}