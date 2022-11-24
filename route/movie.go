package route

import (
	"github.com/gin-gonic/gin"
	"github.com/marioheryanto/DOT-Hiring-Application/controller"
)

func MovieRoutes(router *gin.Engine) {
	movieGroup := router.Group("/movie")
	movieGroup.GET("/", controller.GetALLMovieHandler)
	movieGroup.POST("/create", controller.CreateMovieHandler)
	movieGroup.PUT("/replace", controller.ReplaceMovieHandler)
	movieGroup.PATCH("/edit", controller.EditMovieHandler)
	movieGroup.DELETE("/delete", controller.DeleteMovieHandler)
}
