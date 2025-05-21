package handler

import (
    "encoding/json"
    "net/http"
    
    "github.com/Non2d/ArtVideoTube/backend/internal/domain/item"
    itemService "github.com/Non2d/ArtVideoTube/backend/internal/service/item"
)

type ItemHandler struct {
    service *itemService.Service
}

func NewItemHandler(service *itemService.Service) *ItemHandler {
    return &ItemHandler{service: service}
}

func (h *ItemHandler) GetItems(w http.ResponseWriter, r *http.Request) {
    items, err := h.service.GetItems()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(items)
}

func (h *ItemHandler) CreateItem(w http.ResponseWriter, r *http.Request) {
    var newItem item.Item
    if err := json.NewDecoder(r.Body).Decode(&newItem); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    
    if err := h.service.CreateItem(newItem); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(newItem)
}
