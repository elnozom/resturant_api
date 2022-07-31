package handler

import (
	"rms/repo"

	"github.com/jinzhu/gorm"
)

type Handler struct {
	db        *gorm.DB
	groupRepo repo.GroupRepo
	userRepo  repo.UserRepo
	tax       float64
}

func NewHandler(databaase *gorm.DB, groupRepo repo.GroupRepo, userRepo repo.UserRepo) *Handler {
	return &Handler{
		groupRepo: groupRepo,
		userRepo:  userRepo,
		db:        databaase,
		tax:       0,
	}
}
