package main

import (
	"net/http"
	"github.com/havyx/Golang-Its-time-to-go/model"
	"github.com/labstack/echo/v4"
)

type WebServer struct {
	Products *model.Products
}

func NewWebServer() *WebServer {
	return &WebServer{}
}

func (w WebServer) Serve() {
	e := echo.new()
	e.GET("/product", w.getAll)
	e.Logger.Fatal(e.start(":8585"))
}

func (w WebServer) getAll(c echo.Context) error {
	return c.JSON(http.StatusOK, w.Products)
}
