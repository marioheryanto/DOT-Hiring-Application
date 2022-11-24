package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"testing"

	"github.com/marioheryanto/DOT-Hiring-Application/model"
	"github.com/stretchr/testify/assert"
)

func TestCreateMovieHandler(t *testing.T) {
	c := http.Client{}

	payload := model.Movie{
		Name:   "black panther",
		Genre:  "action",
		Year:   "2017",
		Actors: []model.Actor{{Name: "Budi", Gender: "male"}},
	}

	payloadByte, _ := json.Marshal(payload)

	r, err := c.Post("http://localhost:8080/movie/create", "application/json", bytes.NewBuffer(payloadByte))
	if err != nil {
		log.Printf("error when call api, err => %v", err.Error())
	}

	assert.Equal(t, http.StatusCreated, r.StatusCode, "can't create movie")
}

func TestGetALLMovieHandler(t *testing.T) {
	c := http.Client{}

	r, err := c.Get("http://localhost:8080/movie")
	if err != nil {
		log.Printf("error when get api, err => %v", err.Error())
	}

	resp, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("error when read body, err => %v", err.Error())
	}

	defer r.Body.Close()

	responseData := model.Response{}
	err = json.Unmarshal(resp, &responseData)
	if err != nil {
		log.Printf("error when unmarshalling, err => %v", err.Error())
	}

	log.Printf("response data => %+v", responseData)

	assert.Equal(t, true, len(responseData.Data) > 0, "can't get movies")
}

func TestReplaceMovieHandler(t *testing.T) {
	c := &http.Client{}

	payload := model.Movie{
		Name:   "black panther 1",
		Genre:  "action",
		Year:   "2017",
		Actors: []model.Actor{{Name: "Budi", Gender: "male"}, {Name: "Bambang", Gender: "male"}},
	}

	payloadByte, _ := json.Marshal(payload)

	url := fmt.Sprintf("http://localhost:8080/movie/replace?name=%s", url.QueryEscape("black panther"))

	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(payloadByte))
	if err != nil {
		log.Printf("error when create new request, err => %v", err.Error())
	}

	r, err := c.Do(req)
	if err != nil {
		log.Printf("error when call api, err => %v", err.Error())
	}

	assert.Equal(t, http.StatusOK, r.StatusCode, "can't replace movie")
}

func TestEditMovieHandler(t *testing.T) {
	c := &http.Client{}

	payload := model.Movie{
		Name:   "black panther one",
		Actors: []model.Actor{{Name: "Budi", Gender: "male"}, {Name: "Bambang", Gender: "male"}, {Name: "Rudi", Gender: "male"}},
	}

	payloadByte, _ := json.Marshal(payload)

	url := fmt.Sprintf("http://localhost:8080/movie/edit?name=%s", url.QueryEscape("black panther 1"))

	req, err := http.NewRequest("PATCH", url, bytes.NewBuffer(payloadByte))
	if err != nil {
		log.Printf("error when create new request, err => %v", err.Error())
	}

	r, err := c.Do(req)
	if err != nil {
		log.Printf("error when call api, err => %v", err.Error())
	}

	assert.Equal(t, http.StatusOK, r.StatusCode, "can't edit movie")
}

func TestDeleteMovieHandler(t *testing.T) {
	c := &http.Client{}

	url := fmt.Sprintf("http://localhost:8080/movie/delete?name=%s", url.QueryEscape("black panther one"))

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		log.Printf("error when create new request, err => %v", err.Error())
	}

	r, err := c.Do(req)
	if err != nil {
		log.Printf("error when call api, err => %v", err.Error())
	}

	assert.Equal(t, http.StatusOK, r.StatusCode, "can't delete movie")
}
