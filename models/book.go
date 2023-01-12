package models

type BookPrimeryKey struct {
	Id string `json:"id"`
}

type CreateBook struct {
	Name        string   `json:"name"`
	Price       float64  `json:"price"`
	Description string   `json:"description"`
	CategoryIds []string `json:"category_ids"`
}
type CategoryBook struct {
	Id          string      `json:"id"`
	Name        string      `json:"name"`
	Price       float64     `json:"price"`
	Description string      `json:"description"`
	CreatedAt   string      `json:"created_at"`
	UpdatedAt   string      `json:"updated_at"`
}
type Book struct {
	Id          string      `json:"id"`
	Name        string      `json:"name"`
	Price       float64     `json:"price"`
	Description string      `json:"description"`
	CreatedAt   string      `json:"created_at"`
	UpdatedAt   string      `json:"updated_at"`
	Categories  []*Category1 `json:"categories"`
}

type UpdateBook struct {
	Id          string  `json:"id"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
}

type UpdateBookSwag struct {
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
}

type GetListBookRequest struct {
	Offset int64 `json:"offset"`
	Limit  int64 `json:"limit"`
}

type GetListBookResponse struct {
	Count int64   `json:"count"`
	Books []*UpdateBook `json:"books"`
}

type Empty struct{}
