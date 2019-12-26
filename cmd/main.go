package main

// This is the main entrance of Gitlab-Dingbot
import (
	"github.com/labstack/echo/v4/middleware"

	"github.com/ray1888/self-defined-dingbot/src/api"
	"github.com/ray1888/self-defined-dingbot/src/logger"
)


func main() {

	logger.Inst().Info("get logger successful")
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", api.Hello)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
