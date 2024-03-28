// go build integration
package handler_test

import (
	"covid-lmwn/handler"
	"covid-lmwn/repository"
	"covid-lmwn/service"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetSummaryIntegrationService(t *testing.T) {
	t.Run("success", func(t *testing.T) {

		expected := service.CovidSummaryResponse{
			Province: map[string]int{"Bangkok": 2, "Rayong": 3},
			AgeGroup: map[string]int{"0-30": 1, "31-60": 2, "60+": 1, "N/A": 1},
		}

		mockRepo := repository.NewNewCovidRepositoryMock()
		mockRepo.On("FetchCovidData").Return([]repository.Case{
			{Age: 25, Province: "Bangkok"},
			{Age: 35, Province: "Bangkok"},
			{Age: 45, Province: "Rayong"},
			{Age: 70, Province: "Rayong"},
			{Age: -10, Province: "Rayong"},
		}, nil)

		serviceCovid := service.NewCovidService(mockRepo)
		handler := handler.NewCovidHandler(serviceCovid)

		gin.SetMode(gin.TestMode)
		app := gin.Default()
		app.GET("/covid/summary", handler.GetSummary)

		req := httptest.NewRequest("GET", "/covid/summary", nil)
		w := httptest.NewRecorder()
		app.ServeHTTP(w, req)
		assert.Equal(t, 200, w.Code)

		var responseBody struct {
			Code    int                          `json:"code"`
			Data    service.CovidSummaryResponse `json:"data"`
			Message string                       `json:"message"`
		}
		if assert.Equal(t, 200, w.Code) {
			err := json.Unmarshal(w.Body.Bytes(), &responseBody)
			if err != nil {
				t.Fatal(err)
			}

			assert.Equal(t, expected, responseBody.Data)
			assert.Equal(t, "get data success !", responseBody.Message)
		}
	})

}
