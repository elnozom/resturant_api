package handler

import (
	"fmt"
	"net/http"
	"rms/model"

	"github.com/labstack/echo/v4"
)

// this function is responsible for listing all group tables by calling stored procedure [GroupTablesList]
func (h *Handler) GroupTablesList(c echo.Context) error {
	var groups []model.TableGroup
	rows, err := h.db.Raw("EXEC GroupTablesList").Rows()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	defer rows.Close()
	for rows.Next() {
		var group model.TableGroup
		err = rows.Scan(&group.GroupTableNo, &group.GroupTableName, &group.TableCount)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		groups = append(groups, group)
	}

	return c.JSON(http.StatusOK, groups)
}

// this function is responsible for listing all tables for desired group
// it uses (GroupTableNo) to select tables under this specific group
// by calling stored procedure [TablesListByGroupNo]
func (h *Handler) TablesListByGroupNo(c echo.Context) error {
	groupNo := c.Param("group")
	var tables []model.Table
	rows, err := h.db.Raw("EXEC TablesListByGroupNo @GroupTableNo = ?", groupNo).Rows()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	defer rows.Close()
	for rows.Next() {
		var table model.Table
		err = rows.Scan(&table.Serial, &table.TableNo, &table.TableName, &table.Pause, &table.State, &table.PrintTimes, &table.OpenDateTime)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		if table.State == "Free" {
			table.Status = table.State
		} else {
			if table.PrintTimes > 0 {
				// table.Status = table.State + "with_chque" + "printed : " + table.PrintTimes
				table.Status = fmt.Sprintf("%s_with_cheque", table.State)
			} else {
				table.Status = fmt.Sprintf("%s_without_cheque", table.State)
			}
		}
		tables = append(tables, table)
	}

	return c.JSON(http.StatusOK, tables)
}

// this function is responsible for opening new order on a specific table
// it will return true if order was succesfully opened
// and false if the table was already pause
func (h *Handler) TablesOpenOrder(c echo.Context) error {
	req := new(model.TablesOpenOrderReq)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	var resp model.TablesOpenOrderResp
	err := h.db.Raw("EXEC TablesOpenOrder @Serial = ? , @EmpCode = ? , @Imei = ?", req.Serial, req.EmpCode, req.Imei).Row().Scan(&resp.IsOrderOpened, &resp.Msg)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, resp)
}

// this function is responsible for opening new order on a specific table
// it will return true if order was succesfully opened
// and false if the table was already pause
func (h *Handler) TablesUnPause(c echo.Context) error {
	req := new(model.TablesOpenOrderReq)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	var IsTableClosed bool
	err := h.db.Raw("EXEC TablesUnPause @Serial = ?  , @Imei = ?", req.Serial, req.Imei).Row().Scan(&IsTableClosed)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, IsTableClosed)
}
