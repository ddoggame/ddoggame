package handler_test

import (
	"covid-lmwn/handler"
	"covid-lmwn/service"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// go test covid-lmwn/service -cover

func TestGetSummary(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		expected := service.CovidSummaryResponse{
			Province: map[string]int{"Bangkok": 100},
			AgeGroup: map[string]int{"0-30": 100, "31-60": 100, "60+": 100, "N/A": 100},
		}

		mockCovidService := service.NewCovidServiceMock()
		mockCovidService.On("GetSummary").Return(expected, nil)

		handler := handler.NewCovidHandler(mockCovidService)

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
