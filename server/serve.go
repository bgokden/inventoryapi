package server

import (
	"github.com/deepmap/oapi-codegen/pkg/middleware"
	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
	"github.com/pkg/errors"

	"github.com/bgokden/inventoryapi/api"
)

// CreateEchoServer creates an echo server with inventory api implementation
func CreateEchoServer() (*echo.Echo, error) {

	swagger, err := api.GetSwagger()
	if err != nil {
		return nil, errors.Errorf("Error loading swagger spec\n: %s", err)
	}

	// Create an instance of our handler which satisfies the generated interface
	serverImpl := NewInventoryAPI()

	// This is how you set up a basic Echo router
	e := echo.New()

	// Added this locally to show queries in the Swagger UI
	e.Use(echomiddleware.CORSWithConfig(echomiddleware.CORSConfig{
		AllowOrigins: []string{"https://editor.swagger.io"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	// Log all requests
	e.Use(echomiddleware.Logger())
	// Use our validation middleware to check all requests against the
	// OpenAPI schema.
	e.Use(middleware.OapiRequestValidator(swagger))

	// We now register our serverImple above as the handler for the interface
	// It is possible to register multiple services and versions
	api.RegisterHandlersWithBaseURL(e, serverImpl, "/v0")

	return e, nil
}
