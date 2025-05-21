package item

import "github.com/Non2d/ArtVideoTube/backend/internal/domain/item"

type memoryRepository struct {
	items []item.Item
}

func NewMemoryRepository() Repository {
	return &memoryRepository{
		items: []item.Item{
			{ID: "1", Title: "First Video", Description: "Test video", URL: "https://example.com/1"},
		},
	}
}

func (r *memoryRepository) GetAll() ([]item.Item, error) {
	return r.items, nil
}

func (r *memoryRepository) Create(item item.Item) error {
	r.items = append(r.items, item)
	return nil
}
