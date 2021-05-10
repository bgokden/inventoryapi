package server

import (
	"github.com/deepmap/oapi-codegen/pkg/middleware"
	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
	"github.com/pkg/errors"

	"github.com/bgokden/inventoryapi/api"
)

func CreateEchoServer() (*echo.Echo, error) {

	swagger, err := api.GetSwagger()
	if err != nil {
		return nil, errors.Errorf("Error loading swagger spec\n: %s", err)
	}

	// Clear out the servers array in the swagger spec, that skips validating
	// that server names match. We don't know how this thing will be run.
	// swagger.Servers = nil

	// Create an instance of our handler which satisfies the generated interface
	serverImpl := NewInventoryAPI()

	// This is how you set up a basic Echo router
	e := echo.New()
	// Log all requests
	e.Use(echomiddleware.Logger())
	// Use our validation middleware to check all requests against the
	// OpenAPI schema.
	e.Use(middleware.OapiRequestValidator(swagger))

	// We now register our petStore above as the handler for the interface
	// api.RegisterHandlers(e, serverImpl)
	api.RegisterHandlersWithBaseURL(e, serverImpl, "/v0")

	return e, nil
}
