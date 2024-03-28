package service

import (
	"covid-lmwn/errs"
	"covid-lmwn/logs"
	"covid-lmwn/repository"
	"errors"
)

type CovidService interface {
	GetSummary() (CovidSummaryResponse, error)
}

type covidService struct {
	covidRepo repository.CovidRepository
}

func NewCovidService(covidRepo repository.CovidRepository) CovidService {
	return &covidService{
		covidRepo: covidRepo,
	}
}

func (s *covidService) GetSummary() (CovidSummaryResponse, error) {

	data, err := s.covidRepo.FetchCovidData()
	if err != nil {
		logs.Error(err)
		return CovidSummaryResponse{}, errs.NewUnexpectedError()
	}
	ageGroups := map[string]int{
		AgeGroup0to30:  0,
		AgeGroup31to60: 0,
		AgeGroup60plus: 0,
		AgeGroupNA:     0,
	}

	provinceCases := make(map[string]int)
	for _, c := range data {
		group := GetAgeGroup(c.Age)
		ageGroups[group]++
		if c.Province != "" {
			provinceCases[c.Province]++
			if c.Province == "new york" {
				return CovidSummaryResponse{}, errors.New("new york not thai")
			}
		}
	}
	response := CovidSummaryResponse{Province: provinceCases, AgeGroup: ageGroups}

	return response, nil
}
