package handler

import (
	"net/http"
	"rms/model"

	"github.com/labstack/echo/v4"
)

func (h *Handler) ItemsListBarcodes(c echo.Context) error {
	group := c.Param("group")
	var items []string
	rows, err := h.db.Raw("SELECT BarCode FROM StkMs01 WHERE GroupCode = ? ", group).Rows()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	defer rows.Close()
	for rows.Next() {
		var item string
		err = rows.Scan(&item)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		items = append(items, item)
	}

	return c.JSON(http.StatusOK, items)
}

// this function is responsible for listing all group tabls by calling stored procedure [GroupTablesList]
func (h *Handler) ItemsListByGroupAndMenu(c echo.Context) error {
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

func (h *Handler) AddonsListAll(c echo.Context) error {
	rows, err := h.db.Raw("EXEC ISCodeListAll").Rows()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	var items []string
	for rows.Next() {
		var item string
		err = rows.Scan(&item)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		items = append(items, item)
	}
	defer rows.Close()
	return c.JSON(http.StatusOK, items)
}

// this function is responsible for listing all group tabls by calling stored procedure [GroupTablesList]
func (h *Handler) ItemsGetModifiersBySerial(c echo.Context) error {
	serial := c.Param("serial")

	var resp []map[int][]model.Item
	rows, err := h.db.Raw("EXEC StkMs01GetModifiersBySerial	@Serial = ?", serial).Rows()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	defer rows.Close()
	var screen int
	items := make(map[int][]model.Item)
	for rows.Next() {
		var item model.Item
		err = rows.Scan(&item.ItemSerial, &item.ItemCode, &item.ItemName, &item.Screen, &item.ScreenTimes)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		if item.Screen != screen {
			screen = item.Screen
		}
		items[screen] = append(items[screen], item)
	}
	resp = append(resp, items)

	// resp := handleModifiersResponse(items)
	return c.JSON(http.StatusOK, resp)
}

// func handleModifiersResponse(items []model.Item) [][]model.Item {
// 	var resp [][]model.Item

// 	for i, s := range items {
// 		fmt.Println(i, s)
// 	}

// 	return resp
// }
