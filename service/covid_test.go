package service_test

import (
	"covid-lmwn/errs"
	"covid-lmwn/repository"
	"covid-lmwn/service"
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAgeGroup(t *testing.T) {

	type testCase struct {
		name     string
		age      int
		expected string
	}

	cases := []testCase{
		{
			name:     "age 25",
			age:      25,
			expected: service.AgeGroup0to30,
		},
		{
			name:     "age 40",
			age:      40,
			expected: service.AgeGroup31to60,
		},
		{
			name:     "age 70",
			age:      70,
			expected: service.AgeGroup60plus,
		},
		{
			name:     "age 100",
			age:      100,
			expected: service.AgeGroup60plus,
		},
		{
			name:     "age -10",
			age:      -10,
			expected: service.AgeGroupNA,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			result := service.GetAgeGroup(c.age)
			assert.Equal(t, c.expected, result)
		})
	}

}

// go test covid-lmwn/service -cover
// go test covid-lmwn/service -bench=. -benchmem
func BenchmarkGetAgeGroup(b *testing.B) {
	for i := 0; i < b.N; i++ {
		service.GetAgeGroup(25)
	}
}

func TestGetSummaryService(t *testing.T) {

	type testCase struct {
		name          string
		ProvinceCount map[string]int
		AgeGroupCount map[string]int
		Arg           []repository.Case
	}
	cases := []testCase{
		{
			name:          "case 1",
			ProvinceCount: map[string]int{"Bangkok": 2, "Rayong": 3},
			AgeGroupCount: map[string]int{"0-30": 1, "31-60": 2, "60+": 1, "N/A": 1},
			Arg: []repository.Case{
				{Age: 25, Province: "Bangkok"},
				{Age: 35, Province: "Bangkok"},
				{Age: 45, Province: "Rayong"},
				{Age: 70, Province: "Rayong"},
				{Age: -10, Province: "Rayong"},
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			mockRepo := repository.NewNewCovidRepositoryMock()
			mockRepo.On("FetchCovidData").Return(c.Arg, nil)
			service := service.NewCovidService(mockRepo)
			result, err := service.GetSummary()
			assert.Nil(t, err)
			assert.Equal(t, c.ProvinceCount, result.Province)
			assert.Equal(t, c.AgeGroupCount, result.AgeGroup)
		})

	}
	t.Run("c.name", func(t *testing.T) {
		mockRepo := repository.NewNewCovidRepositoryMock()
		mockRepo.On("FetchCovidData").Return([]repository.Case{}, errors.New("test error"))
		service := service.NewCovidService(mockRepo)
		_, err := service.GetSummary()
		assert.Error(t, err)
	})

	t.Run("case error repo", func(t *testing.T) {
		mockRepo := repository.NewNewCovidRepositoryMock()
		mockRepo.On("FetchCovidData").Return([]repository.Case{}, errs.NewUnexpectedError())
		service := service.NewCovidService(mockRepo)
		_, err := service.GetSummary()
		assert.ErrorIs(t, err, errs.NewUnexpectedError())
	})

	t.Run("case error repo", func(t *testing.T) {
		mockRepo := repository.NewNewCovidRepositoryMock()
		mockRepo.On("FetchCovidData").Return([]repository.Case{{Age: 25, Province: "new york"}}, nil)
		service := service.NewCovidService(mockRepo)
		_, err := service.GetSummary()
		// assert.ErrorIs(t, err, errors.New("new york not thai"))
		fmt.Println("err", err)
		fmt.Println("expected", "new york not thai")
	})

}
