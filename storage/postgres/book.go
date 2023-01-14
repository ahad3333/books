package postgres

import (
	"add/models"
	"database/sql"
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/lib/pq"
	
)

type BookRepo struct {
	db *pgxpool.Pool
}

func NewBookRepo(db *pgxpool.Pool) *BookRepo {
	return &BookRepo{
		db: db,
	}
}

func (r *BookRepo) Insert(ctx context.Context, req *models.CreateBook) (string, error) {

	var (
		id = uuid.New().String()
	)

	query := `
	INSERT INTO books (
		id,
		name,
		price,
		description,
		updated_at
	) VALUES ($1, $2, $3, $4, now())
`

	_, err := r.db.Exec(ctx, query,
		id,
		req.Name,
		req.Price,
		req.Description,
	)

	if err != nil {
		return "", err
	}

	if len(req.CategoryIds) > 0 {

		bookCategoryQuery := `
				INSERT INTO book_category (
					category_id, 
					books_id
				) VALUES
		`

		for _, categoryId := range req.CategoryIds {
			bookCategoryQuery += fmt.Sprintf("('%s', '%s'),", categoryId, id)
		}

		bookCategoryQuery = bookCategoryQuery[:len(bookCategoryQuery)-1]

		_, err := r.db.Exec(ctx, bookCategoryQuery)
		if err != nil {
			return "", err
		}
	}

	return id, nil
}

func (r *BookRepo) GetByID(ctx context.Context, req *models.BookPrimeryKey) (*models.Book, error) {

	query := `
		SELECT
			b.id,
			b.name,
			b.price,
			b.description,
			b.created_at,
			b.updated_at,
			(
				SELECT
					ARRAY_AGG(category_id)
				FROM book_category AS bc 
				WHERE bc.books_id = $1
			) AS category_ids
		FROM
			books AS b
		WHERE b.id = $1
	`

	var (
		id          sql.NullString
		name        sql.NullString
		price       sql.NullFloat64
		description sql.NullString
		createdAt   sql.NullString
		updatedAt   sql.NullString
		categoryIds []string
	)

	err := r.db.QueryRow(ctx, query, req.Id).
		Scan(
			&id,
			&name,
			&price,
			&description,
			&createdAt,
			&updatedAt,
			(*pq.StringArray)(&categoryIds),
		)

	if err != nil {
		return nil, err
	}

	book := &models.Book{
		Id:          id.String,
		Name:        name.String,
		Price:       price.Float64,
		Description: description.String,
		CreatedAt:   createdAt.String,
		UpdatedAt:   updatedAt.String,
	}

	if len(categoryIds) > 0 {
		categoryQuery := `
			SELECT
				id,
				name,
				created_at,
				updated_at
			FROM
				category
			WHERE id IN (`

		for _, categoryId := range categoryIds {
			categoryQuery += fmt.Sprintf("'%s',", categoryId)
		}
		categoryQuery = categoryQuery[:len(categoryQuery)-1]
		categoryQuery += ")"

		rows, err := r.db.Query(ctx, categoryQuery)
		if err != nil {
			return nil, err
		}

		for rows.Next() {

			var (
				id        sql.NullString
				name      sql.NullString
				createdAt sql.NullString
				updatedAt sql.NullString
			)

			err = rows.Scan(
				&id,
				&name,
				&createdAt,
				&updatedAt,
			)

			book.Categories = append(book.Categories, &models.Category1{
				Id:        id.String,
				Name:      name.String,
				CreatedAt: createdAt.String,
				UpdatedAt: updatedAt.String,
			})
		}
	}

	return book, nil
}
func (r *BookRepo)GetList(ctx context.Context, req *models.GetListBookRequest) (*models.GetListBookResponse, error) {

	var (
		resp   models.GetListBookResponse
		offset = " OFFSET 0"
		limit  = " LIMIT 10"
	)

	query := `
		SELECT
			COUNT(*) OVER(),
			id,
			name,
			price,
			description,
			created_at,
			updated_at 
		FROM books
	`

	if req.Offset > 0 {
		offset = fmt.Sprintf(" OFFSET %d", req.Offset)
	}

	if req.Limit > 0 {
		limit = fmt.Sprintf(" LIMIT %d", req.Limit)
	}

	query += offset + limit
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return &models.GetListBookResponse{}, err
	}
	var (
		id          sql.NullString
		name        sql.NullString
		price       sql.NullFloat64
		description sql.NullString
		createdAt   sql.NullString
		updatedAt   sql.NullString
	)

	for rows.Next() {


		err = rows.Scan(
			&resp.Count,
			&id,
			&name,
			&price,
			&description,
			&createdAt,
			&updatedAt,
		)
		book := models.CategoryBook{
			Id:          id.String,
			Name:        name.String,
			Price:       price.Float64,
			Description: description.String,
			CreatedAt:   createdAt.String,
			UpdatedAt:   updatedAt.String,
		}
		if err != nil {
			return &models.GetListBookResponse{}, err
		}
		
		resp.Books = append(resp.Books, &book)


	}
	return &resp, nil
}


func (r *BookRepo)Update(ctx context.Context, book *models.UpdateBook) error {
	query := `
		UPDATE 
			books 
		SET 
			name = $2,
			price = $3,
			description = $4,
			updated_at = now()
		WHERE id = $1
	`

	_, err := r.db.Exec(ctx,query,
		book.Id,
		book.Name,
		book.Price,
		book.Description,
	)

	if err != nil {
		return err
	}

	return nil
}

func (r *BookRepo)Delete(ctx context.Context, req *models.BookPrimeryKey) error {
	_, err := r.db.Exec(ctx,"DELETE FROM book_category WHERE books_id  = $1 ", req.Id)
	if err != nil {
		return err
	}
	_, err = r.db.Exec(ctx,"DELETE FROM books WHERE id = $1", req.Id)

	if err != nil {
		return err
	}

	return nil
}
