package impl

import (
	"errors"
	"fmt"
	"strconv"
	"sync"

	"github.com/jinzhu/copier"

	"github.com/bgokden/inventoryapi/api"
)

type Store interface {
	UpsertInventory(*api.Inventory) error
	GetStock(string) (*api.Stock, error)
	ListInventory() (*api.Inventory, error)
	UpsertProducts(*api.Products) error
}

type InMemoryStore struct {
	sync.RWMutex
	Inventory map[string]api.Stock
	Products  map[string]api.Product
}

func NewInMemoryStore() Store {
	return &InMemoryStore{
		Inventory: make(map[string]api.Stock, 0),
	}
}

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

func (s *InMemoryStore) GetStock(artId string) (*api.Stock, error) {
	s.RLock()
	defer s.RUnlock()

	if stock, ok := s.Inventory[artId]; ok {
		stockCopy := api.Stock{}
		copier.Copy(&stockCopy, &stock)
		return &stockCopy, nil
	}

	return nil, errors.New(fmt.Sprintf("artId %v does not exists", artId))
}

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

func (s *InMemoryStore) UpsertProducts(products *api.Products) error {
	s.Lock()
	defer s.Unlock()
	for _, product := range *products.Products {
		s.Products[product.Name] = product
	}
	return nil
}

type Order struct {
	ProductName string
	Number      int
}

type SellOrder struct {
	Orders []Order
}

func (s *InMemoryStore) SellProducts(orders *SellOrder) error {
	inventoryChangeMap, err := s.CalculateInventoryChangeMap(orders)
	if err != nil {
		return err
	}
	err = s.CheckInventory(inventoryChangeMap)
	if err != nil {
		return err
	}
	return s.DecrementInventory(inventoryChangeMap)
}

func (s *InMemoryStore) CalculateInventoryChangeMap(orders *SellOrder) (map[string]int, error) {
	inventoryChangeMap := make(map[string]int, len(orders.Orders))
	for _, order := range orders.Orders {
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
				inventoryChangeMap[order.ProductName] = change
			}
		} else {
			return nil, errors.New(fmt.Sprintf("Product doesn't exists: %v\n", order.ProductName))
		}
	}
	return inventoryChangeMap, nil
}

func (s *InMemoryStore) DecrementInventory(inventoryChangeMap map[string]int) error {
	return s.decrementInventoryWithCheck(inventoryChangeMap, true)
}

func (s *InMemoryStore) CheckInventory(inventoryChangeMap map[string]int) error {
	return s.decrementInventoryWithCheck(inventoryChangeMap, false)
}

func (s *InMemoryStore) decrementInventoryWithCheck(inventoryChangeMap map[string]int, apply bool) error {
	s.Lock()
	defer s.Unlock()
	for articleID, change := range inventoryChangeMap {
		if stock, ok := s.Inventory[articleID]; ok {
			currentStock, err := strconv.Atoi(stock.Stock)
			if err != nil {
				return err
			}
			if currentStock < change {
				return errors.New(fmt.Sprintf("There is not enough stock for %v", articleID))
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
			return errors.New(fmt.Sprintf("Article %v does not exist.", articleID))
		}
	}
	return nil
}
