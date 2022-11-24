package library

import (
	"github.com/imdario/mergo"
	"github.com/marioheryanto/DOT-Hiring-Application/config"
	"github.com/marioheryanto/DOT-Hiring-Application/model"
	"github.com/marioheryanto/DOT-Hiring-Application/repository"
)

func GetALLMovieLibrary(name *string) ([]model.Movie, error) {
	if *name != "" {
		movie, err := repository.GetMovieRepository(*name)
		if err != nil {
			return nil, err
		}

		return []model.Movie{movie}, nil
	}

	return repository.GetALLMovieRepository()
}

func GetMovieLibrary(name string) (model.Movie, error) {

	return repository.GetMovieRepository(name)
}

func CreateMovieLibrary(request *model.Movie) error {
	err := repository.CreateMovieRepository(request)
	if err != nil {
		return err
	}

	repository.DeleteRedis("movies")

	return nil
}

func ReplaceMovieLibrary(request *model.Movie, name string) error {
	movieExisting, err := repository.GetMovieRepository(name)
	if err != nil {
		return err
	}

	request.Id = movieExisting.Id

	// begin trx
	trx := config.ConnectDatabase().Begin()

	err = repository.TRXUpdateMovieActorRepository(trx, request)
	if err != nil {
		trx.Rollback()
		return err
	}

	err = repository.TRXReplaceMovieRepository(trx, request)
	if err != nil {
		trx.Rollback()
		return err
	}

	trx.Commit()

	repository.DeleteRedis("movies", name)

	return nil
}

func EditMovieLibrary(request *model.Movie, name string) error {

	movieExisting, err := repository.GetMovieRepository(name)
	if err != nil {
		return err
	}

	err = mergo.Merge(&movieExisting, *request, mergo.WithOverride)
	if err != nil {
		return err
	}

	// err = repository.ReplaceMovieRepository(&movieExisting)
	// if err != nil {
	// 	return err
	// }

	// begin trx
	trx := config.ConnectDatabase().Begin()

	err = repository.TRXUpdateMovieActorRepository(trx, &movieExisting)
	if err != nil {
		trx.Rollback()
		return err
	}

	err = repository.TRXReplaceMovieRepository(trx, &movieExisting)
	if err != nil {
		trx.Rollback()
		return err
	}

	trx.Commit()

	repository.DeleteRedis("movies", name)

	return nil
}

func DeleteMovieLibrary(name string) error {
	err := repository.DeleteMovieRepository(name)
	if err != nil {
		return err
	}

	repository.DeleteRedis("movies", name)

	return nil
}
