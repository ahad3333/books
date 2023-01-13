package postgres

import (
	"add/models"
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/lib/pq"
)
type categoryRepo struct {
	db *pgxpool.Pool
}

func NewCategoryRepo(db *pgxpool.Pool) *categoryRepo {
	return &categoryRepo{
		db: db,
	}

}

func (r *categoryRepo) Insert(ctx context.Context, req *models.CreateCategory) (string, error) {
	var (
		id = uuid.New().String()
	)

	query := `
	INSERT INTO category (
		id,
		name,
		updated_at
	) VALUES ($1, $2, now())
`

	_, err := r.db.Exec(ctx, query,
		id,
		req.Name,
	)

	if err != nil {
		return "", err
	}

	if len(req.Books_id) > 0 {

		CategoryBookQuery := `
				INSERT INTO book_category (
					books_id,
					category_id
				) VALUES
		`

		for _, booksId:= range req.Books_id {
			CategoryBookQuery += fmt.Sprintf("('%s', '%s'),", booksId, id)
		}

		CategoryBookQuery = CategoryBookQuery[:len(CategoryBookQuery)-1]

		_, err := r.db.Exec(ctx, CategoryBookQuery)
		if err != nil {
			return "", err
		}
	}

	return id, nil
}

func (r *categoryRepo) GetByID(ctx context.Context, req *models.CategoryPrimeryKey) (*models.Category, error) {
	query := `
	SELECT
		c.id,
		c.name,
		c.created_at,
		updated_at,
		(
			SELECT
				ARRAY_AGG(books_id)
			FROM book_category AS bc 
			WHERE bc.category_id = $1
		) AS Books_id
	FROM
		category AS c
	WHERE c.id = $1
`

var (
	id          sql.NullString
	name        sql.NullString
	createdAt   sql.NullString
	updatedAt   sql.NullString
	book_ids []string
)

err := r.db.QueryRow(ctx, query, req.Id).
	Scan(
		&id,
		&name,
		&createdAt,
		&updatedAt,
		(*pq.StringArray)(&book_ids),
	)

if err != nil {
	return nil, err
}

category := &models.Category{
	Id:          id.String,
	Name:        name.String,
	CreatedAt:   createdAt.String,
	UpdatedAt:   updatedAt.String,
}

if len(book_ids) > 0 {
	booksQuery := `
		SELECT
			id,
			name,
			price,
			description,
			created_at,
			updated_at
		FROM
			books
		WHERE id IN (`

	for _, bookId := range book_ids {
		booksQuery += fmt.Sprintf("'%s',", bookId)
	}
	booksQuery = booksQuery[:len(booksQuery)-1]
	booksQuery += ")"

	rows, err := r.db.Query(ctx, booksQuery)
	if err != nil {
		return nil, err
	}

	for rows.Next() {

		var (
			id        sql.NullString
			name      sql.NullString
			price       sql.NullFloat64
			description sql.NullString
			createdAt sql.NullString
			updatedAt sql.NullString
		)

		err = rows.Scan(
			&id,
			&name,
			&price,
			&description,
			&createdAt,
			&updatedAt,
		)

		category.Books = append(category.Books, &models.CategoryBook{
			Id:        id.String,
			Name:      name.String,
			Price: 		price.Float64,
			Description: description.String,
			CreatedAt: createdAt.String,
			UpdatedAt: updatedAt.String,
		})
	}
}

return category, nil
}


func (r *categoryRepo)GetList(ctx context.Context, req *models.GetListCategoryRequest) (*models.GetListCategoryResponse, error) {

	var (
		resp   models.GetListCategoryResponse
		offset = " OFFSET 0"
		limit  = " LIMIT 10"
	)

	query := `
		SELECT
			COUNT(*) OVER(),
			id,
			name
		FROM Category
	`

	if req.Offset > 0 {
		offset = fmt.Sprintf(" OFFSET %d", req.Offset)
	}

	if req.Limit > 0 {
		limit = fmt.Sprintf(" LIMIT %d", req.Limit)
	}

	query += offset + limit

	rows, err := r.db.Query(ctx,query)
	if err != nil {
		return &models.GetListCategoryResponse{}, err
	}
	var (
		id        sql.NullString
		name      sql.NullString
		price       sql.NullFloat64
		description sql.NullString
		createdAt sql.NullString
		updatedAt sql.NullString
	)

	for rows.Next() {
		err = rows.Scan(
			&id,
			&name,
			&price,
			&description,
			&createdAt,
			&updatedAt,
		)
		if err != nil {
			return &models.GetListCategoryResponse{}, err
		}
		category := models.Category1{
			Id:          id.String,
			Name:        name.String,
			CreatedAt:   createdAt.String,
			UpdatedAt:   updatedAt.String,
		}
		resp.Categories = append(resp.Categories, &category)

	}


	return &resp, nil
}

func (r *categoryRepo)Update(ctx context.Context, Category *models.UpdateCategory) error {

	query := `
		UPDATE 
		Category
		SET 
			name = $2
		WHERE id = $1
	`

	_, err := r.db.Exec(ctx,query,
		Category.Id,
		Category.Name,
	)

	if err != nil {
		return err
	}

	return nil
}

func (r *categoryRepo)Delete(ctx context.Context, req *models.CategoryPrimeryKey) error {
	_, err := r.db.Exec(ctx,"DELETE FROM book_category WHERE category_id  = $1 ", req.Id)
	if err != nil {
		return err
	}
	_, err = r.db.Exec(ctx,"DELETE FROM Category WHERE id = $1 ", req.Id)

	if err != nil {
		return err
	}

	return nil
}
