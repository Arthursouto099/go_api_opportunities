package router

import (
	"github.com/gin-gonic/gin"
)

// exports e imports -> convenção de nomeclatura
// A -> visivel a todos os outros pacotes, a -> não
func Initialize() {

	// inicializando o router
	// initialize router
	router := gin.Default()

	// initialize routers
	initializeRoutes(router)

	//definindo a rota de ping
	// router.GET("/ping", func(ctx *gin.Context) {
	// 	ctx.JSON(http.StatusOK, gin.H{
	// 		"message": "pong",
	// 	})
	// })

	// run the server
	router.Run("localhost:8080")
}
