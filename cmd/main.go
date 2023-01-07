package main

import (
	"add/pkg/db"
	"log"
	"github.com/gin-gonic/gin"
	"add/config"
	"add/api"
)

func main()  {
	cfg := config.Load() 
	db,err:= db.NewConnectPostgres(cfg)
	if err!=nil {
		log.Fatal("Fatal:", err.Error() )
		
	}

	r:= gin.New()

	r.Use(gin.Logger(),gin.Recovery())
	api.NewApi(r, db )
	err = r.Run(cfg.HTTPPort)
	if err!= nil {
		panic(err)
	}
}