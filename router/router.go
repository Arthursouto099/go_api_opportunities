package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// exports e imports -> convenção de nomeclatura
// A -> visivel a todos os outros pacotes, a -> não
func Initialize() {

	// inicializando o router
	router := gin.Default()

	//definindo a rota de ping
	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.Run("localhost:8080")
}
