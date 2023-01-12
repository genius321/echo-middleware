package app

import (
	"github.com/genius321/echo-middleware/internal/app/endpoint"
	"github.com/genius321/echo-middleware/internal/app/mw"
	"github.com/genius321/echo-middleware/internal/app/service"
	"github.com/labstack/echo/v4"
)

type App struct {
	e    *endpoint.Endpoint
	s    *service.Service
	echo *echo.Echo
}

func New() *App {
	a := &App{}

	a.s = service.New()

	a.e = endpoint.New(a.s)

	a.echo = echo.New()

	a.echo.GET("/status", a.e.Status, mw.RoleCheck)

	return a
}

func (a *App) Run() error {
	return a.echo.Start(":8080")
}
