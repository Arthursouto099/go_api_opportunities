package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func PutOpeningHandler(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, gin.H{
		"message": "PUT opening",
	})
}
