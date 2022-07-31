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
	h := handler.NewHandler(db, groupRepo, userRepo)
	h.Register(v1)
	port := fmt.Sprintf(":%s", config.Config("PORT"))
	fmt.Println(port)
	r.Logger.Fatal(r.Start(port))

}
