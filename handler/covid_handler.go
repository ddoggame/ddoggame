// internal/handler/covid_handler.go
package handler

import (
	"covid-lmwn/service"

	"github.com/gin-gonic/gin"
)

type CovidHandler interface {
	GetSummary(c *gin.Context)
}

type covidHandler struct {
	covidService service.CovidService
}

func NewCovidHandler(covidService service.CovidService) CovidHandler {
	return &covidHandler{
		covidService: covidService,
	}
}

func (h *covidHandler) GetSummary(c *gin.Context) {
	summary, err := h.covidService.GetSummary()
	if err != nil {
		handleError(c, err)
		return
	}
	handleResponse(c, AppResponse{Code: 200, Data: summary, Message: "get data success !"})

}
