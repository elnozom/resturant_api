package handler

import "github.com/jinzhu/gorm"

type Handler struct {
	db  *gorm.DB
	tax float64
}

func NewHandler(databaase *gorm.DB) *Handler {
	return &Handler{
		db:  databaase,
		tax: 0,
	}
}
