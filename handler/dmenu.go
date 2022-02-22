package handler

import (
	"net/http"
	"rms/model"

	"github.com/labstack/echo/v4"
)

func (h *Handler) CreateCart(c echo.Context) error {
	req := new(model.CartCreateReq)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	var resp int
	err := h.db.Raw("EXEC CartCreate @CustomerSerial = ? , @DeviceId = ? , @TableSerial = ?  ", req.CustomerSerial, req.DeviceId, req.TableSerial).Row().Scan(&resp)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, resp)
}

// this function is responsible for listing cart item by a table
func (h *Handler) ListCartItems(c echo.Context) error {
	var items []model.CartItem
	rows, err := h.db.Raw("EXEC CartItemList @Serial = ?", c.QueryParam("Serial")).Rows()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	defer rows.Close()
	for rows.Next() {
		var item model.CartItem
		err = rows.Scan(
			&item.CartItemSerial,
			&item.Qnt,
			&item.Price,
			&item.ItemSerial,
			&item.ItemName,
			&item.IsMod,
			&item.MainModSerial,
			&item.AddItems,
		)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		items = append(items, item)
	}

	return c.JSON(http.StatusOK, items)
}

// this function is responsible for create cart item
func (h *Handler) CreateCartItem(c echo.Context) error {
	req := new(model.CartCreateItemReq)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	var resp bool
	err := h.db.Raw("EXEC CartItemCreate @CartSerial = ? , @ItemSerial = ? , @Price = ?", req.CartSerial, req.ItemSerial, req.Price).Row().Scan(&resp)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, resp)
}

// this function is responsible for creaet cart ietm
func (h *Handler) DeleteCartItem(c echo.Context) error {
	var resp bool
	err := h.db.Raw("EXEC CartItemDelete @Serial = ?", c.Param("serial")).Row().Scan(&resp)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, resp)
}
