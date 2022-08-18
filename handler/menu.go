package handler

import (
	"net/http"
	"rms/model"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (h *Handler) MenuList(c echo.Context) error {
	resp, err := h.menuRepo.List()
	if err != nil {
		return c.JSON(http.StatusOK, resp)
	}
	return c.JSON(http.StatusOK, resp)
}
func (h *Handler) MenuListItems(c echo.Context) error {
	req := new(model.MenuItemsListReq)
	if err := c.Bind(req); err != nil {
		return err
	}
	resp, err := h.menuRepo.ListItems(req)
	if err != nil {
		return c.JSON(http.StatusOK, resp)
	}
	return c.JSON(http.StatusOK, resp)
}

// this function is responsible for listing all group tabls by calling stored procedure [GroupTablesList]
func (h *Handler) MenuFind(c echo.Context) error {
	group := c.Param("group")
	tableSerial := c.Param("tableSerial")

	var items []model.Item
	rows, err := h.db.Raw("EXEC StkMs01ListByMenuAndGroup 	@GroupCode = ? , @TableSerial = ?", group, tableSerial).Rows()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	defer rows.Close()
	for rows.Next() {
		var item model.Item
		err = rows.Scan(&item.ItemSerial, &item.ItemPrice, &item.ImagePath, &item.ItemCode, &item.ItemName, &item.ItemDesc, &item.WithModifier, &item.Qnt)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		items = append(items, item)
	}

	return c.JSON(http.StatusOK, items)
}

func (h *Handler) MenuEditAdd(c echo.Context) error {
	req := new(model.MenuInsertReq)
	if err := c.Bind(req); err != nil {
		return err
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "error scanning id : "+err.Error())
	}
	req.Id = id
	resp, err := h.menuRepo.EditAdd(req)
	if err != nil {
		return c.JSON(http.StatusOK, resp)
	}
	return c.JSON(http.StatusOK, resp)
}
