package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

const apiName = "/api/v1/"

type Handler struct {
	Handler *http.ServeMux
}

func HelloPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, "+apiName)
}

func SetHandlersTags(r *gin.Engine, dbpool *pgxpool.Pool) {
	r.GET(apiName+"tags", GetTags(dbpool))
	r.GET(apiName+"tags/:id", GetTagByID(dbpool))

	r.PUT(apiName+"tags/:id", EditTag(dbpool))
	r.POST(apiName+"tags", AddTag(dbpool))
	r.DELETE(apiName+"tags/:id", DeleteTag(dbpool))
}
