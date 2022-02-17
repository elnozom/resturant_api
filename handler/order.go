package handler

import (
	"fmt"
	"net/http"
	"rms/model"
	"strings"

	"github.com/labstack/echo/v4"
)

func (h *Handler) OrderInsert(c echo.Context) error {
	req := new(model.Order)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	var resp model.OrderCreateResp
	err := h.db.Raw("EXEC Stktr03Insert	@TableSerial = ? ,@Imei = ? ,@OrderType = ? ,@WaiterCode = ? ", req.TableSerial, req.Imei, req.OrderType, req.WaiterCode).Row().Scan(&resp.HeadSerial, &resp.DocNo)
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

func (h *Handler) OrderChangeTable(c echo.Context) error {
	req := new(model.OrderChangeTableReq)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	var resp bool
	err := h.db.Raw("EXEC StkTr03ChangeTable  @NewTableSerial = ? , @OldTableSerial = ?", req.NewSerial, req.OldSerial).Row().Scan(&resp)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, resp)
}

func (h *Handler) OrderChangeCustomer(c echo.Context) error {
	req := new(model.OrderChangeCustomerReq)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	var resp bool
	err := h.db.Raw("EXEC StkTr03ChangeCustomer  @HeadSerial = ? , @CustomerSerial = ?", req.HeadSerial, req.CustomerSerial).Row().Scan(&resp)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, resp)
}

func (h *Handler) OrderChangeWaiter(c echo.Context) error {
	req := new(model.OrderChangeWaiterReq)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	var resp bool
	err := h.db.Raw("EXEC StkTr03ChangeWaiter  @HeadSerial = ? , @WaiterCode = ?", req.HeadSerial, req.WaiterCode).Row().Scan(&resp)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, resp)
}

func (h *Handler) OrderApplyDiscount(c echo.Context) error {
	req := new(model.ApplyDiscountReq)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	fmt.Println(req)
	var resp bool
	err := h.db.Raw("EXEC StkTr03ApplyDiscount  @HeadSerial = ?, @DiscCode = ? , @DiscPercent = ? , @Comment = ?  ", req.HeadSerial, req.DiscCode, req.DiscValue, req.Comment).Row().Scan(&resp)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, resp)
}

func (h *Handler) OrderSetNoOfGuests(c echo.Context) error {
	req := new(model.NoOfGuestsReq)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	fmt.Println(req)
	var resp bool
	err := h.db.Raw("EXEC StkTr03SetNoOfGuests  @HeadSerial = ?, @Guests = ?   ", req.HeadSerial, req.Guests).Row().Scan(&resp)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, resp)
}

func (h *Handler) OrderUpdateAddons(c echo.Context) error {
	req := new(model.AddonsReq)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	_ = h.db.Raw("EXEC StkTr04ApplyAddons  @Serial = ?, @Addons = ? ", req.Serial, req.Addons).Row()

	return c.JSON(http.StatusOK, true)
}
func (h *Handler) OrderTransferItems(c echo.Context) error {
	req := new(model.TransferItemsReq)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	fmt.Println(req)
	_ = h.db.Raw("EXEC Stktr04TransferItems  @TableSerial = ?, @ItemsSerials = ?  , @Imei = ?  , @WaiterCode = ?  , @Split = ? ", req.TableSerial, req.ItemsSerials, req.Imei, req.WaiterCode, req.Split).Row()

	return c.JSON(http.StatusOK, true)
}

func (h *Handler) OrderItemInsertWithModifiers(c echo.Context) error {
	req := new(model.InsertItemWithModifiersReq)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	var resp bool
	err := h.db.Raw("EXEC Stktr04InsertModifiers @ItemsSerials = ? , @HeadSerial = ? , @OrderItemSerial = ?", req.ItemsSerials, req.HeadSerial, req.OrderItemSerial).Row().Scan(&resp)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, resp)
}

func (h *Handler) OrderListItemsBySerial(c echo.Context) error {
	req := new(model.InsertItemWithModifiersReq)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	var items []model.OrderItemsResp
	rows, err := h.db.Raw("EXEC StkTr03ListItemsBySerial @Serial = ? ", c.Param("serial")).Rows()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	defer rows.Close()
	for rows.Next() {
		var item model.OrderItemsResp
		err = rows.Scan(&item.OrderItemSerial, &item.Qnt, &item.ItemPrice, &item.ItemSerial, &item.ItemName, &item.IsMod, &item.MainModSerial, &item.AddItems, &item.Printed)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		items = append(items, item)
	}
	return c.JSON(http.StatusOK, items)
}
func (h *Handler) OrderListItemsForPrint(c echo.Context) error {
	req := new(model.InsertItemWithModifiersReq)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	var resp model.PrintResp
	rows, err := h.db.Raw("EXEC StkTr03PrintItemsBySerial @Serial = ? ", c.Param("serial")).Rows()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	defer rows.Close()
	for rows.Next() {
		var item model.PrintItemResp
		err = rows.Scan(&resp.Config.DocDate, &resp.Config.DocTime, &resp.Config.CashtryNo, &resp.Config.CustomerName, &resp.Config.OrderNo, &resp.Config.BonNo, &item.ItemName, &resp.Config.WaiterName, &resp.Config.TableNO, &resp.Config.GroupTableName, &resp.Config.GuestsNo, &resp.Config.DiscountPercent, &resp.Config.WaiterCode, &resp.Config.SaleTax, &item.Qnt, &item.Price, &item.Total)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		resp.Config.SubTotal += item.Total
		resp.Items = append(resp.Items, item)
	}

	resp.Config.DocDate = strings.Split(resp.Config.DocDate, "T")[0]
	resp.Config.DocTime = strings.Split(resp.Config.DocTime, "T")[1]
	resp.Config.DocTime = resp.Config.DocTime[0:5]
	resp.Config.DiscountValue = float64(resp.Config.DiscountPercent) * resp.Config.SubTotal / 100
	resp.Config.SaleTax = (h.tax * resp.Config.SubTotal) / 100
	resp.Config.TaxPercent = h.tax
	resp.Config.Total = (resp.Config.SubTotal + resp.Config.SaleTax) - resp.Config.DiscountValue
	return c.JSON(http.StatusOK, resp)
}

func (h *Handler) DiscountsListAll(c echo.Context) error {
	var codes []model.Discount
	rows, err := h.db.Raw("EXEC DisCodesListAll  ").Rows()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	defer rows.Close()
	for rows.Next() {
		var code model.Discount
		err = rows.Scan(&code.DiscCode, &code.DiscDesc, &code.DiscValue, &code.DelTax, &code.SecLevel)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		codes = append(codes, code)
	}
	return c.JSON(http.StatusOK, codes)
}
