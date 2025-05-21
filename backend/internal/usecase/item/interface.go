package item

import "github.com/Non2d/ArtVideoTube/backend/internal/domain/item"

type UseCase interface {
    GetAllItems() ([]item.Item, error)
    CreateItem(item item.Item) error
}