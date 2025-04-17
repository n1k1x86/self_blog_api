package handler

import (
	"api/config"
	"fmt"
	"net/http"
)

const apiName = "/api/v1/"

type Handler struct {
	Handler *http.ServeMux
}

func HelloPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, "+apiName)
}

func AddArticle(w http.ResponseWriter, r *http.Request) {
}

func EditArticle(w http.ResponseWriter, r *http.Request) {
}

func DeleteArticle(w http.ResponseWriter, r *http.Request) {
}

func GetArticleByID(w http.ResponseWriter, r *http.Request) {
}

func GetArticleList(w http.ResponseWriter, r *http.Request) {
}

func CreateNewHandler(cfg config.BlogDBConfig) *Handler {
	handler := http.NewServeMux()
	tagHandler := TagsHandlerStruct{cfg: cfg}

	handler.HandleFunc(apiName, HelloPage)
	handler.HandleFunc(apiName+"tags/", tagHandler.TagsHandler)
	return &Handler{
		Handler: handler,
	}
}
