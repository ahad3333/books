package storage

import (
	"add/models"
	"database/sql"
	"fmt"
	"github.com/google/uuid"
)


func InsertCategory(db *sql.DB, Category models.CreateCategory) (string, error) {

	var (
		id = uuid.New().String()
	)

	query := `
		INSERT INTO Categorys (
			id,
			name
		) VALUES ($1, $2)
	`

	_, err := db.Exec(query,
		id,
		Category.Name,
	)

	if err != nil {
		return "", err
	}

	return id, nil
}

func GetByIdCategory(db *sql.DB, req models.CategoryPrimeryKey) (models.Category1, error) {

	var (
		Category models.Category1
	)
	query :=`select
	c.id,
	c.name
	from BookCategory as cb
	join categorys as c on c.id = cb.categoryId
	where cb.categoryId = $1`
	query1 := `
	select
	b.id,
	b.name,
	b.price,
	b.Description
	from BookCategory as cb
	join books as b on b.id = cb.bookId
	where cb.categoryId = $1
	group by b.name,b.id;
	`
rows, err := db.Query(query,req.Id)
if err != nil {
	return models.Category1{}, err
}

for rows.Next() {
err =  rows.Scan(
	&Category.Id,
	&Category.Name,	
)
rows, err = db.Query(query1,req.Id)
for rows.Next() {
	var book models.BookInfo
	err =  rows.Scan(
		&book.Id,
		&book.Name,
		&book.Price,
		&book.Description,
		)
Category.BookInfos = append(Category.BookInfos,book)		
}
}
	if err != nil {
		return models.Category1{}, err
	}

	return Category, nil
}

func GetListCategory(db *sql.DB, req models.GetListCategoryRequest) (models.GetListCategoryResponse, error) {

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
		FROM Categorys
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
		return models.GetListCategoryResponse{}, err
	}

	for rows.Next() {
		var Category models.Category

		err = rows.Scan(
			&resp.Count,
			&Category.Id,
			&Category.Name,
		)

		if err != nil {
			return models.GetListCategoryResponse{}, err
		}

		resp.Categorys= append(resp.Categorys, Category)
	}

	return resp, nil
}

func UpdateCategory(db *sql.DB, Category models.UpdateCategory) error {

	query := `
		UPDATE 
		Categorys 
		SET 
			name = $2
		WHERE id = $1
	`

	_, err := db.Exec(query,
		Category.Id,
		Category.Name,
	)

	if err != nil {
		return err
	}

	return nil
}

func DeleteCategory(db *sql.DB, req models.CategoryPrimeryKey)  error {
	_, err := db.Exec("DELETE FROM Categorys WHERE id = $1", req.Id)

	if err != nil {
		return err
	}

	return nil
}
