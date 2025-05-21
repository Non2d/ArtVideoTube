package item

import "github.com/Non2d/ArtVideoTube/backend/internal/domain/item"

type mockRepository struct {
	items []item.Item
}

func NewMockRepository() Repository {
	return &mockRepository{
		items: []item.Item{
			{ID: "1", Title: "First Video", Description: "Test video", URL: "https://example.com/1"},
		},
	}
}

func (r *mockRepository) GetAll() ([]item.Item, error) {
	return r.items, nil
}

func (r *mockRepository) Create(item item.Item) error {
	r.items = append(r.items, item)
	return nil
}
