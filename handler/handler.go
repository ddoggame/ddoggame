package handler

import (
	"covid-lmwn/errs"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AppResponse struct {
	Message string      `json:"message"`
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
}

func handleError(c *gin.Context, err error) {
	switch e := err.(type) {
	case errs.AppError:
		c.JSON(e.Code, e)
		return
	case error:
		c.JSON(http.StatusInternalServerError, "Internal Server Error")
		return

	}
}

func handleResponse(c *gin.Context, d AppResponse) {
	c.JSON(d.Code, d)
}
