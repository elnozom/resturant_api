package handler

import (
	"github.com/labstack/echo/v4"
)

func (h *Handler) Register(v1 *echo.Group) {

	//global routes [ authinticate ]
	v1.GET("/employee/waiters", h.EmpListWaiters)
	v1.GET("/employee/:code", h.EmpGetByCode)
	v1.GET("/employee/barcode/:code", h.EmpGetByBarCode)
	v1.GET("/authorize/:imei", h.CheckDeviceAuthorization)
	v1.POST("/authorize", h.InsertDevice)
	v1.GET("/customers", h.CustomersListByName)
	v1.GET("/discounts", h.DiscountsListAll)

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
	ordersG.GET("/:serial", h.OrderListItemsBySerial)
	ordersG.GET("/print/:serial", h.OrderListItemsForPrint)
	ordersG.POST("", h.OrderInsert)
	ordersG.POST("/item", h.OrderItemInsert)
	ordersG.POST("/item/modifers", h.OrderItemInsertWithModifiers)
	ordersG.DELETE("/item/:serial", h.OrderItemDelete)
	// update routes
	ordersG.PUT("/table", h.OrderChangeTable)
	ordersG.PUT("/customer", h.OrderChangeCustomer)
	ordersG.PUT("/waiter", h.OrderChangeWaiter)
	ordersG.PUT("/discount", h.OrderApplyDiscount)
	ordersG.PUT("/guests", h.OrderSetNoOfGuests)
	ordersG.PUT("/transfer", h.OrderTransferItems)

	// items routes
	itemsG := v1.Group("/item")
	itemsG.GET("/:group/:tableSerial", h.ItemsListByGroupAndMenu)
	itemsG.GET("/modifiers/:serial", h.ItemsGetModifiersBySerial)

}
