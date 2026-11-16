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
	Public bool `json:"public"`
}

var Echoes = []*Echo{
	{Id: "1", Title: "Hello", Body: "This is the text", Public: true},
	{Id: "2", Title: "Another echo", Body: "I am another echo", Public: true},
}


// The request payload for the Echo data model
type EchoRequest struct{
	*Echo
}

// Bind method for render.Binder interface
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
