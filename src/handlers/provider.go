package handlers

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"storage"
)

type Provider struct {
	Storage     storage.Storage
	MasterToken string
	instance    *echo.Echo
}

func (p *Provider) init() {
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("/ALL", p.showAll)
	e.GET("/:key", p.getHandler)
	e.POST("/:key", p.setHandler, p.checkMasterToken)
	e.DELETE("/:key", p.deleteHandler, p.checkMasterToken)
	p.instance = e
}

func (p *Provider) Serve() error {
	p.init()
	return p.instance.Start(":8888")
}
