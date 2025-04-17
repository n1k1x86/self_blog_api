package server

import (
	"api/config"
	handler "api/internal/services/server/handlers"
	"net/http"
)

type ServerHTTP struct {
	Server *http.Server
}

func CreateNewServer(cfg config.BlogDBConfig) *ServerHTTP {
	handler := handler.CreateNewHandler(cfg)
	server := &http.Server{
		Addr:    ":8080",
		Handler: handler.Handler,
	}
	return &ServerHTTP{
		Server: server,
	}
}
