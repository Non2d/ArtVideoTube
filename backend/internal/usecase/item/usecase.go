package item

import (
	"github.com/Non2d/ArtVideoTube/backend/internal/domain/item"
	itemRepo "github.com/Non2d/ArtVideoTube/backend/internal/repository/item"
)

type useCase struct {
	repo itemRepo.Repository
}

func NewUseCase(repo itemRepo.Repository) UseCase {
	return &useCase{repo: repo}
}

func (uc *useCase) GetAllItems() ([]item.Item, error) {
	return uc.repo.GetAll()
}

func (uc *useCase) CreateItem(item item.Item) error {
	return uc.repo.Create(item)
}
