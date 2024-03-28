package service

import "github.com/stretchr/testify/mock"

type covidServiceMock struct {
	mock.Mock
}

func NewCovidServiceMock() *covidServiceMock {
	return &covidServiceMock{}
}

func (m *covidServiceMock) GetSummary() (CovidSummaryResponse, error) {
	args := m.Called()
	return args.Get(0).(CovidSummaryResponse), args.Error(1)
}
