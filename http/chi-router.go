package router

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

type chiRouter struct{}

var (
	chiDispatcher = chi.NewRouter()
)

//NewChiRouter
func NewChiRouter() Router {
	return &chiRouter{}
}

func (*chiRouter) REQUEST(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	chiDispatcher.HandleFunc(uri, f)
}

func (*chiRouter) GET(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	chiDispatcher.Get(uri, f)
}

func (*chiRouter) POST(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	chiDispatcher.Post(uri, f)
}

func (*chiRouter) DELETE(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	chiDispatcher.Delete(uri, f)
}

func (*chiRouter) SERVE(port string) {
	fmt.Println("Chi HTTP server listing on port", port)
	_ = http.ListenAndServe(port, chiDispatcher)
}
