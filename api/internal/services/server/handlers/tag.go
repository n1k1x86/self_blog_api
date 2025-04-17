package handler

import (
	"api/config"
	"api/db/models"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type TagsHandlerStruct struct {
	cfg config.BlogDBConfig
}

func (h *TagsHandlerStruct) AddTag(w http.ResponseWriter, r *http.Request) {
	bodyData, err := io.ReadAll(r.Body)
	if err != nil {
		Handle505Error(err, w)
		return
	}
	id, err := models.AddTag(h.cfg, bodyData)
	if err != nil {
		Handle505Error(err, w)
		return
	}
	successRes, err := json.Marshal(fmt.Sprintf(`{"tag_id": %d}`, id))
	if err != nil {
		Handle505Error(err, w)
		return
	}
	w.Write(successRes)
	w.Header().Add("Content-Type", "application/json")
}

func (h *TagsHandlerStruct) EditTag(w http.ResponseWriter, r *http.Request) {
}

func (h *TagsHandlerStruct) DeleteTag(w http.ResponseWriter, r *http.Request) {
}

func (h *TagsHandlerStruct) GetTags(w http.ResponseWriter, r *http.Request) {
	result, err := models.GetTags(h.cfg)
	if err != nil {
		Handle505Error(err, w)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.Write(result)
}

func (h *TagsHandlerStruct) GetTagByID(w http.ResponseWriter, r *http.Request) {
}

func (h *TagsHandlerStruct) TagsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		h.GetTags(w, r)
		return
	}
	if r.Method == http.MethodPost {
		h.AddTag(w, r)
		return
	}
}

func (h *TagsHandlerStruct) TagHandler(w http.ResponseWriter, r *http.Request) {
}
