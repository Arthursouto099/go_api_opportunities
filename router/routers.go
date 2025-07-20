package router

import (
	"github.com/Arthursouto099/go_api_opportunities/handler"
	"github.com/gin-gonic/gin"
)

// funções e variaveis são visiveis dentro do mesmo package -> não necessita de imports

// Receive pointer of gin.Engine
func initializeRoutes(router *gin.Engine) {

	v1 := router.Group("/api/v1")
	// delimiting the scope
	{
		v1.GET("/opening", handler.ShowOpeningHandler)
		v1.POST("/opening", handler.CreateOpeningHandler)
		v1.DELETE("/opening", handler.DeleteOpeningHandler)
		v1.PUT("/opening", handler.PutOpeningHandler)
		v1.GET("/openings", handler.ShowAllOpeningHandler)
	}

}
