package util

import (
	"sort"

	"github.com/bgokden/inventoryapi/api"
)

/*
These are utility functions that are added to be used in the tests
and it is possible to use them in non-test code in the future.
*/

// SortInventoryList sorts given list of api.Stock based on ArtId
func SortInventoryList(list []api.Stock) []api.Stock {
	sort.Slice(list, func(i, j int) bool {
		return list[i].ArtId < list[j].ArtId
	})
	return list
}

// SortProductStocksList sorts given list of api.ProductStock based on Product.Name
func SortProductStocksList(list []api.ProductStock) []api.ProductStock {
	sort.Slice(list, func(i, j int) bool {
		return list[i].Product.Name < list[j].Product.Name
	})
	return list
}

// SortProductList sorts given list of api.Product based on Name
func SortProductList(list []api.Product) []api.Product {
	sort.Slice(list, func(i, j int) bool {
		return list[i].Name < list[j].Name
	})
	return list
}
