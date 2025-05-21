package handlers

import (
    "net/http"
    "encoding/json"
    "github.com/gorilla/mux"
    "go-render-api/internal/models"
)

var items []models.Item

// GetItems retrieves all items
func GetItems(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(items)
}

// CreateItem adds a new item to the list
func CreateItem(w http.ResponseWriter, r *http.Request) {
    var newItem models.Item
    if err := json.NewDecoder(r.Body).Decode(&newItem); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    items = append(items, newItem)
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(newItem)
}

// RegisterRoutes sets up the API routes
func RegisterRoutes(r *mux.Router) {
    r.HandleFunc("/items", GetItems).Methods("GET")
    r.HandleFunc("/items", CreateItem).Methods("POST")
}