package handlers

import (
	"github.com/labstack/echo"
	"net/http"
)

func (p *Provider) checkMasterToken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		authToken := ctx.Request().Header.Get("Authorization")
		if authToken != p.MasterToken {
			return ctx.String(http.StatusUnauthorized, "MASTER_TOKEN required")
		}
		return next(ctx)
	}
}
