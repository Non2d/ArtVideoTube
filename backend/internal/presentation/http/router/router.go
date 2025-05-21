package router

import (
	"github.com/Non2d/ArtVideoTube/backend/internal/presentation/http/handler"
	itemRepo "github.com/Non2d/ArtVideoTube/backend/internal/repository/item"
	itemService "github.com/Non2d/ArtVideoTube/backend/internal/service/item"
	itemUseCase "github.com/Non2d/ArtVideoTube/backend/internal/usecase/item"

	"github.com/gorilla/mux"
)

func New() *mux.Router {
	// Initialize dependencies for the "item" domain
	repo := itemRepo.NewMockRepository()
	useCase := itemUseCase.NewUseCase(repo)
	service := itemService.NewService(useCase)
	h := handler.NewItemHandler(service)

	// Setup router
	r := mux.NewRouter()
	r.HandleFunc("/items", h.GetItems).Methods("GET")
	r.HandleFunc("/items", h.CreateItem).Methods("POST")

	return r
}
