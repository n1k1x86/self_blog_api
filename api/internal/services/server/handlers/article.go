package handler

import (
	"api/db/models"
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func AddArticle(dbpool *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c.Request.Context(), time.Second*3)
		defer cancel()
		article := &models.ArticleToSave{}
		data, err := io.ReadAll(c.Request.Body)
		if err != nil {
			Handle505Error(c, err)
			return
		}
		err = json.Unmarshal(data, article)
		if err != nil {
			Handle400Error(c)
			return
		}
		err = models.AddArticle(ctx, dbpool, article)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				Handle404Error(c, ErrNoTag{TagID: article.TagID})
				return
			}
			Handle505Error(c, err)
			return
		}
		c.Status(201)
	}
}

func EditArticle(dbpool *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c.Request.Context(), time.Second*3)
		defer cancel()

		idQuery, ok := c.Params.Get("id")
		if !ok {
			Handle400Error(c)
			return
		}
		id, err := strconv.Atoi(idQuery)
		if err != nil {
			Handle505Error(c, err)
			return
		}
		body, err := io.ReadAll(c.Request.Body)
		if err != nil {
			fmt.Println("ERROR OSHIBKA1")
			Handle505Error(c, err)
			return
		}
		article := models.ArticleToSave{}
		fmt.Println("TYTA2")
		err = json.Unmarshal(body, &article)
		if err != nil {
			fmt.Println("ERROR OSHIBKA2")
			Handle505Error(c, err)
			return
		}
		fmt.Println(article)
		err = models.EditArticle(ctx, dbpool, &article, int64(id))
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				Handle404Error(c, err)
				return
			}
			Handle505Error(c, err)
			return
		}
		c.Status(http.StatusOK)
	}
}

func DeleteArticle(dbpool *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c.Request.Context(), time.Second*3)
		defer cancel()

		idQuery, ok := c.Params.Get("id")
		if !ok {
			Handle400Error(c)
			return
		}
		id, err := strconv.Atoi(idQuery)
		if err != nil {
			Handle505Error(c, err)
			return
		}
		err = models.DeleteArticle(ctx, dbpool, int64(id))
		if err != nil {
			Handle505Error(c, err)
			return
		}
		c.Status(http.StatusNoContent)
	}
}

func GetArticleByID(dbpool *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c.Request.Context(), time.Second*3)
		defer cancel()

		idQuery, ok := c.Params.Get("id")
		if !ok {
			Handle400Error(c)
			return
		}
		id, err := strconv.Atoi(idQuery)
		if err != nil {
			Handle505Error(c, err)
			return
		}

		article, err := models.GetArticleByID(ctx, dbpool, int64(id))
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				Handle404Error(c, err)
				return
			}
			Handle505Error(c, err)
			return
		}
		c.JSON(200, article)
	}
}

func GetArticles(dbpool *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c.Request.Context(), time.Second*3)
		defer cancel()

		articles, err := models.GetArticles(ctx, dbpool)
		if err != nil {
			Handle505Error(c, err)
			return
		}
		c.JSON(http.StatusOK, articles)
	}
}
