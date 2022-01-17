package handler

import (
	"net/http"
	"rms/model"

	"github.com/labstack/echo/v4"
)

func (h *Handler) CustomersListByName(c echo.Context) error {
	name := c.QueryParam("name")
	var customers []model.Customer
	rows, err := h.db.Raw("EXEC AccMs01ListByCodeNameType @Name = ? , @Type = 1", name).Rows()
	if err != nil {
		return c.JSON(http.StatusNoContent, err.Error())
	}
	defer rows.Close()
	for rows.Next() {
		var customer model.Customer
		err = rows.Scan(&customer.Serial, &customer.AccountCode, &customer.AccountName)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		customers = append(customers, customer)
	}
	return c.JSON(http.StatusOK, customers)
}
