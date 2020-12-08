package yahoofinance

import (
	"bytes"
	"fmt"
	"net/http"
)

// Error contains an error response from the server.
type Error struct {
	// Code is the HTTP response status code and will always be populated.
	Code int `json:"code,omitempty"`
	// Message is the server response message and is only populated when
	// explicitly referenced by the JSON server response.
	Message string `json:"message,omitempty"`
	// Body is the raw response returned by the server.
	// It is often but not always JSON, depending on how the request fails.
	Body string
	// Header contains the response header fields from the server.
	Header http.Header
}

func (e *Error) Error() string {
	if e.Message == "" {
		return fmt.Sprintf("API: got HTTP response code %d with body: %v", e.Code, e.Body)
	}
	var buf bytes.Buffer
	fmt.Fprintf(&buf, "API: Error %d: ", e.Code)
	if e.Message != "" {
		fmt.Fprintf(&buf, "%s", e.Message)
	}

	return buf.String()
}

type errorReply struct {
	Chart struct {
		Error struct {
			Code        string `json:"code"`
			Description string `json:"description"`
		} `json:"error"`
	} `json:"chart"`
}
