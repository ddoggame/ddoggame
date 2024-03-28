package repository

import "github.com/stretchr/testify/mock"

type NewCovidRepositoryMock struct {
	mock.Mock
}

func NewNewCovidRepositoryMock() *NewCovidRepositoryMock {
	return &NewCovidRepositoryMock{}
}

func (m *NewCovidRepositoryMock) FetchCovidData() ([]Case, error) {
	args := m.Called()
	return args.Get(0).([]Case), args.Error(1)
}
