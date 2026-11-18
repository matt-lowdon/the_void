package main

import(
	"net/http"
	"errors"

	"github.com/go-chi/render"
)

//An Echo is a post. This is the basic structure
type Echo struct{
	Id string `json:"id"`
	Title string `json:"title"`
	Body string `json:"body"`
	ReportCount int `json:"report_count"`
}

// This is the database for the time being
var Echoes = []*Echo{
	{Id: "1", Title: "Hello", Body: "This is the text"},
	{Id: "2", Title: "Another echo", Body: "I am another echo"},
}


// The request payload struct for the Echo data model
type EchoRequest struct{
	*Echo

	Public bool `json:"public"`
}

// Bind method for render.Binder interface
// Safe to ignore checks on Public. If it's missing it safely defaults to false
func (e *EchoRequest) Bind(r *http.Request) error {
	if e.Echo == nil {
		return errors.New("missing required Echo field")
	}

	return nil
}

// The response payload for the Echo data model
type EchoResponse struct{
	*Echo
}

// Render method for render.Renderer interface
func (e *EchoResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

// Consistent interface for returning 
type SubmissionResponse struct{
	Received bool `json:"received"`
	Saved bool `json:"saved"`
	Echo *Echo `json:"echo,omitempty"`
}

func (s *SubmissionResponse) Render(w http.ResponseWriter, r *http.Request) error {
	if s.Saved {
		render.Status(r, http.StatusCreated)
	} else {
		render.Status(r, http.StatusAccepted)
	}
	return nil
}

type ErrResponse struct{
	Err error `json:"-"` // low-level runtime error
	HTTPStatusCode int `json:"-"` // http response status code

	StatusText string `json:"status"` //user-level error message
	AppCode int64 `json:"code,omitempty"` //application-specific error code
	ErrorText string `json:"error,omitempty"`
}

func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

func ErrInvalidRequest(err error) render.Renderer {
	return &ErrResponse{
		Err: err,
		HTTPStatusCode: 400,
		StatusText: "Invalid request",
		ErrorText: err.Error(),
	}
}

func ErrRender(err error) render.Renderer {
	return &ErrResponse{
		Err: err,
		HTTPStatusCode: 422,
		StatusText: "Error rendering response",
		ErrorText: err.Error(),
	}
}

var ErrNotFound = &ErrResponse{HTTPStatusCode: 404, StatusText: "Resource not found"}
