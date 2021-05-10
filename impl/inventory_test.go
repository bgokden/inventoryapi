package impl_test

import (
	"testing"

	"github.com/bgokden/inventoryapi/api"
	"github.com/bgokden/inventoryapi/impl"
	"github.com/stretchr/testify/assert"
)

func TestInMemoeryStoreUpsertAndList(t *testing.T) {
	s := impl.NewInMemoryStore()
	input := &api.Inventory{
		Inventory: &[]api.Stock{
			{
				ArtId: "1",
				Name:  "test1",
				Stock: "1",
			},
		},
	}
	err := s.UpsertInventory(input)
	assert.Nil(t, err)

	output, err := s.ListInventory()
	assert.Nil(t, err)
	assert.Equal(t, input, output)
}

func TestInMemoeryStoreUpsertNameChangeAndList(t *testing.T) {
	s := impl.NewInMemoryStore()
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
	err := s.UpsertInventory(input)
	assert.Nil(t, err)

	sell := &api.Inventory{
		Inventory: &[]api.Stock{
			{
				ArtId: "1",
				Name:  "test1_updated",
				Stock: "0",
			},
		},
	}
	err = s.UpsertInventory(sell)
	assert.Nil(t, err)

	expected := &api.Inventory{
		Inventory: &[]api.Stock{
			{
				ArtId: "1",
				Name:  "test1_updated",
				Stock: "5",
			},
			{
				ArtId: "2",
				Name:  "test2",
				Stock: "3",
			},
		},
	}

	output, err := s.ListInventory()
	assert.Nil(t, err)
	assert.Equal(t, expected, output)
}

func TestInMemoeryStoreUpsertNegativeAndList(t *testing.T) {
	s := impl.NewInMemoryStore()
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
	err := s.UpsertInventory(input)
	assert.Nil(t, err)

	sell := &api.Inventory{
		Inventory: &[]api.Stock{
			{
				ArtId: "1",
				Name:  "test1",
				Stock: "-1",
			},
		},
	}
	err = s.UpsertInventory(sell)
	assert.Nil(t, err)

	expected := &api.Inventory{
		Inventory: &[]api.Stock{
			{
				ArtId: "1",
				Name:  "test1",
				Stock: "4",
			},
			{
				ArtId: "2",
				Name:  "test2",
				Stock: "3",
			},
		},
	}

	output, err := s.ListInventory()
	assert.Nil(t, err)
	assert.Equal(t, expected, output)
}

func TestInMemoeryStoreUpsertDeleteAndList(t *testing.T) {
	s := impl.NewInMemoryStore()
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
	err := s.UpsertInventory(input)
	assert.Nil(t, err)

	sell := &api.Inventory{
		Inventory: &[]api.Stock{
			{
				ArtId: "1",
				Name:  "test1",
				Stock: "-5",
			},
		},
	}
	err = s.UpsertInventory(sell)
	assert.Nil(t, err)

	expected := &api.Inventory{
		Inventory: &[]api.Stock{
			{
				ArtId: "2",
				Name:  "test2",
				Stock: "3",
			},
		},
	}

	output, err := s.ListInventory()
	assert.Nil(t, err)
	assert.Equal(t, expected, output)
}

func TestInMemoeryStoreUpsertDeleteAndGetStock(t *testing.T) {
	s := impl.NewInMemoryStore()
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
	err := s.UpsertInventory(input)
	assert.Nil(t, err)

	stock1, err := s.GetStock("1")
	expectedStock1 := &api.Stock{
		ArtId: "1",
		Name:  "test1",
		Stock: "5",
	}
	assert.Equal(t, expectedStock1, stock1)

	stock2, err := s.GetStock("2")
	expectedStock2 := &api.Stock{
		ArtId: "2",
		Name:  "test2",
		Stock: "3",
	}
	assert.Equal(t, expectedStock2, stock2)
}
