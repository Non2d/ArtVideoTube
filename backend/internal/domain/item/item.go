package item

type Item struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	URL         string `json:"url"`
}

type Repository interface {
	GetAll() ([]Item, error)
	Create(item Item) error
}
