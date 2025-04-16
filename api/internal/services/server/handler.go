package server

import (
	"fmt"
	"net/http"
)

type Handler struct {
	Handler *http.ServeMux
}

func HelloPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, web page")
}

func CreateNewHandler() *Handler {
	handler := http.NewServeMux()
	handler.HandleFunc("/", HelloPage)
	return &Handler{
		Handler: handler,
	}
}
