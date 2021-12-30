package handler

import (
	"net/http"
	"rms/model"

	"github.com/labstack/echo/v4"
)

// this function is responsible for listing all group tabls by calling stored procedure [GroupTablesList]
func (h *Handler) ItemsListByGroupAndMenu(c echo.Context) error {
	group := c.Param("group")

	var items []model.Item
	rows, err := h.db.Raw("EXEC StkMs01ListByMenuAndGroup 	@GroupCode = ?", group).Rows()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	defer rows.Close()
	for rows.Next() {
		var item model.Item
		err = rows.Scan(&item.ItemSerial, &item.ItemPrice, &item.ItemCode, &item.ItemName, &item.WithModifier)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		items = append(items, item)
	}

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
		err = rows.Scan(&item.ItemSerial, &item.ItemCode, &item.ItemName, &item.Screen)
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
