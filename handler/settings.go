package handler

import (
	"net/http"
	"rms/model"

	"github.com/labstack/echo/v4"
)

// this function is responsible for listing all group tables by calling stored procedure [GroupTablesList]
func (h *Handler) PosOptionsGet(c echo.Context) error {
	var resp model.Option
	err := h.db.Raw("EXEC POSOptionsGet").Row().Scan(
		&resp.StoreCode,
		&resp.MultiPOS,
		&resp.TransSerial,
		&resp.UseWaiter,
		&resp.AccountSerialint,
		&resp.SaleTax,
	)
	h.tax = resp.SaleTax
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, resp)
}
