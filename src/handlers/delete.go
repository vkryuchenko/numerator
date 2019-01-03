package handlers

import (
	"github.com/labstack/echo"
	"net/http"
)

func (p *Provider) deleteHandler(ctx echo.Context) error {
	key := ctx.Param("key")
	err := p.Storage.Delete(key)
	if err != nil {
		return err
	}
	return ctx.String(http.StatusOK, key+" dropped")
}
