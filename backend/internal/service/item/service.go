package item

import (
    "github.com/Non2d/ArtVideoTube/backend/internal/domain/item"
    itemUseCase "github.com/Non2d/ArtVideoTube/backend/internal/usecase/item"
)

type Service struct {
    useCase itemUseCase.UseCase
}

func NewService(uc itemUseCase.UseCase) *Service {
    return &Service{useCase: uc}
}

func (s *Service) GetItems() ([]item.Item, error) {
    return s.useCase.GetAllItems()
}

func (s *Service) CreateItem(item item.Item) error {
    return s.useCase.CreateItem(item)
}