package repository

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/marioheryanto/DOT-Hiring-Application/config"
	"github.com/marioheryanto/DOT-Hiring-Application/model"
	"gorm.io/gorm"
)

func GetALLMovieRepository() ([]model.Movie, error) {
	movies := []model.Movie{}

	// get data from redis
	redisCli := config.ConnectRedis()
	data, err := redisCli.Get(context.Background(), "movies").Result()
	if err == nil {
		err = json.Unmarshal([]byte(data), &movies)
		if err == nil {
			return movies, nil
		}
	}

	// get data from DB
	db := config.ConnectDatabase()
	err = db.Model(&model.Movie{}).Preload("Actors").Find(&movies).Error
	if err != nil {
		return movies, model.ErrorService{Code: http.StatusInternalServerError, Message: err.Error()}
	}

	if len(movies) == 0 {
		return movies, model.ErrorService{Code: http.StatusNotFound, Message: "data not found"}
	}

	// set to redis, we dont mind the status
	dataBytes, _ := json.Marshal(movies)
	redisCli.Set(context.Background(), "movies", dataBytes, 0)

	return movies, nil
}

func GetMovieRepository(name string) (model.Movie, error) {
	movie := model.Movie{}

	// get data from redis first
	redisCli := config.ConnectRedis()
	data, err := redisCli.Get(context.Background(), name).Result()
	if err == nil {
		err = json.Unmarshal([]byte(data), &movie)
		if err == nil {
			return movie, nil
		}
	}

	// get data from DB
	db := config.ConnectDatabase()
	err = db.Where("name = ?", name).Model(&model.Movie{}).Preload("Actors").Find(&movie).Error
	if err != nil {
		return movie, model.ErrorService{Code: http.StatusInternalServerError, Message: err.Error()}
	}

	if movie.Name == "" {
		return movie, model.ErrorService{Code: http.StatusNotFound, Message: "data not found"}
	}

	// set to redis, we dont mind the status
	dataBytes, _ := json.Marshal(movie)
	redisCli.Set(context.Background(), name, dataBytes, 0)

	return movie, nil
}

func CreateMovieRepository(request *model.Movie) error {
	db := config.ConnectDatabase()
	err := db.Create(request).Error
	if err != nil {
		return model.ErrorService{Code: http.StatusInternalServerError, Message: err.Error()}
	}

	return nil
}

func DeleteMovieRepository(name string) error {
	var movie model.Movie

	db := config.ConnectDatabase()
	err := db.Where("name = ?", name).Delete(&movie).Error
	if err != nil {
		return model.ErrorService{Code: http.StatusInternalServerError, Message: err.Error()}
	}

	return nil
}

func TRXUpdateMovieActorRepository(trx *gorm.DB, request *model.Movie) error {
	err := trx.Model(request).Association("Actors").Replace(request.Actors)
	if err != nil {
		return model.ErrorService{Code: http.StatusInternalServerError, Message: err.Error()}
	}

	return nil
}

func TRXReplaceMovieRepository(trx *gorm.DB, request *model.Movie) error {
	err := trx.Save(request).Error
	if err != nil {
		return model.ErrorService{Code: http.StatusInternalServerError, Message: err.Error()}
	}

	return nil
}
