package repo

import (
	"fmt"
	"rms/model"
	"rms/utils"
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

func (ur *GroupRepo) Find(id *int) (*model.GroupHierarchy, error) {
	var resp model.GroupHierarchy
	err := ur.db.Raw("EXEC GroupCodeFind @groupCode = ?", id).Row().Scan(
		&resp.GroupCode,
		&resp.GroupName,
		&resp.GroupNameEn,
		&resp.Parent,
		&resp.Code,
	)
	if utils.CheckErr(&err) {
		return nil, err
	}
	return &resp, nil
}

func (ur *GroupRepo) List(lang *string, parent *string) (*[]model.GroupListResp, error) {
	var resp []model.GroupListResp
	rows, err := ur.db.Raw("EXEC GroupCodeList @parentCode = ?", parent).Rows()
	if utils.CheckErr(&err) {
		return nil, err
	}
	for rows.Next() {
		var rec model.GroupListResp
		err := rows.Scan(&rec.GroupCode, &rec.GroupName, &rec.GroupNameEn, &rec.Code)
		if utils.CheckErr(&err) {
			return nil, err
		}
		if *lang == "ar" {
			rec.Name = fmt.Sprintf("%s %s", rec.GroupName, rec.Code)
		} else {
			rec.Name = fmt.Sprintf("%s %s", rec.GroupNameEn, rec.Code)
		}

		resp = append(resp, rec)
	}
	return &resp, nil
}

func (ur *GroupRepo) InsertUpdate(req *model.GroupInsertUpdateReq) (*model.GroupHierarchy, error) {
	var resp model.GroupHierarchy
	err := ur.db.Raw("EXEC GroupCodeInsertUpdate @groupCode = ? , @imagePath = ? , @groupName = ? , @groupNameEn = ? , @parentCode = ?",
		&req.GroupCode,
		&req.ImagePath,
		&req.GroupName,
		&req.GroupNameEn,
		&req.ParentCode,
	).Row().Scan(
		&resp.GroupCode,
		&resp.Name,
		&resp.GroupNameEn,
		&resp.Parent,
		&resp.Code,
	)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (ur *GroupRepo) ListHierarchy(lang *string) (*[]model.GroupHierarchy, error) {
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
			&rec.GroupName,
			&rec.GroupNameEn,
			&rec.Parent,
			&rec.Code,
			&rec.ChildrenLength,
			&rec.Level,
		)

		if *lang == "ar" {
			rec.Name = fmt.Sprintf("%s %s", rec.GroupName, rec.Code)
		} else {
			rec.Name = fmt.Sprintf("%s %s", rec.GroupNameEn, rec.Code)
		}
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
