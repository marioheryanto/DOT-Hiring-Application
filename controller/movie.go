package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/marioheryanto/DOT-Hiring-Application/library"
	"github.com/marioheryanto/DOT-Hiring-Application/model"
)

func GetALLMovieHandler(ctx *gin.Context) {
	name := ctx.Query("name")

	data, err := library.GetALLMovieLibrary(&name)
	if err != nil {
		errService, _ := err.(*model.ErrorService)
		ctx.JSON(errService.Code, model.Response{Message: errService.Error()})
		return
	}

	ctx.JSON(http.StatusOK, model.Response{Data: data})
}

func CreateMovieHandler(ctx *gin.Context) {
	movie := model.Movie{}
	ctx.BindJSON(&movie)

	err := library.CreateMovieLibrary(&movie)
	if err != nil {
		errService, _ := err.(*model.ErrorService)
		ctx.JSON(errService.Code, model.Response{Message: errService.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, model.Response{Message: "Movie Created !"})
}

func ReplaceMovieHandler(ctx *gin.Context) {
	movie := model.Movie{}
	ctx.BindJSON(&movie)

	name := ctx.Query("name")
	if name == "" {
		ctx.JSON(http.StatusBadRequest, model.Response{Message: "Missing Movie Name !"})
		return
	}

	err := library.ReplaceMovieLibrary(&movie, name)
	if err != nil {
		errService, _ := err.(*model.ErrorService)
		ctx.JSON(errService.Code, model.Response{Message: errService.Error()})
		return
	}

	ctx.JSON(http.StatusOK, model.Response{Message: "Success Replace the Movie !"})
}

func EditMovieHandler(ctx *gin.Context) {
	movie := model.Movie{}
	ctx.BindJSON(&movie)

	name := ctx.Query("name")
	if name == "" {
		ctx.JSON(http.StatusBadRequest, model.Response{Message: "Missing Movie Name !"})
		return
	}

	err := library.EditMovieLibrary(&movie, name)
	if err != nil {
		errService, _ := err.(*model.ErrorService)
		ctx.JSON(errService.Code, model.Response{Message: errService.Error()})
		return
	}

	ctx.JSON(http.StatusOK, model.Response{Message: "Success Edit the Movie !"})
}

func DeleteMovieHandler(ctx *gin.Context) {
	name := ctx.Query("name")
	if name == "" {
		ctx.JSON(http.StatusBadRequest, model.Response{Message: "Missing Movie Name !"})
		return
	}

	err := library.DeleteMovieLibrary(name)
	if err != nil {
		errService, _ := err.(*model.ErrorService)
		ctx.JSON(errService.Code, model.Response{Message: errService.Error()})
		return
	}

	ctx.JSON(http.StatusOK, model.Response{Message: "Success Delete the Movie !"})
}
