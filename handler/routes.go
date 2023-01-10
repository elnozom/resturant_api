package handler

import (
	"github.com/labstack/echo/v4"
)

// comment
func (h *Handler) Register(v1 *echo.Group) {
	v1.GET("/validate", h.ValidateUser)
	v1.GET("/health", h.CheckHealth)
	v1.POST("/login", h.Login)
	v1.POST("/upload", h.Upload)
	v1.PUT("/image", h.UpdateImage)

	//global routes [ authinticate ]
	v1.GET("/employee/waiters", h.EmpListWaiters)
	v1.GET("/employee/:code", h.EmpGetByCode)
	v1.GET("/employee/barcode/:code", h.EmpGetByBarCode)
	v1.GET("/authorize/:imei", h.CheckDeviceAuthorization)
	v1.POST("/authorize", h.InsertDevice)
	v1.GET("/customers", h.CustomersListByName)
	v1.GET("/discounts", h.DiscountsListAll)
	v1.GET("/options", h.PosOptionsGet)

	// grouptables routes
	groupTablesG := v1.Group("/groupTable")
	groupTablesG.GET("", h.GroupTablesListBo)
	groupTablesG.POST("/editadd", h.GroupTablesEditAdd)
	groupTablesG.PUT("/editadd/:id", h.GroupTablesEditAdd)
	groupTablesG.GET("/:id", h.GroupTablesFind)

	// tables routes
	tablesG := v1.Group("/tables")
	tablesG.GET("/groups", h.GroupTablesList)
	tablesG.GET("/:group", h.TablesListByGroupNo)
	tablesG.PUT("/open", h.TablesOpenOrder)
	tablesG.PUT("/close", h.TablesUnPause)
	tablesG.PUT("/device/close/:imei", h.TablesCloseByImei)

	// groups routes
	groupG := v1.Group("/group")
	groupG.GET("", h.MainGroupsList)
	groupG.GET("/hierarchy", h.GroupsListHierarchy)
	groupG.POST("", h.GroupsEditAdd)
	groupG.PUT("/:id", h.GroupsEditAdd)
	groupG.GET("/:group", h.SubGroupsByParent)
	groupG.DELETE("/:id", h.GroupsDelete)

	// groups routes
	groups := v1.Group("/groups")
	groups.GET("/list", h.GroupCodeList)
	groups.PUT("/editadd/:id", h.GroupsEditAdd)
	groups.POST("/editadd", h.GroupsEditAdd)
	// groups.GET("", h.groupsListAll)
	groups.GET("/:id", h.GroupsFind)

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
	itemsG.GET("/barcodes/:group", h.ItemsListBarcodes)

	// cart routes
	cartG := v1.Group("/cart")
	cartG.GET("", h.ListCartItems)
	cartG.POST("", h.CreateCart)
	cartG.POST("/item", h.CreateCartItem)
	cartG.DELETE("/:serial", h.DeleteCartItem)
	cartG.POST("/call", h.CreateCartCall)
	cartG.GET("/call/:Imei", h.CheckCartCalls)
	cartG.GET("/call/:Imei/list", h.ListCartCalls)
	cartG.PUT("/call/respond", h.RespondCartCall)

	// guests rotues
	guestsG := v1.Group("/guests")
	guestsG.GET("/:device", h.GetGuestByDevice)
	guestsG.POST("", h.CreateGuest)

	// items backoffice routes
	// items routes
	items := v1.Group("/items")
	items.GET("/list", h.ItemsList)
	items.GET("/:id", h.ItemsFind)
	items.PUT("/editadd/:id", h.ItemsEditAdd)
	items.POST("/editadd", h.ItemsEditAdd)

	menus := v1.Group("/menus")
	menus.GET("/list", h.MenuList)
	menus.GET("/items", h.MenuListItems)
	menus.GET("/:id", h.MenuFind)
	menus.PUT("/attach/:id", h.MenuAttach)
	menus.PUT("/detach/:id", h.MenuDetach)
	menus.PUT("/price/edit", h.MenuPriceEdit)
	menus.POST("/editadd", h.MenuEditAdd)
	menus.PUT("/editadd/:id", h.MenuEditAdd)

	settings := v1.Group("/settings")
	settings.PUT("/:id", h.SettingsUpdate)

}
