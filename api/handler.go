package main

import (
	"net/http"
	"fmt"
	"math/rand"

	"github.com/go-chi/render"
)

type EchoHandler struct{
}

func (eh *EchoHandler) ListEchoes(w http.ResponseWriter, r *http.Request) {
	if err := render.RenderList(w, r, RenderEchoesList()); err != nil {
		render.Render(w, r, ErrRender(err))
		return
	}
}

func (eh *EchoHandler) CreateEcho(w http.ResponseWriter, r *http.Request) {
	data := &EchoRequest{}
	err := render.Bind(r, data); if err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}

	if data.Public == false {
		render.Render(w, r, &SubmissionResponse{Received: true, Saved: false})
		return
	}

	echo := data.Echo
	newId, err := dbNewEcho(echo)
	if err != nil{
		render.Render(w, r, ErrRender(err))
		return
	}
	echo.Id = newId
	
	render.Render(w, r, &SubmissionResponse{Received: true, Saved: true, Echo: echo})
}

func RenderEchoesList() []render.Renderer{
	list := []render.Renderer{}
	for _, echo := range Echoes {
		list = append(list, &EchoResponse{Echo: echo})
	}

	return list
}

func dbNewEcho(e *Echo) (string, error) {
	e.Id = fmt.Sprintf("%d", rand.Intn(100) + 10)
	Echoes = append(Echoes, e)
	return e.Id, nil
}
