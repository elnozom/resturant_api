package handler

import (
	"github.com/labstack/echo/v4"
)

//comment
func (h *Handler) Register(v1 *echo.Group) {

	//global routes [ authinticate ]
	v1.GET("/employee/waiters", h.EmpListWaiters)
	v1.GET("/employee/:code", h.EmpGetByCode)
	v1.GET("/employee/barcode/:code", h.EmpGetByBarCode)
	v1.GET("/authorize/:imei", h.CheckDeviceAuthorization)
	v1.POST("/authorize", h.InsertDevice)
	v1.GET("/customers", h.CustomersListByName)
	v1.GET("/discounts", h.DiscountsListAll)
	v1.GET("/options", h.PosOptionsGet)

	// tables routes
	tablesG := v1.Group("/tables")
	tablesG.GET("/groups", h.GroupTablesList)
	tablesG.GET("/:group", h.TablesListByGroupNo)
	tablesG.PUT("/open", h.TablesOpenOrder)
	tablesG.PUT("/close", h.TablesUnPause)
	tablesG.PUT("/device/close/:imei", h.TablesCloseByImei)

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
	ordersG.PUT("/addons", h.OrderUpdateAddons)

	// items routes
	itemsG := v1.Group("/item")
	itemsG.GET("/:group/:tableSerial", h.ItemsListByGroupAndMenu)
	itemsG.GET("/modifiers/:serial", h.ItemsGetModifiersBySerial)
	itemsG.GET("/addons", h.AddonsListAll)

	// cart routes
	cartG := v1.Group("/cart")
	cartG.GET("", h.ListCartItems)
	cartG.POST("", h.CreateCart)
	cartG.POST("/item", h.CreateCartItem)
	cartG.DELETE("/:serial", h.DeleteCartItem)
	cartG.POST("/call", h.CreateCartCall)

	// guests rotues
	v1.POST("/guests", h.CreateGuest)

}
