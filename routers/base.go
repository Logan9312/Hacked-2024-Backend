package routers

import (
	"net/http"
	"os"

	"github.com/Logan9312/Hacked-2024-Backend/src"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func HealthCheck() {

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Enable CORS for all origins
	e.Use(middleware.CORS())

	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, struct{ Status string }{Status: "OK"})
	})

	e.GET("/products", src.FetchPayments)

	httpPort := os.Getenv("PORT")
	if httpPort == "" {
		httpPort = "8080"
	}

	e.Logger.Fatal(e.Start(":" + httpPort))
}
