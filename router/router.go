package router

import "net/http"

type Router interface {
	Get(url string, f func(w http.ResponseWriter, r *http.Request))
	Post(url string, f func(w http.ResponseWriter, r *http.Request))
	Patch(url string, f func(w http.ResponseWriter, r *http.Request))
	Put(url string, f func(w http.ResponseWriter, r *http.Request))
	Delete(url string, f func(w http.ResponseWriter, r *http.Request))
	Serve(port string)
}
