package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) Root(c echo.Context) error {

	return c.Render(http.StatusOK, "hallo.html", map[string]interface{}{
		"name": "Test!",
	})

}
