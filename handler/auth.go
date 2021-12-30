package handler

import (
	"net/http"
	"rms/model"

	"github.com/labstack/echo/v4"
)

func (h *Handler) EmpGetByCode(c echo.Context) error {
	code := c.Param("code")
	var employee model.Emp
	err := h.db.Raw("EXEC EmployeeGetByCode @EmpCode = ?", code).Row().Scan(&employee.EmpName, &employee.EmpPassword, &employee.EmpCode, &employee.SecLevel)
	if err != nil {
		return c.JSON(http.StatusNoContent, err.Error())
	}
	return c.JSON(http.StatusOK, employee)
}
func (h *Handler) EmpGetByBarCode(c echo.Context) error {
	code := c.Param("code")
	var employee model.Emp
	err := h.db.Raw("EXEC EmployeeGetByCode @BarCode = ?", code).Row().Scan(&employee.EmpName, &employee.EmpPassword, &employee.EmpCode, &employee.SecLevel)
	if err != nil {
		return c.JSON(http.StatusNoContent, err.Error())
	}
	return c.JSON(http.StatusOK, employee)
}

func (h *Handler) CheckDeviceAuthorization(c echo.Context) error {
	imei := c.Param("imei")
	var setting model.Setting
	err := h.db.Raw("EXEC ComUseGetDevice @Imei = ?", imei).Row().Scan(&setting.ComName, &setting.Capital, &setting.CashtraySerial)
	if err != nil {
		return c.JSON(http.StatusOK, nil)
	}

	return c.JSON(http.StatusOK, setting)
}

func (h *Handler) InsertDevice(c echo.Context) error {
	req := new(model.InsertDeviceReq)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	var resp int
	err := h.db.Raw("EXEC ComUseInsert @Imei = ? , @ComName = ?", req.Imei, req.ComName).Row().Scan(&resp)
	if err != nil {
		return c.JSON(http.StatusNoContent, err.Error())
	}
	return c.JSON(http.StatusOK, resp)
}
