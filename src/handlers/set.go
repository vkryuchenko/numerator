package handlers

import (
	"github.com/labstack/echo"
	"strconv"
)

func (p *Provider) setHandler(ctx echo.Context) error {
	r := ctx.Request()
	err := r.ParseForm()
	if err != nil {
		return err
	}
	value := r.PostFormValue("value")
	v, err := strconv.Atoi(value)
	if err != nil {
		return err
	}
	return p.Storage.Set(ctx.Param("key"), uint(v))
}
