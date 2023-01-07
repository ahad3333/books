package api

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"add/api/handler"
)

func NewApi(r *gin.Engine, db *sql.DB) {

	handlerV1 := handler.NewHandler(db)

	r.POST("/book", handlerV1.CreateBook)
	r.GET("/book/:id", handlerV1.GetByIDBook)
	r.GET("/book", handlerV1.GetListBook)
	r.PUT("/book/:id",handlerV1.UpdateBook)
	r.DELETE("/book/:id",handlerV1.DeleteBook)

	r.POST("/category", handlerV1.CreateCategory)
	r.GET("/category/:id", handlerV1.GetByIdCategory)
	r.GET("/category", handlerV1.GetListCategory)
	r.PUT("/category/:id",handlerV1.UpdateCategory)
	r.DELETE("/category/:id",handlerV1.DeleteCategory)


}

