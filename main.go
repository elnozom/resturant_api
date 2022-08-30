package main

import (
	"fmt"
	"rms/config"
	"rms/db"
	"rms/handler"
	"rms/repo"
	"rms/router"
)

func main() {
	r := router.New()
	v1 := r.Group("/api")
	db.InitDatabase()
	db := db.DBConn
	groupRepo := repo.NewGroupRepo(db)
	userRepo := repo.NewUserRepo(db)
	itemRepo := repo.NewItemRepo(db)
	tableRepo := repo.NewTableRepo(db)
	menuRepo := repo.NewMenuRepo(db)
	h := handler.NewHandler(db, groupRepo, userRepo, itemRepo, tableRepo, menuRepo)
	h.Register(v1)
	port := fmt.Sprintf(":%s", config.Config("PORT"))
	r.Logger.Fatal(r.Start(port))

}
