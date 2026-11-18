package main

import (
	"net/http"
	"fmt"
	"math/rand"
	"context"
	"errors"
	"sync"

	"github.com/go-chi/render"
	"github.com/go-chi/chi/v5"
)

// Handler struct for requests. Easier to pass values in (e.g postgres connections)
type EchoHandler struct{
	mu sync.RWMutex
}

// Listing all echos
// TODO when switching to Postgres, limit response and implement pagination
func (eh *EchoHandler) ListEchoes(w http.ResponseWriter, r *http.Request) {
	eh.mu.RLock()
	defer eh.mu.RUnlock()
	if err := render.RenderList(w, r, RenderEchoesList()); err != nil {
		render.Render(w, r, ErrRender(err))
		return
	}
}

// Creating a post. Filters depending on whether the post is public or not
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
	eh.mu.Lock()
	newId, err := dbNewEcho(echo)
	eh.mu.Unlock()
	if err != nil{
		render.Render(w, r, ErrRender(err))
		return
	}
	echo.Id = newId
	
	render.Render(w, r, &SubmissionResponse{Received: true, Saved: true, Echo: echo})
}

// Context is responsible for loading specified Echo from DB
func (eh *EchoHandler) EchoContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var err error
		var echo *Echo

		if echoId := chi.URLParam(r, "echoId"); echoId != "" {
			eh.mu.Lock()
			echo, err = dbGetEcho(echoId)
			eh.mu.Unlock()
		} else {
			render.Render(w, r, ErrNotFound)
			return
		}

		if err != nil {
			render.Render(w, r, ErrNotFound)
			return
		}

		ctx := context.WithValue(r.Context(), "echo", echo)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// Return specified Echo
func (eh *EchoHandler) GetEcho(w http.ResponseWriter, r *http.Request) {
	echo := r.Context().Value("echo").(*Echo)
	if err := render.Render(w, r, &EchoResponse{Echo: echo}); err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}
}

// Delete specified Echo
func (eh *EchoHandler) DeleteEcho(w http.ResponseWriter, r *http.Request) {
	echo := r.Context().Value("echo").(*Echo)
	if err := dbDeleteEcho(echo.Id); err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}
	render.NoContent(w, r)
}

// Function for converting list of Echo structs to EchoResponse for rendering
// Will remain useful when loading from Postgres
func RenderEchoesList() []render.Renderer{
	list := []render.Renderer{}
	for _, echo := range Echoes {
		list = append(list, &EchoResponse{Echo: echo})
	}

	return list
}

// Creates a new Echo and saves it to the database
// TODO when switching to Postgres, update ID to UUIDv7. API will be responsible for creation
func dbNewEcho(e *Echo) (string, error) {
	e.Id = fmt.Sprintf("%d", rand.Intn(100) + 10)
	Echoes = append(Echoes, e)
	return e.Id, nil
}

func dbGetEcho(id string) (*Echo, error) {
	for _, echo := range Echoes {
		if id == echo.Id {
			return echo, nil
		}
	}

	return nil, errors.New("no match for provided id")
}

func dbDeleteEcho(id string) error {
	for i := range Echoes {
		if Echoes[i].Id == id {
			Echoes = append(Echoes[:i], Echoes[i+1:]...)
			return nil
		}
	}
	return errors.New("echo not found")
}
