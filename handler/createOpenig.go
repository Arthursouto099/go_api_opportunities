package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateOpeningHandler(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, gin.H{
		"message": "POST opening",
	})
}
