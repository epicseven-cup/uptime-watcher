package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	app := echo.New()
	// Reouting for path '/'
	app.GET("/", func(c echo.Context) error {
		// Return the status 200, taking constant from the http package
		// Return content of "Test"
		return c.String(http.StatusOK, "Test")
	})
	// Start logger on fatal on local port 1323
	app.Logger.Fatal(app.Start(":1323"))
}
