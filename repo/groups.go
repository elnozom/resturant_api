package repo

import (
	"rms/model"
	"strconv"

	"github.com/jinzhu/gorm"
)

type GroupRepo struct {
	db *gorm.DB
}

func NewGroupRepo(db *gorm.DB) GroupRepo {
	return GroupRepo{
		db: db,
	}
}

func (ur *GroupRepo) InsertUpdate(req *model.GroupInsertUpdateReq) (*model.GroupHierarchy, error) {
	var resp model.GroupHierarchy
	err := ur.db.Raw("EXEC GroupCodeInsertUpdate @groupCode = ? , @groupName = ? , @groupNameEn = ? , @parentCode = ? , @code = ?",
		&req.GroupCode,
		&req.GroupName,
		&req.GroupNameEn,
		&req.ParentCode,
		&req.Code,
	).Row().Scan(
		&resp.Name,
		&resp.NameEn,
		&resp.Parent,
		&resp.Code,
	)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (ur *GroupRepo) ListHierarchy() (*[]model.GroupHierarchy, error) {
	var resp []model.GroupHierarchy
	rows, err := ur.db.Raw("EXEC GroupCodeListHierarchy").Rows()
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var currentGrandParent *model.GroupHierarchy
	var currentParent *model.GroupHierarchy
	for rows.Next() {
		var rec model.GroupHierarchy
		err := rows.Scan(
			&rec.GroupCode,
			&rec.Name,
			&rec.NameEn,
			&rec.Parent,
			&rec.Code,
			&rec.ChildrenLength,
			&rec.Level,
		)
		rec.Children = make([]model.GroupHierarchy, 0)
		if err != nil {
			return nil, err
		}

		if rec.Level == 0 {
			resp = append(resp, rec)
		}
		if rec.Level == 1 {
			if currentGrandParent == nil || currentGrandParent.Code != rec.Parent {
				intParent, _ := strconv.Atoi(rec.Parent)
				currentGrandParent = &resp[intParent-1]
			}
			currentGrandParent.Children = append(currentGrandParent.Children, rec)
		}
		if rec.Level == 2 {
			grandParentCode := rec.Parent[0:2]
			if currentGrandParent == nil || currentGrandParent.Code != grandParentCode {
				intParent, _ := strconv.Atoi(grandParentCode)
				currentGrandParent = &resp[intParent-1]
				// currentGrandParent = _findGrandParent(&grandParentCode, &resp)
			}
			parentCode := rec.Parent[0:4]
			if currentParent == nil || currentParent.Code != parentCode {
				currentParent = _findGrandParent(&parentCode, &currentGrandParent.Children)
			}
			currentParent.Children = append(currentParent.Children, rec)
		}

	}
	return &resp, nil
}

func _findGrandParent(code *string, list *[]model.GroupHierarchy) *model.GroupHierarchy {
	for i := 0; i < len(*list); i++ {
		if (*list)[i].Code == *code {
			return &(*list)[i]
		}
	}
	return nil
}
