package main

import(
	"net/http"
	"errors"
)

type Echo struct{
	Id string `json:"id"`
	Title string `json:"title"`
	Body string `json:"body"`
	Public bool `json:"public"`
}

var echoes = []*Echo{
	{Id: "1", Title: "Hello", Body: "This is the text", Public: true},
	{Id: "2", Title: "Another echo", Body: "I am another echo", Public: true},
}

type EchoRequest struct{
	*Echo
}

func (e *EchoRequest) Bind(r *http.Request) error {
	if e.Echo == nil {
		return errors.New("missing required Echo field")
	}

	return nil
}
