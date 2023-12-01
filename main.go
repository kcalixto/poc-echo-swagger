package main

import (
	"encoding/json"
	"net/http"
	"time"

	_ "poc-echo-swagger/docs" // Import the docs

	"github.com/labstack/echo/v4"
	"github.com/swaggo/echo-swagger"
)

// @title Echo Swagger Example API
// @version 1.0
// @description This is a sample server server.
// @host localhost:1323
// @BasePath /
// @schemes http
func main() {
	e := echo.New()
	e.GET("/health", HealthCheck)
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.Logger.Fatal(e.Start(":1323"))
}

type RequestBody struct {
	FakeValue  string `json:"fake_value"`
	FakeValue2 string `json:"fake_value2"`
	FakeValue3 string `json:"fake_value3"`
}

type ResponseStruct struct {
	Status string
	Time   string
}

// HealthCheck godoc
// @Accept json
// @Produce json
// @Summary Show the status of server.
// @Description get the status of server.
// @Tags root
// @Accept */*
// @Produce json
// @Param body body RequestBody true "request body"
// @Success 200 {object} ResponseStruct
// @Router /health [post]
func HealthCheck(c echo.Context) error {
	str, _ := json.Marshal(ResponseStruct{
		Status: "OK",
		Time:   time.Now().String(),
	})

	return c.JSON(http.StatusOK, str)
}
