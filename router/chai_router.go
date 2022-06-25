package router

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type chaiRouter struct {
}

func NewChaiRouter() Router {
	return &chaiRouter{}
}

var (
	chaiDispatcher = chi.NewRouter()
)

func (*chaiRouter) Get(url string, f func(w http.ResponseWriter, r *http.Request)) {
	chaiDispatcher.Get(url, f)
}

func (*chaiRouter) Post(url string, f func(w http.ResponseWriter, r *http.Request)) {
	chaiDispatcher.Post(url, f)
}

func (*chaiRouter) Patch(url string, f func(w http.ResponseWriter, r *http.Request)) {
	chaiDispatcher.Patch(url, f)
}

func (*chaiRouter) Put(url string, f func(w http.ResponseWriter, r *http.Request)) {
	chaiDispatcher.Put(url, f)
}

func (*chaiRouter) Delete(url string, f func(w http.ResponseWriter, r *http.Request)) {
	chaiDispatcher.Delete(url, f)
}

func (*chaiRouter) Serve(port string) {
	fmt.Printf("Server running on port %s\n", port)
	panic(http.ListenAndServe(port, chaiDispatcher))
}
