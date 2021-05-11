package server

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/bgokden/inventoryapi/api"
	"github.com/bgokden/inventoryapi/store"
)

// InventoryAPI implements api.ServerInterface and holds a store
type InventoryAPI struct {
	Store store.Store
}

func NewInventoryAPI() *InventoryAPI {
	return &InventoryAPI{
		// Here is it hardcoded to use a in memory store, in production this can be a database store
		// And it can be parametrized.
		Store: store.NewInMemoryStore(),
	}
}

// GetInventory converts echo context to params and calls store get inventory
func (ia *InventoryAPI) GetInventory(ctx echo.Context) error {
	result, err := ia.Store.ListInventory()
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, result)
}

// UpsertInventory converts echo context to params and calls store upsert inventory
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

// ListProducts converts echo context to params and calls store list products
func (ia *InventoryAPI) ListProducts(ctx echo.Context) error {
	result, err := ia.Store.ListProducts()
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, result)
}

// UpsertProducts converts echo context to params and calls store upsert products
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

// ListProductStocks converts echo context to params and calls Store List Product Stocks
func (ia *InventoryAPI) ListProductStocks(ctx echo.Context) error {
	result, err := ia.Store.ListProductStocks()
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, result)
}

// SellFromInventory converts echo context to params and calls Store Sell Products
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
