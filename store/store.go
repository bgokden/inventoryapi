package store

import (
	"errors"
	"fmt"
	"strconv"
	"sync"

	"github.com/jinzhu/copier"

	"github.com/bgokden/inventoryapi/api"
)

// Store interface implements defines a common functions
type Store interface {
	UpsertInventory(*api.Inventory) error
	GetStock(string) (*api.Stock, error)
	ListInventory() (*api.Inventory, error)
	UpsertProducts(*api.Products) error
	ListProducts() (*api.Products, error)
	ListProductStocks() (*api.ProductStocks, error)
	SellProducts(*api.SellOrder) error
}

// InMemoryStore implements Store interface and holds data
type InMemoryStore struct {
	sync.RWMutex
	Inventory map[string]api.Stock
	Products  map[string]api.Product
}

// NewInMemoryStore creates an InMemoryStore with default values
func NewInMemoryStore() Store {
	return &InMemoryStore{
		Inventory: make(map[string]api.Stock, 0),
		Products:  make(map[string]api.Product, 0),
	}
}

// UpsertInventory updates in memory representation of the inventory
func (s *InMemoryStore) UpsertInventory(inventory *api.Inventory) error {
	s.Lock()
	defer s.Unlock()
	for _, additionalStockEntry := range *inventory.Inventory {
		additionalStock, err := strconv.Atoi(additionalStockEntry.Stock)
		if err != nil {
			return err
		}
		if stock, ok := s.Inventory[additionalStockEntry.ArtId]; ok {
			currentStock, err := strconv.Atoi(stock.Stock)
			if err != nil {
				return err
			}
			newStock := currentStock + additionalStock
			if newStock != 0 {
				s.Inventory[additionalStockEntry.ArtId] = api.Stock{
					ArtId: additionalStockEntry.ArtId,
					Name:  additionalStockEntry.Name, // It is possible to update name
					Stock: strconv.Itoa(newStock),
				}
			} else {
				delete(s.Inventory, additionalStockEntry.ArtId)
			}
		} else {
			s.Inventory[additionalStockEntry.ArtId] = additionalStockEntry
		}
	}
	return nil
}

// GetStock returns stock for an article
func (s *InMemoryStore) GetStock(artId string) (*api.Stock, error) {
	s.RLock()
	defer s.RUnlock()

	if stock, ok := s.Inventory[artId]; ok {
		stockCopy := api.Stock{}
		copier.Copy(&stockCopy, &stock)
		return &stockCopy, nil
	}

	return nil, errors.New(fmt.Sprintf("ArtId %v does not exists", artId))
}

// ListInventory return the current representation of the inventory
func (s *InMemoryStore) ListInventory() (*api.Inventory, error) {
	s.RLock()
	defer s.RUnlock()
	stocks := make([]api.Stock, 0, len(s.Inventory))
	for _, stock := range s.Inventory {
		stockCopy := api.Stock{}
		copier.Copy(&stockCopy, &stock)
		stocks = append(stocks, stock)
	}
	return &api.Inventory{
		Inventory: &stocks,
	}, nil
}

// ListProducts returns the current representation of the products
func (s *InMemoryStore) ListProducts() (*api.Products, error) {
	s.RLock()
	defer s.RUnlock()
	productList := make([]api.Product, 0, len(s.Products))
	for _, product := range s.Products {
		productCopy := api.Product{}
		copier.Copy(&productCopy, &product)
		productList = append(productList, product)
	}
	return &api.Products{
		Products: &productList,
	}, nil
}

// UpsertProducts updates in memory representation of the inventory
func (s *InMemoryStore) UpsertProducts(products *api.Products) error {
	s.Lock()
	defer s.Unlock()
	for _, product := range *products.Products {
		s.Products[product.Name] = product
	}
	return nil
}

// SellProducts processes a sell order and updates the inventory accordingly
func (s *InMemoryStore) SellProducts(sellOrder *api.SellOrder) error {
	s.Lock()
	defer s.Unlock()
	// Calculate and reduce sell order of multiple products to a map of article id to required stock
	inventoryChangeMap, err := s.CalculateInventoryChangeMap(sellOrder)
	if err != nil {
		return err
	}
	// First check if it is possible process this order
	// Without this a rollback would have been needed for failed orders.
	err = s.CheckInventory(inventoryChangeMap)
	if err != nil {
		return err
	}
	// Apply the actual order. Sice there is a write lock this process it atomic for inventory.
	return s.DecrementInventory(inventoryChangeMap)
}

