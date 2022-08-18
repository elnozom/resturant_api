package repo

import (
	"rms/model"

	"github.com/jinzhu/gorm"
)

type TableRepo struct {
	db *gorm.DB
}

func NewTableRepo(db *gorm.DB) TableRepo {
	return TableRepo{
		db: db,
	}
}

func (ur *TableRepo) List(id int) ([]model.TableGroupBo, error) {
	var groups []model.TableGroupBo
	rows, err := ur.db.Raw("EXEC GroupTablesList @id = ?", id).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var group model.TableGroupBo
		err = rows.Scan(&group.GroupTableNo, &group.GroupTableName, &group.StartNo, &group.TableCount)
		if err != nil {
			return nil, err
		}
		groups = append(groups, group)
	}

	return groups, nil

}
