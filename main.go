package main

import (
	"github.com/gin-gonic/gin"
	"github.com/marioheryanto/DOT-Hiring-Application/config"
	"github.com/marioheryanto/DOT-Hiring-Application/model"
	"github.com/marioheryanto/DOT-Hiring-Application/route"
)

func main() {
	router := gin.New()

	db := config.ConnectDatabase()
	db.AutoMigrate(&model.Movie{})
	db.AutoMigrate(&model.Actor{})

	route.MovieRoutes(router)

	router.Run(":8080")
}
