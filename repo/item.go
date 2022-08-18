package repo

import (
	"rms/model"

	"github.com/jinzhu/gorm"
)

type ItemRepo struct {
	db *gorm.DB
}

func NewItemRepo(db *gorm.DB) ItemRepo {
	return ItemRepo{
		db: db,
	}
}

func (ur *ItemRepo) List(req *model.ProductListReq) ([]model.ProductListResp, error) {
	var resp []model.ProductListResp
	rows, err := ur.db.Raw("EXEC ItemsListBo @serial = ? , @name = ? , @groupCode =? ,@priceFrom =? ,@priceTo =? ,@dateFrom =? ,@dateTo =?",
		req.Serial,
		req.ItemName,
		req.GroupCode,
		req.PriceFrom,
		req.PriceTo,
		req.DateFrom,
		req.DateTo,
	).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var rec model.ProductListResp
		err := rows.Scan(
			&rec.Id,
			&rec.ItemName,
			&rec.ItemNameEn,
			&rec.ImagePath,
			&rec.EstimatedTime,
			&rec.GroupCode,
			&rec.Category,
			&rec.BarCode,
			&rec.Price,
			&rec.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		resp = append(resp, rec)
	}
	return resp, nil
}

func (ur *ItemRepo) EditAdd(req *model.ProductEditAddReq) (*int, error) {
	var resp int
	err := ur.db.Raw("EXEC ItemsEditAdd  @serial = ?, @name = ?, @nameEn = ?,@imagePath = ? , @groupCode = ?, @bardCode = ?, @price = ?, @estimatedTime = ?",
		req.Serial,
		req.Name,
		req.NameEn,
		req.ImagePath,
		req.GroupCode,
		req.BarCode,
		req.Price,
		req.EstimatedTime,
	).Row().Scan(&resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
