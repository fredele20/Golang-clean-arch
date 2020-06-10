package router

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type muxRouter struct {}

func NewMuxRouter() Router {
	return &muxRouter{}
}

var muxDispatcher = mux.NewRouter()

func (*muxRouter) GET(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	muxDispatcher.HandleFunc(uri, f).Methods("GET")
}
func (*muxRouter) POST(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	muxDispatcher.HandleFunc(uri, f).Methods("POST")
}
func (*muxRouter) SERVER(port string) {
	fmt.Printf("Mux HTTP server running on port %v", port)
	http.ListenAndServe(port, muxDispatcher)
}

