package model

type Menu struct {
	Serial    int    `json:"Id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
}

type MenuItem struct {
	Serial         int    `json:"Id"`
	MenuSerial     int    `json:"menuId"`
	Name           string `json:"name"`
	GroupCode      string `json:"groupCode"`
	GroupName      string `json:"groupName"`
	BarCode        string `json:"barCode"`
	Price          string `json:"price"`
	MenuItemSerial int    `json:"MenuItemId"`
}

type MenuItemsListResp struct {
	Items     []MenuItem `json:"items"`
	MenuItems []MenuItem `json:"menuItems"`
}

type MenuItemsListReq struct {
	MenuId    int `query:"menuId"`
	GroupCode int `query:"groupCode"`
}

type MenuEditAddReq struct {
	Id   int
	Name string
}

type MenuAttachDetachReq struct {
	Id    int
	Items string `json:"items"`
}

type MenuPriceEditReq struct {
	Id    int
	Price float64
}
