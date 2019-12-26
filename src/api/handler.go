package api

import (
	"net/http"
)

func Handle(c echo.Content) {

}

func Hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
