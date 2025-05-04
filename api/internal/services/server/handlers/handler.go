package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

const apiName = "/api/v1/"

func SetHandlersTags(r *gin.Engine, dbpool *pgxpool.Pool) {
	r.GET(apiName+"tags", GetTags(dbpool))
	r.GET(apiName+"tags/:id", GetTagByID(dbpool))

	r.PUT(apiName+"tags/:id", EditTag(dbpool))

	r.POST(apiName+"tags", AddTag(dbpool))

	r.DELETE(apiName+"tags/:id", DeleteTag(dbpool))
}

func SetHandlersArticle(r *gin.Engine, dbpool *pgxpool.Pool) {
	r.GET(apiName+"articles", GetArticles(dbpool))
	r.GET(apiName+"articles/:id", GetArticleByID(dbpool))

	r.PUT(apiName+"articles/:id", EditArticle(dbpool))

	r.POST(apiName+"articles", AddArticle(dbpool))

	r.DELETE(apiName+"articles/:id", DeleteArticle(dbpool))
}
