// internal/repository/covid_repository.go
package repository

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type CovidRepository interface {
	FetchCovidData() ([]Case, error)
}

type covidRepository struct{}

func NewCovidRepository() CovidRepository {
	return &covidRepository{}
}

func (r *covidRepository) FetchCovidData() ([]Case, error) {
	url := "https://static.wongnai.com/devinterview/covid-cases.json"
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var res ResponseCovidTrack
	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, err
	}
	return res.Data, nil
}
