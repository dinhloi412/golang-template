package api

import (
	"net/http"
	modelv1 "project/model/v1"

	"github.com/labstack/echo/v4"
)

type Api struct {
	e *echo.Echo
}

func NewApi(svc modelv1.Controller) *Api {
	e := echo.New()

	// Get user
	e.GET("/api/v1/user/:id", apiGetUser(svc))

	return &Api{
		e: e,
	}
}

func (c *Api) Run(server *http.Server) error {

	c.e.HideBanner = true
	c.e.Debug = true
	c.e.Logger.SetOutput(NewNullWriter())
	return c.e.StartServer(server)
}
