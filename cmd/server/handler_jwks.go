package server

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type jwksHandler struct {
	jwksBytes []byte
}

func (h *jwksHandler) GetJSONWebKeySet(c echo.Context) error {
	c.Response().Header().Set(echo.HeaderContentType, "application/jwks+json")
	return c.JSONBlob(http.StatusOK, h.jwksBytes)
}
