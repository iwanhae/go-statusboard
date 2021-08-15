package server

import (
	"github.com/iwanhae/go-statusboard/pkg/monitor"
	"github.com/labstack/echo"
)

func NewServer(looker *monitor.Looker) *echo.Echo {
	e := echo.New()
	e.Server.SetKeepAlivesEnabled(true)
	e.Use(RequestIDGenerator)
	e.Use(LoggerMiddleware)

	e.File("/", "front/out/index.html")
	e.Static("/", "front/out")
	e.GET("/meta", GetMetaHandler(looker))
	e.GET("/stream", GetStreamHandler(looker))
	return e
}
