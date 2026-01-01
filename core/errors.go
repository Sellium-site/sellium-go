package core

import "fmt"

type APIErrorBody struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type APIError struct {
	Status  int
	Code    string
	Message string
	Raw     []byte
}

func (e *APIError) Error() string {
	if e.Code != "" {
		return fmt.Sprintf("sellium API error (%d) %s: %s", e.Status, e.Code, e.Message)
	}
	return fmt.Sprintf("sellium API error (%d): %s", e.Status, e.Message)
}
