package handler

import (
	"rms/repo"

	"github.com/jinzhu/gorm"
)

type Handler struct {
	db        *gorm.DB
	groupRepo repo.GroupRepo
	userRepo  repo.UserRepo
	itemRepo  repo.ItemRepo
	tableRepo repo.TableRepo
	menuRepo  repo.MenuRepo
	tax       float64
}

func NewHandler(databaase *gorm.DB, groupRepo repo.GroupRepo, userRepo repo.UserRepo, itemRepo repo.ItemRepo, tableRepo repo.TableRepo, menuRepo repo.MenuRepo) *Handler {
	return &Handler{
		groupRepo: groupRepo,
		userRepo:  userRepo,
		itemRepo:  itemRepo,
		tableRepo: tableRepo,
		menuRepo:  menuRepo,
		db:        databaase,
		tax:       0,
	}
}
