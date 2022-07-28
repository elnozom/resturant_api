package handler

import (
	"net/http"
	"rms/model"
	"strconv"

	"github.com/labstack/echo/v4"
)

// this function is responsible for listing all group tabls by calling stored procedure [GroupTablesList]
func (h *Handler) MainGroupsList(c echo.Context) error {
	var groups []model.MainGroup
	rows, err := h.db.Raw("EXEC GroupTypeList").Rows()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	defer rows.Close()
	for rows.Next() {
		var group model.MainGroup
		err = rows.Scan(&group.GroupCode, &group.GroupName, &group.Icon)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		groups = append(groups, group)
	}

	return c.JSON(http.StatusOK, groups)
}

// this function is responsible for listing all groups hierachy
func (h *Handler) GroupsListHierarchy(c echo.Context) error {
	result, err := h.groupRepo.ListHierarchy()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, result)
}

func (h *Handler) GroupsInsertUpdate(c echo.Context) error {
	codeFromUrl := c.Param("id")
	var codeInt int
	codeInt, _ = strconv.Atoi(codeFromUrl)
	req := new(model.GroupInsertUpdateReq)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	req.GroupCode = codeInt
	result, err := h.groupRepo.InsertUpdate(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, result)
}

// this function is responsible for listing all groups from groupcode table by calling GroupCodeListByGroupTypeId
// it uses (groupID) to select groups under this specific main group
func (h *Handler) SubGroupsByParent(c echo.Context) error {
	groupID := c.Param("group")
	var groups []model.SubGroup
	rows, err := h.db.Raw("EXEC GroupCodeListByGroupTypeId @GroupTypeID = ?", groupID).Rows()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	defer rows.Close()
	for rows.Next() {
		var group model.SubGroup
		err = rows.Scan(&group.GroupCode, &group.GroupName, &group.ImagePath)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		groups = append(groups, group)
	}

	return c.JSON(http.StatusOK, groups)
}
