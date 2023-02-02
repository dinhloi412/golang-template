package api

import (
	"net/http"

	modelv1 "project/model/v1"

	"github.com/labstack/echo/v4"
)

func apiGetUser(svc modelv1.Controller) echo.HandlerFunc {
	return echo.HandlerFunc(func(c echo.Context) error {
		var user modelv1.User

		err := c.Bind(&user)
		if err != nil {
			return c.String(http.StatusBadRequest, "bad request")
		}

		user, err := svc.SaveUser(&user)
		if err != nil {
			c.JSON(http.StatusOK, JsonError(http.StatusInternalServerError, err.Error()))
		}

		return c.JSON(http.StatusOK, JsonSuccess("success", user))
	})
}
