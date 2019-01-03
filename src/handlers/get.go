package handlers

import (
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

func (p *Provider) getHandler(ctx echo.Context) error {
	value, err := p.Storage.Get(ctx.Param("key"))
	if err != nil {
		return err
	}
	return ctx.String(http.StatusOK, strconv.Itoa(int(value)))
}
