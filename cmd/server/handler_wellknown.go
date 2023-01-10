package server

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type wellKnownHandler struct {
	providerBytes []byte
}

func (h *wellKnownHandler) GetConfiguration(c echo.Context) error {
	return c.JSONBlob(http.StatusOK, h.providerBytes)
}
