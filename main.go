package main

import (
	"fmt"
	"rms/config"
	"rms/db"
	"rms/handler"
	"rms/router"
)

func main() {
	r := router.New()
	v1 := r.Group("/api")
	db.InitDatabase()
	db := db.DBConn
	h := handler.NewHandler(db)
	h.Register(v1)
	port := fmt.Sprintf(":%s", config.Config("PORT"))
	fmt.Println(port)
	r.Logger.Fatal(r.Start(port))

}
