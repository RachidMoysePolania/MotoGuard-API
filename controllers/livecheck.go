package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HealtCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"mensaje": "pong"})
}
