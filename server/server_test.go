package server_test

import (
	"net/http"
	"testing"

	"github.com/deepmap/oapi-codegen/pkg/testutil"
	"github.com/stretchr/testify/assert"

	"github.com/bgokden/inventoryapi/api"
	"github.com/bgokden/inventoryapi/server"
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

var productList = []api.Product{
	{
		Name: "Dining Chair",
		ContainArticles: &[]api.Article{
			{
				ArtId:    "1",
				AmountOf: "4",
			},
			{
				ArtId:    "2",
				AmountOf: "8",
			},
			{
				ArtId:    "3",
				AmountOf: "1",
			},
		},
	},
	{
		Name: "Dinning Table",
		ContainArticles: &[]api.Article{
			{
				ArtId:    "1",
				AmountOf: "4",
			},
			{
				ArtId:    "2",
				AmountOf: "8",
			},
			{
				ArtId:    "4",
				AmountOf: "1",
			},
		},
	},
}

var stockList = []api.Stock{
	{
		ArtId: "1",
		Name:  "leg",
		Stock: "12",
	},
	{
		ArtId: "2",
		Name:  "screw",
		Stock: "17",
	},
	{
		ArtId: "3",
		Name:  "seat",
		Stock: "2",
	},
	{
		ArtId: "4",
		Name:  "table top",
		Stock: "1",
	},
}

func TestUpsertInventoryAndGetInventory(t *testing.T) {
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

	// Test the getter function.
	result = testutil.NewRequest().Get("/v0/inventory").WithAcceptJson().Go(t, e)
	var resultInventory api.Inventory
	err = result.UnmarshalBodyToObject(&resultInventory)
	assert.NoError(t, err, "error getting pet")
	assert.Equal(t, newInventory, resultInventory)

	addtionalInventory := api.Inventory{
		Inventory: &[]api.Stock{
			{
				ArtId: "1",
				Name:  "test1_updated",
				Stock: "3",
			},
			{
				ArtId: "2",
				Name:  "test2",
				Stock: "-3",
			},
			{
				ArtId: "3",
				Name:  "test3",
				Stock: "2",
			},
		},
	}
	result = testutil.NewRequest().Post("/v0/inventory").WithJsonBody(addtionalInventory).Go(t, e)
	// We expect 200 code on successful pet insertion
	assert.Equal(t, http.StatusOK, result.Code())

	expectedInventory := api.Inventory{
		Inventory: &[]api.Stock{
			{
				ArtId: "1",
				Name:  "test1_updated",
				Stock: "8",
			},
			{
				ArtId: "3",
				Name:  "test3",
				Stock: "2",
			},
		},
	}

	// Test the getter function.
	result = testutil.NewRequest().Get("/v0/inventory").WithAcceptJson().Go(t, e)
	err = result.UnmarshalBodyToObject(&resultInventory)
	assert.NoError(t, err, "error getting pet")
	assert.Equal(t, expectedInventory, resultInventory)
}

func TestUpsertProductsAndListProducts(t *testing.T) {
	var err error
	// Here, we Initialize echo
	e, err := server.CreateEchoServer()
	assert.Nil(t, err)

	newProducts := api.Products{
		Products: &productList,
	}
	result := testutil.NewRequest().Post("/v0/products").WithJsonBody(newProducts).Go(t, e)
	// We expect 200 code on successful pet insertion
	assert.Equal(t, http.StatusOK, result.Code())

	// Test the getter function.
	result = testutil.NewRequest().Get("/v0/products").WithAcceptJson().Go(t, e)
	var resultProducts api.Products
	err = result.UnmarshalBodyToObject(&resultProducts)
	assert.NoError(t, err, "error getting pet")
	assert.Equal(t, newProducts, resultProducts)

	expectedProducts := api.Products{
		Products: &productList,
	}

	// Test the getter function.
	result = testutil.NewRequest().Get("/v0/products").WithAcceptJson().Go(t, e)
	err = result.UnmarshalBodyToObject(&resultProducts)
	assert.NoError(t, err, "error getting pet")
	assert.Equal(t, expectedProducts, resultProducts)
}

func TestUpsertsAndListProductStocksAndSellProduct(t *testing.T) {
	var err error
	// Here, we Initialize echo
	e, err := server.CreateEchoServer()
	assert.Nil(t, err)

	newProducts := api.Products{
		Products: &productList,
	}
	result := testutil.NewRequest().Post("/v0/products").WithJsonBody(newProducts).Go(t, e)
	// We expect 200 code on successful pet insertion
	assert.Equal(t, http.StatusOK, result.Code())

	newInventory := api.Inventory{
		Inventory: &stockList,
	}

	result = testutil.NewRequest().Post("/v0/inventory").WithJsonBody(newInventory).Go(t, e)
	// We expect 200 code on successful pet insertion
	assert.Equal(t, http.StatusOK, result.Code())

	expectedProductStockList := []api.ProductStock{
		{
			Product: productList[0],
			Stock:   2,
		},
		{
			Product: productList[1],
			Stock:   1,
		},
	}
	expectedProductStocks := api.ProductStocks{
		Products: &expectedProductStockList,
	}
	// Test the getter function.
	result = testutil.NewRequest().Get("/v0/productstock").WithAcceptJson().Go(t, e)
	var resultProductStocks api.ProductStocks
	err = result.UnmarshalBodyToObject(&resultProductStocks)
	assert.NoError(t, err, "error getting pet")
	assert.Equal(t, expectedProductStocks, resultProductStocks)

	sellOrder := &api.SellOrder{
		Orders: []api.Order{
			{
				ProductName: productList[1].Name,
				Number:      1,
			},
		},
	}

	result = testutil.NewRequest().Post("/v0/sell").WithJsonBody(sellOrder).Go(t, e)
	// We expect 200 code on successful pet insertion
	assert.Equal(t, http.StatusOK, result.Code())

	// Product stocks after sale
	expectedProductStockList2 := []api.ProductStock{
		{
			Product: productList[0],
			Stock:   1,
		},
	}
	expectedProductStocks2 := api.ProductStocks{
		Products: &expectedProductStockList2,
	}

	result = testutil.NewRequest().Get("/v0/productstock").WithAcceptJson().Go(t, e)
	err = result.UnmarshalBodyToObject(&resultProductStocks)
	assert.NoError(t, err, "error getting pet")
	assert.Equal(t, expectedProductStocks2, resultProductStocks)
}
