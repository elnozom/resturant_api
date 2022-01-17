package handler

import (
	"database/sql"
	"fmt"
	"net/http"
	"rms/model"

	"github.com/labstack/echo/v4"
)

func (h *Handler) EmpListWaiters(c echo.Context) error {
	var employees []model.Emp
	rows, err := h.db.Raw("EXEC EmployeeGetByCode @IsWaiter = 1").Rows()
	defer rows.Close()
	if err != nil {
		return c.JSON(http.StatusNoContent, err.Error())
	}
	for rows.Next() {
		employee := scanEmpResult(rows)
		employees = append(employees, employee)
	}
	return c.JSON(http.StatusOK, employees)
}
func scanEmpResult(row *sql.Rows) model.Emp {
	var employee model.Emp
	err := row.Scan(&employee.EmpName, &employee.EmpPassword, &employee.EmpCode, &employee.SecLevel)
	if err != nil {
		fmt.Println(err.Error())
	}
	return employee
}
