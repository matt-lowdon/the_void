package main

import (
	"net/http"

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

func RenderEchoesList() []render.Renderer{
	list := []render.Renderer{}
	for _, echo := range Echoes {
		list = append(list, &EchoResponse{Echo: echo})
	}

	return list
}
