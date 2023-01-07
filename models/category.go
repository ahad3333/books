package models

type CategoryPrimeryKey struct {
	Id string `json:"id"`
}

type CreateCategory struct {
	Name        string  `json:"name"`
}

type Category struct {
	Id          string  `json:"id"`
	Name        string  `json:"name"`
}

type UpdateCategory struct {
	Id          string  `json:"id"`
	Name        string  `json:"name"`
}

type GetListCategoryRequest struct {
	Offset int64 `json:"offset"`
	Limit  int64 `json:"limit"`
}

type GetListCategoryResponse struct {
	Count int64  `json:"count"`
	Categorys []Category `json:"Categorys"`
}