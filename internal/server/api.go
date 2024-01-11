package server

import (
	"os"

	"github.com/labstack/echo/v4"
)

func (s ServerEcho) Start() {
	e := echo.New()

	s.Routes(e)

	port := os.Getenv("PORT")
	e.Logger.Fatal(e.Start(":" + port))
}

func (s ServerEcho) Routes(e *echo.Echo) {

}
