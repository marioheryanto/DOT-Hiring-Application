package library

import (
	"github.com/marioheryanto/DOT-Hiring-Application/config"
	"github.com/marioheryanto/DOT-Hiring-Application/model"
	"github.com/marioheryanto/DOT-Hiring-Application/repository"
)

func AddActorToMovie(movieName string, actors ...model.Actor) error {

	movie, _ := repository.GetMovieRepository(movieName)

	db := config.ConnectDatabase()
	err := db.Model(&movie).Association("Actors").Error
	if err != nil {
		return err
	}

	if len(actors) > 0 {
		repository.CreateActor(actors...)
	}

	return repository.AddActorToMovie(&movie, actors...)
}
