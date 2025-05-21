package item

import "github.com/Non2d/ArtVideoTube/backend/internal/domain/item"

type Repository interface {
    GetAll() ([]item.Item, error)
    Create(item item.Item) error
}