package server

import "net/http"

type ServerHTTP struct {
	Server *http.Server
}

func CreateNewServer() *ServerHTTP {
	handler := CreateNewHandler()
	server := &http.Server{
		Addr:    ":8080",
		Handler: handler.Handler,
	}
	return &ServerHTTP{
		Server: server,
	}
}
