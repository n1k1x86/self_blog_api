package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ErrNoTag struct {
	TagID int64
}

func (e ErrNoTag) Error() string {
	return fmt.Sprintf(`tag %d is not found`, e.TagID)
}

func Handle400Error(c *gin.Context) {
	c.AbortWithStatus(http.StatusBadRequest)
}

func Handle404Error(c *gin.Context, err error) {
	c.AbortWithError(http.StatusNotFound, err)
}

func Handle505Error(c *gin.Context, err error) {
	c.AbortWithError(http.StatusInternalServerError, err)
}
