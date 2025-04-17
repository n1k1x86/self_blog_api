package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Handle505Error(err error, w http.ResponseWriter) {
	err505, _ := json.Marshal(fmt.Sprintf(`{"error": %s}`, err.Error()))
	w.Header().Add("Content-Type", "application/json")
	w.Write(err505)
}
