// main.go
package main

import (
	"covid-lmwn/handler"
	"covid-lmwn/repository"
	"covid-lmwn/service"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {

	covidRepo := repository.NewCovidRepository()
	covidService := service.NewCovidService(covidRepo)
	covidHandler := handler.NewCovidHandler(covidService)

	// Setup Gin server
	r := gin.Default()

	r.GET("/covid/summary", covidHandler.GetSummary)
	err := r.Run(":8087")
	if err != nil {
		fmt.Println(err)
	}

}