// CalculateInventoryChangeMap calculates required articles based on multiple products.
// A sell order can hold multiple products and products can share the same article
// this method sums and return a map of article id to int values
func (s *InMemoryStore) CalculateInventoryChangeMap(sellOrder *api.SellOrder) (map[string]int, error) {
	inventoryChangeMap := make(map[string]int, len(sellOrder.Orders))
	for _, order := range sellOrder.Orders {
		if product, ok := s.Products[order.ProductName]; ok {
			for _, articles := range *product.ContainArticles {
				changePerArticle, err := strconv.Atoi(articles.AmountOf)
				if err != nil {
					return nil, err
				}
				change := changePerArticle * order.Number
				if currentChange, ok := inventoryChangeMap[articles.ArtId]; ok {
					change = change + currentChange
				}
				inventoryChangeMap[articles.ArtId] = change
			}
		} else {
			return nil, errors.New(fmt.Sprintf("Product %v doesn't exists", order.ProductName))
		}
	}
	return inventoryChangeMap, nil
}

// DecrementInventory updates the inventory with giveninventoryChangeMap
func (s *InMemoryStore) DecrementInventory(inventoryChangeMap map[string]int) error {
	return s.decrementInventoryWithCheck(inventoryChangeMap, true)
}

// CheckInventory checks if it is possible to update the inventory with giveninventoryChangeMap
func (s *InMemoryStore) CheckInventory(inventoryChangeMap map[string]int) error {
	return s.decrementInventoryWithCheck(inventoryChangeMap, false)
}

// decrementInventoryWithCheck allow updateing or checking the inventory with giveninventoryChangeMap
// When apply set to false it only check if it is possible to process a sell order
func (s *InMemoryStore) decrementInventoryWithCheck(inventoryChangeMap map[string]int, apply bool) error {
	for articleID, change := range inventoryChangeMap {
		if stock, ok := s.Inventory[articleID]; ok {
			currentStock, err := strconv.Atoi(stock.Stock)
			if err != nil {
				return err
			}
			if currentStock < change {
				return errors.New(fmt.Sprintf("There is not enough stock for Article %v", articleID))
			}
			if apply {
				newStock := currentStock - change
				if newStock != 0 {
					stock.Stock = strconv.Itoa(newStock)
					s.Inventory[articleID] = stock
				} else {
					delete(s.Inventory, articleID)
				}
			}
		} else {
			return errors.New(fmt.Sprintf("Article %v does not exist", articleID))
		}
	}
	return nil
}

// ListProductStocks list the currenly available products based on inventory
func (s *InMemoryStore) ListProductStocks() (*api.ProductStocks, error) {
	s.RLock()
	defer s.RUnlock()
	productStocks := make([]api.ProductStock, 0, len(s.Products))
	for _, product := range s.Products {
		stock := s.GetStockForProduct(&product)
		if stock > 0 {
			productStocks = append(productStocks, api.ProductStock{
				Product: product,
				Stock:   stock,
			})
		}
	}
	return &api.ProductStocks{
		Products: &productStocks,
	}, nil
}

// GetStockForProduct returns the possible stock for a product
// It returns the stock of the least available article in a product.
// So lack of an article can create bottlenack for a product
func (s *InMemoryStore) GetStockForProduct(product *api.Product) int {
	minimumNumberOfPossibleArticles := -1
	for _, articles := range *product.ContainArticles {
		changePerArticle, err := strconv.Atoi(articles.AmountOf)
		if err != nil || changePerArticle <= 0 {
			return 0
		}
		if currentArticleStock, ok := s.Inventory[articles.ArtId]; ok {
			currentArticleAvailability, err := strconv.Atoi(currentArticleStock.Stock)
			if err != nil || currentArticleAvailability == 0 {
				return 0
			}
			numberOfPossibleArticles := currentArticleAvailability / changePerArticle
			if numberOfPossibleArticles < minimumNumberOfPossibleArticles || minimumNumberOfPossibleArticles == -1 {
				minimumNumberOfPossibleArticles = numberOfPossibleArticles
			}
		} else {
			return 0
		}
	}
	if minimumNumberOfPossibleArticles == -1 {
		// This shows an input validation problem
		return 0
	}
	return minimumNumberOfPossibleArticles
}
