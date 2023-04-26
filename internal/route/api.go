package route

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sparkymat/finchphone/internal/auth"
)

func registerAPIRoutes(app *echo.Group, cfg ConfigService, db DatabaseService) {
	apiGroup := app.Group("api")

	if cfg.ReverseProxyAuthentication() {
		apiGroup.Use(auth.ProxyAuthMiddleware(cfg, db))
	} else {
		apiGroup.Use(auth.APIMiddleware(cfg, db))
	}

	apiGroup.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		TokenLookup: "header:X-CSRF-Token",
	}))
}
