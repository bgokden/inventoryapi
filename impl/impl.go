package impl

import (
	"net/http"

	"github.com/bgokden/inventoryapi/api"
	"github.com/labstack/echo/v4"
)

type InventoryAPI struct {
	Store Store
}

func NewInventoryAPI() *InventoryAPI {
	ia := &InventoryAPI{
		Store: NewInMemoryStore(),
	}
	input := &api.Inventory{
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
	ia.Store.UpsertInventory(input)

	return ia
}

// GetInventory converts echo context to params.
func (w *InventoryAPI) GetInventory(ctx echo.Context) error {
	// Invoke the callback with all the unmarshalled arguments
	result, err := w.Store.ListInventory()
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, result)
}

// UpsertInventory converts echo context to params.
func (w *InventoryAPI) UpsertInventory(ctx echo.Context) error {
	var err error
	var newInventory api.Inventory
	err = ctx.Bind(&newInventory)
	if err != nil {
		return err
	}
	err = w.Store.UpsertInventory(&newInventory)
	return err
}

/*
curl -v --header "Content-Type: application/json" \
  --request POST \
  --data '{"inventory":[{"art_id":"1","name":"test1","stock":"5"},{"art_id":"2","name":"test2","stock":"3"}]}' \
  http://localhost:8080/inventory
*/

// ListProducts converts echo context to params.
func (w *InventoryAPI) ListProducts(ctx echo.Context) error {
	// Invoke the callback with all the unmarshalled arguments
	result, err := w.Store.ListProducts()
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, result)
}

// UpsertProducts converts echo context to params.
func (w *InventoryAPI) UpsertProducts(ctx echo.Context) error {
	var err error
	var newProducts api.Products
	err = ctx.Bind(&newProducts)
	if err != nil {
		return err
	}
	err = w.Store.UpsertProducts(&newProducts)
	return err
}

// ListProductStocks converts echo context to params.
func (w *InventoryAPI) ListProductStocks(ctx echo.Context) error {
	// Invoke the callback with all the unmarshalled arguments
	result, err := w.Store.ListProductStocks()
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, result)
}

// SellFromInventory converts echo context to params.
func (w *InventoryAPI) SellFromInventory(ctx echo.Context) error {
	var err error
	var sellOrder api.SellOrder
	err = ctx.Bind(&sellOrder)
	if err != nil {
		return err
	}
	err = w.Store.SellProducts(&sellOrder)
	return err
}
