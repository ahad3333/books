package storage

import (
	"add/models"
	"database/sql"
	"fmt"
	"github.com/google/uuid"
)


func InsertBook(db *sql.DB, book models.CreateBook) (string, error) {

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

	_, err := db.Exec(query,
		id,
		book.Name,
		book.Price,
		book.Description,
	)

	if err != nil {
		return "", err
	}

	return id, nil
}

func GetByIdBook(db *sql.DB, req models.BookPrimeryKey) (models.Book1, error) {

	var (
		book models.Book1
	)

	query := `
	select
	b.id,
	b.name,
	b.price,
	b.Description,
	b.updated_at,
	b.created_at
	from BookCategory as cb
	join books as b on b.id = cb.bookId
	where b.id = $1
	group by b.name,b.id;
	`
query1 :=`select
c.id,
c.name
from BookCategory as cb
join categorys as c on c.id = cb.categoryId
where cb.bookId = $1`
rows, err := db.Query(query,req.Id)
if err != nil {
	return models.Book1{}, err
}

for rows.Next() {
err =  rows.Scan(
	&book.Id,
	&book.Name,
	&book.Price,
	&book.Description,
	&book.UpdatedAt,
	&book.CreatedAt,	
)
rows, err = db.Query(query1,req.Id)
for rows.Next() {
	var cato models.CategoryName
	err =  rows.Scan(
		&cato.Id,
		&cato.Name,
		)
book.Categorys = append(book.Categorys,cato)		
}
}
	if err != nil {
		return models.Book1{}, err
	}

	return book, nil

}


func GetListBook(db *sql.DB, req models.GetListBookRequest) (models.GetListBookResponse, error) {

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

	rows, err := db.Query(query)
	if err != nil {
		return models.GetListBookResponse{}, err
	}

	for rows.Next() {
		var book models.Book

		err = rows.Scan(
			&resp.Count,
			&book.Id,
			&book.Name,
			&book.Price,
			&book.Description,
			&book.CreatedAt,
			&book.UpdatedAt,
		)

		if err != nil {
			return models.GetListBookResponse{}, err
		}

		resp.Books = append(resp.Books, book)
	}

	return resp, nil
}

func UpdateBook(db *sql.DB, book models.UpdateBook) error {

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

	_, err := db.Exec(query,
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

func DeleteBook(db *sql.DB, req models.BookPrimeryKey)  error {
	_, err := db.Exec("DELETE FROM books WHERE id = $1", req.Id)

	if err != nil {
		return err
	}

	return nil
}

