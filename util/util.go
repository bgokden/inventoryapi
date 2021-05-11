package util

import (
	"sort"

	"github.com/bgokden/inventoryapi/api"
)

func SortInventoryList(list []api.Stock) []api.Stock {
	sort.Slice(list, func(i, j int) bool {
		return list[i].ArtId < list[j].ArtId
	})
	return list
}

func SortProductStocksList(list []api.ProductStock) []api.ProductStock {
	sort.Slice(list, func(i, j int) bool {
		return list[i].Product.Name < list[j].Product.Name
	})
	return list
}

func SortProductList(list []api.Product) []api.Product {
	sort.Slice(list, func(i, j int) bool {
		return list[i].Name < list[j].Name
	})
	return list
}
