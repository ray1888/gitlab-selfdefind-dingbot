package api

import (
	echo "github.com/labstack/echo/v4"
	"net/http"
)

//func Handle(c echo.Content) {
//
//}

func Hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
