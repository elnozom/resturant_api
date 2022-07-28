package handler

import (
	"rms/repo"

	"github.com/jinzhu/gorm"
)

type Handler struct {
	db        *gorm.DB
	groupRepo repo.GroupRepo
	tax       float64
}

func NewHandler(databaase *gorm.DB, groupRepo repo.GroupRepo) *Handler {
	return &Handler{
		groupRepo: groupRepo,
		db:        databaase,
		tax:       0,
	}
}
