package server

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/bgokden/inventoryapi/api"
	"github.com/bgokden/inventoryapi/store"
)

/*
	// Inventory API implements these methods
	// Get current inventory
	// (GET /inventory)
	GetInventory(ctx echo.Context) error
	// Inserts or Updates stocks in Inventory
	// (POST /inventory)
	UpsertInventory(ctx echo.Context) error
	// Lists products
	// (GET /products)
	ListProducts(ctx echo.Context) error
	// Insert or Update products
	// (POST /products)
	UpsertProducts(ctx echo.Context) error
	// Lists products with stock
	// (GET /productstock)
	ListProductStocks(ctx echo.Context) error
	// Sell specified products and update Inventory
	// (POST /sell)
	SellFromInventory(ctx echo.Context) error
*/
type InventoryAPI struct {
	Store store.Store
}

func NewInventoryAPI() *InventoryAPI {
	return &InventoryAPI{
		Store: store.NewInMemoryStore(),
	}
}

// GetInventory converts echo context to params.
func (ia *InventoryAPI) GetInventory(ctx echo.Context) error {
	// Invoke the callback with all the unmarshalled arguments
	result, err := ia.Store.ListInventory()
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, result)
}

// UpsertInventory converts echo context to params.
func (ia *InventoryAPI) UpsertInventory(ctx echo.Context) error {
	var err error
	var newInventory api.Inventory
	err = ctx.Bind(&newInventory)
	if err != nil {
		return err
	}
	err = ia.Store.UpsertInventory(&newInventory)
	return err
}

/*
curl -v --header "Content-Type: application/json" \
  --request POST \
  --data '{"inventory":[{"art_id":"1","name":"test1","stock":"5"},{"art_id":"2","name":"test2","stock":"3"}]}' \
  http://localhost:8080/inventory
*/

// ListProducts converts echo context to params.
func (ia *InventoryAPI) ListProducts(ctx echo.Context) error {
	// Invoke the callback with all the unmarshalled arguments
	result, err := ia.Store.ListProducts()
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, result)
}

// UpsertProducts converts echo context to params.
func (ia *InventoryAPI) UpsertProducts(ctx echo.Context) error {
	var err error
	var newProducts api.Products
	err = ctx.Bind(&newProducts)
	if err != nil {
		return err
	}
	err = ia.Store.UpsertProducts(&newProducts)
	return err
}

// ListProductStocks converts echo context to params.
func (ia *InventoryAPI) ListProductStocks(ctx echo.Context) error {
	// Invoke the callback with all the unmarshalled arguments
	result, err := ia.Store.ListProductStocks()
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, result)
}

// SellFromInventory converts echo context to params.
func (ia *InventoryAPI) SellFromInventory(ctx echo.Context) error {
	var err error
	var sellOrder api.SellOrder
	err = ctx.Bind(&sellOrder)
	if err != nil {
		return err
	}
	err = ia.Store.SellProducts(&sellOrder)
	return err
}
