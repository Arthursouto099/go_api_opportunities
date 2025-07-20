package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShowOpeningHandler(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, gin.H{
		"message": "GET opening",
	})
}
