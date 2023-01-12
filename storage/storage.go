package storage

import (
	"context"
	"add/models"
)

type StorageI interface {
	CloseDB()
	Book() BookRepoI
	Category() CategoryRepoI
}

type BookRepoI interface {
	Insert(context.Context, *models.CreateBook) (string, error)
	GetByID(context.Context, *models.BookPrimeryKey) (*models.Book, error)
	GetList(ctx context.Context, req *models.GetListBookRequest) (*models.GetListBookResponse, error)
	Update(ctx context.Context,book *models.UpdateBook) error
	Delete(ctx context.Context, req *models.BookPrimeryKey) error 
}

type CategoryRepoI interface {
	Insert(context.Context, *models.CreateCategory) (string, error)
	GetByID(context.Context, *models.CategoryPrimeryKey) (*models.Category, error)
	GetList(ctx context.Context, req *models.GetListCategoryRequest) (*models.GetListCategoryResponse, error)
	Update(ctx context.Context, category *models.UpdateCategory) error
	Delete(ctx context.Context, req *models.CategoryPrimeryKey) error

}