package handler

import (
	"net/http"
	"rms/model"

	"github.com/labstack/echo/v4"
)

func (h *Handler) OrderInsert(c echo.Context) error {
	req := new(model.Order)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	var resp int
	err := h.db.Raw("EXEC Stktr03Insert	@TableSerial = ? ,@Imei = ? ,@OrderType = ? ,@WaiterCode = ? ", req.TableSerial, req.Imei, req.OrderType, req.WaiterCode).Row().Scan(&resp)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, resp)
}

func (h *Handler) OrderItemInsert(c echo.Context) error {
	req := new(model.InsertItemReq)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	var resp int
	err := h.db.Raw("EXEC StkTr04Insert	@HeadSerial = ? ,@ItemSerial = ? ,@WithMod = ? ,@IsMod = ?", req.HeadSerial, req.ItemSerial, req.WithMod, req.IsMod).Row().Scan(&resp)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, resp)
}

func (h *Handler) OrderItemDelete(c echo.Context) error {
	var resp bool
	err := h.db.Raw("EXEC StkTr04Delete	@Serial = ?", c.Param("serial")).Row().Scan(&resp)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, resp)
}
