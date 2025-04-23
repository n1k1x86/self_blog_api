package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Handle400Error(c *gin.Context) {
	c.AbortWithStatus(http.StatusBadRequest)
}

func Handle404Error(c *gin.Context, err error) {
	c.AbortWithError(http.StatusNotFound, err)
}

func Handle505Error(c *gin.Context, err error) {
	c.AbortWithError(http.StatusInternalServerError, err)
}
