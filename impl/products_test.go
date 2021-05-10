package impl_test

import (
	"testing"

	"github.com/bgokden/inventoryapi/api"
	"github.com/bgokden/inventoryapi/impl"
	"github.com/stretchr/testify/assert"
)

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

func TestInMemoryStoreUpsertProductsAndListStocks(t *testing.T) {
	s := impl.NewInMemoryStore()
	inventory := &api.Inventory{
		Inventory: &stockList,
	}
	err := s.UpsertInventory(inventory)
	assert.Nil(t, err)

	products := &api.Products{
		Products: &productList,
	}

	err = s.UpsertProducts(products)
	assert.Nil(t, err)

	expectedProductStocksList := []api.ProductStock{
		{
			Product: productList[0],
			Stock:   2,
		},
		{
			Product: productList[1],
			Stock:   1,
		},
	}

	productStocks, err := s.ListProductStocks()
	assert.Nil(t, err)

	assert.Equal(t, expectedProductStocksList, *productStocks.Products)
}

func TestInMemoryStoreUpsertProductsAndSellAndListStocks(t *testing.T) {
	s := impl.NewInMemoryStore()
	inventory := &api.Inventory{
		Inventory: &stockList,
	}
	err := s.UpsertInventory(inventory)
	assert.Nil(t, err)

	products := &api.Products{
		Products: &productList,
	}

	err = s.UpsertProducts(products)
	assert.Nil(t, err)

	sellOrder := &api.SellOrder{
		Orders: []api.Order{
			{
				ProductName: productList[0].Name,
				Number:      1,
			},
		},
	}

	err = s.SellProducts(sellOrder)
	assert.Nil(t, err)

	expectedProductStocksList := []api.ProductStock{
		{
			Product: productList[0],
			Stock:   1,
		},
		{
			Product: productList[1],
			Stock:   1,
		},
	}

	listProducts, err := s.ListProducts()
	assert.Nil(t, err)

	assert.Equal(t, products.Products, listProducts.Products)

	productStocks, err := s.ListProductStocks()
	assert.Nil(t, err)

	assert.Equal(t, expectedProductStocksList, *productStocks.Products)
}

func TestInMemoryStoreUpsertProductsAndSellFail(t *testing.T) {
	s := impl.NewInMemoryStore()
	inventory := &api.Inventory{
		Inventory: &stockList,
	}
	err := s.UpsertInventory(inventory)
	assert.Nil(t, err)

	products := &api.Products{
		Products: &productList,
	}

	err = s.UpsertProducts(products)
	assert.Nil(t, err)

	sellOrder := &api.SellOrder{
		Orders: []api.Order{
			{
				ProductName: productList[0].Name,
				Number:      3,
			},
		},
	}

	err = s.SellProducts(sellOrder)
	assert.EqualError(t, err, "There is not enough stock for Article 2")
}

func TestInMemoryStoreUpsertProductsAndSellAndRemoveFromListStocks(t *testing.T) {
	s := impl.NewInMemoryStore()
	inventory := &api.Inventory{
		Inventory: &stockList,
	}
	err := s.UpsertInventory(inventory)
	assert.Nil(t, err)

	products := &api.Products{
		Products: &productList,
	}

	err = s.UpsertProducts(products)
	assert.Nil(t, err)

	sellOrder := &api.SellOrder{
		Orders: []api.Order{
			{
				ProductName: productList[1].Name,
				Number:      1,
			},
		},
	}

	err = s.SellProducts(sellOrder)
	assert.Nil(t, err)

	expectedProductStocksList := []api.ProductStock{
		{
			Product: productList[0],
			Stock:   1,
		},
	}

	productStocks, err := s.ListProductStocks()
	assert.Nil(t, err)

	assert.Equal(t, expectedProductStocksList, *productStocks.Products)
}
