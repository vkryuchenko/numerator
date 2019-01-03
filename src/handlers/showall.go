package handlers

import (
	"github.com/labstack/echo"
	"net/http"
)

func (p *Provider) showAll(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, p.Storage.ShowData())
}
