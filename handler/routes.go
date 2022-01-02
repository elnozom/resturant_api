package handler

import (
	"github.com/labstack/echo/v4"
)

func (h *Handler) Register(v1 *echo.Group) {

	//global routes [ authinticate ]
	v1.GET("/employee/:code", h.EmpGetByCode)
	v1.GET("/employee//barcode/:code", h.EmpGetByBarCode)
	v1.GET("/authorize/:imei", h.CheckDeviceAuthorization)
	v1.POST("/authorize", h.InsertDevice)

	// tables routes
	tablesG := v1.Group("/tables")
	tablesG.GET("/groups", h.GroupTablesList)
	tablesG.GET("/:group", h.TablesListByGroupNo)
	tablesG.PUT("/open", h.TablesOpenOrder)
	tablesG.PUT("/close", h.TablesUnPause)

	// groups routes
	groupsG := v1.Group("/group")
	groupsG.GET("", h.MainGroupsList)
	groupsG.GET("/:group", h.SubGroupsByParent)

	// order routes
	ordersG := v1.Group("/order")
	ordersG.POST("", h.OrderInsert)
	ordersG.POST("/item", h.OrderItemInsert)
	ordersG.POST("/item/modifers", h.OrderItemInsertWithModifiers)
	ordersG.DELETE("/item/:serial", h.OrderItemDelete)

	// items routes
	itemsG := v1.Group("/item")
	itemsG.GET("/:group", h.ItemsListByGroupAndMenu)
	itemsG.GET("/modifiers/:serial", h.ItemsGetModifiersBySerial)

}
