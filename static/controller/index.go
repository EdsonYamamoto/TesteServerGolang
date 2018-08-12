package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GETIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "teste1.tmpl.html", nil)
}
