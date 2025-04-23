package handler

import (
	"api/db/models"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/ydb-platform/ydb-go-sdk/v3/log"
)

func AddTag(dbpool *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c.Request.Context(), time.Second*3)
		defer cancel()

		bodyData, err := io.ReadAll(c.Request.Body)
		if err != nil {
			Handle505Error(c, err)
			return
		}

		tag := models.TagREST{}
		err = json.Unmarshal(bodyData, &tag)
		if err != nil {
			Handle505Error(c, err)
			return
		}
		err = models.AddTag(ctx, dbpool, &tag)
		if err != nil {
			Handle505Error(c, err)
			return
		}
		c.JSON(http.StatusOK, tag)
	}
}

func GetTags(dbpool *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {

		ctx, cancel := context.WithTimeout(c.Request.Context(), time.Second*3)
		defer cancel()

		result, err := models.GetTags(ctx, dbpool)
		if err != nil {
			log.Error(err)
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		c.JSON(http.StatusOK, result)
	}
}

func EditTag(dbpool *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		body, err := io.ReadAll(c.Request.Body)
		if err != nil {
			Handle505Error(c, err)
			return
		}
		var tag models.TagREST
		err = json.Unmarshal(body, &tag)
		if err != nil {
			Handle505Error(c, err)
			return
		}

		idQuery, ok := c.Params.Get("id")
		if !ok {
			Handle404Error(c, err)
			return
		}

		id, err := strconv.Atoi(idQuery)
		if err != nil {
			Handle400Error(c)
			return
		}
		tag.ID = int64(id)
		ctx, cancel := context.WithTimeout(c.Request.Context(), time.Second*3)
		defer cancel()

		err = models.EditTag(ctx, dbpool, &tag)
		if err != nil {
			if errors.Is(err, models.ErrNotFound) {
				Handle404Error(c, err)
				return
			}
			Handle505Error(c, err)
			return
		}

		c.JSON(http.StatusOK, tag)
	}

}

func DeleteTag(dbpool *pgxpool.Pool) gin.HandlerFunc {
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
			Handle400Error(c)
			return
		}

		err = models.DeleteTag(ctx, dbpool, int64(id))
		if err != nil {
			if errors.Is(err, models.ErrNotFound) {
				Handle404Error(c, err)
				return
			}
			Handle505Error(c, err)
			return
		}
		c.Status(http.StatusNoContent)
	}
}

func GetTagByID(dbpool *pgxpool.Pool) gin.HandlerFunc {
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
			Handle404Error(c, err)
			return
		}

		tag, err := models.GetTagByID(ctx, dbpool, int64(id))
		if err != nil {
			if errors.Is(err, models.ErrNotFound) {
				Handle404Error(c, err)
				return
			}
			Handle505Error(c, err)
			return
		}
		c.JSON(http.StatusOK, tag)
	}
}
