package server_test

import (
	"net/http"
	"testing"

	"github.com/deepmap/oapi-codegen/pkg/testutil"
	"github.com/stretchr/testify/assert"

	"github.com/bgokden/inventoryapi/api"
	"github.com/bgokden/inventoryapi/server"
)

func TestInventoryAPI(t *testing.T) {
	var err error
	// Here, we Initialize echo
	e, err := server.CreateEchoServer()
	assert.Nil(t, err)

	newInventory := api.Inventory{
		Inventory: &[]api.Stock{
			{
				ArtId: "1",
				Name:  "test1",
				Stock: "5",
			},
			{
				ArtId: "2",
				Name:  "test2",
				Stock: "3",
			},
		},
	}
	result := testutil.NewRequest().Post("/v0/inventory").WithJsonBody(newInventory).Go(t, e)
	// We expect 200 code on successful pet insertion
	assert.Equal(t, http.StatusOK, result.Code())
}
