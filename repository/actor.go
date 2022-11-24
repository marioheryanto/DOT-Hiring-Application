package repository

import (
	"net/http"

	"github.com/marioheryanto/DOT-Hiring-Application/config"
	"github.com/marioheryanto/DOT-Hiring-Application/model"
)

func AddActorToMovie(movie *model.Movie, actors ...model.Actor) error {
	db := config.ConnectDatabase()
	err := db.Model(movie).Association("Actors").Append(actors)
	if err != nil {
		return model.ErrorService{Code: http.StatusInternalServerError, Message: err.Error()}
	}

	return nil
}

func CreateActor(actors ...model.Actor) error {
	db := config.ConnectDatabase()

	for _, actor := range actors {
		err := db.Create(&actor).Error
		if err != nil {
			return model.ErrorService{Code: http.StatusInternalServerError, Message: err.Error()}
		}
	}

	return nil
}
