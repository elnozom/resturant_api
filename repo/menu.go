package repo

import (
	"rms/model"

	"github.com/jinzhu/gorm"
)

type MenuRepo struct {
	db *gorm.DB
}

func NewMenuRepo(db *gorm.DB) MenuRepo {
	return MenuRepo{
		db: db,
	}
}

func (ur *MenuRepo) List() ([]model.Menu, error) {
	var resp []model.Menu
	rows, err := ur.db.Raw("EXEC MenusList").Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var rec model.Menu
		err := rows.Scan(
			&rec.Serial,
			&rec.Name,
			&rec.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		resp = append(resp, rec)
	}
	return resp, nil
}

func (ur *MenuRepo) ListItems(req *model.MenuItemsListReq) (*model.MenuItemsListResp, error) {
	var inMenu []model.MenuItem
	var outMenu []model.MenuItem
	rows, err := ur.db.Raw("EXEC MenuItemsList @menuId = ? , @groupCode = ?", req.MenuId, req.GroupCode).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var rec model.MenuItem
		err := rows.Scan(
			&rec.Serial,
			&rec.Name,
			&rec.GroupCode,
			&rec.GroupName,
			&rec.BarCode,
			&rec.Price,
		)
		if err != nil {
			return nil, err
		}
		inMenu = append(inMenu, rec)
	}
	if rows.NextResultSet() {
		for rows.Next() {
			var rec model.MenuItem
			err := rows.Scan(
				&rec.Serial,
				&rec.Name,
				&rec.GroupCode,
				&rec.GroupName,
				&rec.BarCode,
				&rec.Price,
			)
			if err != nil {
				return nil, err
			}
			outMenu = append(outMenu, rec)
		}
	}

	resp := model.MenuItemsListResp{Items: outMenu, MenuItems: inMenu}
	return &resp, nil
}

func (ur *MenuRepo) EditAdd(req *model.MenuInsertReq) (*int, error) {
	var resp int
	err := ur.db.Raw("EXEC MenusEditAdd  @id = ? , @name = ?, @items = ?",
		req.Id,
		req.Name,
		req.Items,
	).Row().Scan(&resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
