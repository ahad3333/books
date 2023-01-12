package models

type CategoryPrimeryKey struct {
	Id string `json:"id"`
}

type CreateCategory struct {
	Name string `json:"name"`
	Books_id []string `json:"books_id"`
}
type Category struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	Books  []*CategoryBook `json:"categories"`

}
type Category1 struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type UpdateCategory struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type UpdateCategorySwag struct {
	Name string `json:"name"`
}

type GetListCategoryRequest struct {
	Offset int64 `json:"offset"`
	Limit  int64 `json:"limit"`
}

type GetListCategoryResponse struct {
	Count      int64       `json:"count"`
	Categories []*UpdateCategory `json:"categories"`
}
