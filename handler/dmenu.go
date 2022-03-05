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
	rows, err := h.db.Raw("EXEC CartItemList @Table = ? , @DeviceId = ?", c.QueryParam("Table"), c.QueryParam("Device")).Rows()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	defer rows.Close()
	for rows.Next() {
		var item model.CartItem
		err = rows.Scan(
			&item.CartSerial,
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
	var resp int
	err := h.db.Raw("EXEC CartItemCreate @CartSerial = ? , @ItemSerial = ? , @Price = ?", req.CartSerial, req.ItemSerial, req.Price).Row().Scan(&resp)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, resp)
}

// this function will be called when user click on call waiter or call cheque from dmenu
func (h *Handler) CreateCartCall(c echo.Context) error {
	req := new(model.CartCreateCallReq)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	var resp int
	err := h.db.Raw("EXEC CartCallCreate @CallType = ? , @CartSerial = ? , @TableSerial = ? ", req.Type, req.CartSerial, req.TableSerial).Row().Scan(&resp)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, resp)
}

// this function will be called when waiter respond to cart call waiter
func (h *Handler) RespondCartCall(c echo.Context) error {
	var resp bool
	err := h.db.Raw("EXEC CartCallRespond @Serial = ? ", c.Param("Serial")).Row().Scan(&resp)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, resp)
}

// this function will be called every intercval from waiters tablets to check if there is call waiter or cheque request
func (h *Handler) CheckCartCalls(c echo.Context) error {
	var items []model.CartCall
	rows, err := h.db.Raw("EXEC CartCheckCalls @Imei = ? ", c.Param("Imei")).Rows()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	defer rows.Close()
	for rows.Next() {
		var item model.CartCall
		err = rows.Scan(&item.CartSerial, &item.TableSerial, &item.Type, &item.GroupTableNo, &item.TableNo, &item.GuestName, &item.CreatedAt, &item.RespondedAt)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		items = append(items, item)
	}

	return c.JSON(http.StatusOK, items)
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

// this function is responsible for creaet cart ietm
func (h *Handler) CreateGuest(c echo.Context) error {
	var resp int
	req := new(model.GuestCreateReq)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	err := h.db.Raw("EXEC GuestsCreate @DeviceId = ? , @GuestName = ? , @GuestPhone = ?  ", req.DeviceId, req.GuestName, req.GuestPhone).Row().Scan(&resp)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, resp)
}

// this function is responsible for selcting guest name and phone by passing device id
func (h *Handler) GetGuestByDevice(c echo.Context) error {
	var resp model.GuestCreateReq
	device := c.Param("device")
	err := h.db.Raw("EXEC GuestsGetByDeivce @DeviceId = ? ", device).Row().Scan(&resp.GuestName, &resp.GuestPhone)
	resp.DeviceId = device
	if err != nil && err.Error() != "sql: no rows in result set" {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, resp)
}
