package model

type Item struct {
	ItemSerial      int
	ItemPrice       float32
	ItemCode        int
	ItemName        string
	ItemDesc        string
	ImagePath       string
	WithModifier    bool
	Screen          int
	ScreenTimes     int
	OrderItemSerial int
	Qnt             float32
	MainModSerial   int
	AddItems        string
	Printed         bool
}

type InsertItemReq struct {
	HeadSerial int
	ItemSerial int
	WithMod    bool
	IsMod      bool
	Qnt        int
}

type OrderChangeTableReq struct {
	NewSerial    int
	OldSerial    int
	ComputerName string
}

type OrderChangeCustomerReq struct {
	CustomerSerial int
	HeadSerial     int
}

type OrderChangeWaiterReq struct {
	WaiterCode int
	HeadSerial int
}

type InsertItemWithModifiersReq struct {
	ItemsSerials    string
	HeadSerial      int
	OrderItemSerial int
}

type ProductListReq struct {
	Serial    int
	ItemName  string  `query:"name"`
	GroupCode string  `query:"groupCode"`
	PriceFrom float64 `query:"priceFrom"`
	PriceTo   float64 `query:"priceTo"`
	DateFrom  string  `query:"dateFrom"`
	DateTo    string  `query:"dateTo"`
}

type ProductListResp struct {
	Id            int     `json:"Id"`
	ItemName      string  `json:"name"`
	ItemNameEn    string  `json:"name_en"`
	ImagePath     string  `json:"image"`
	Category      string  `json:"category"`
	GroupCode     int     `json:"groupCode"`
	BarCode       string  `json:"barCode"`
	Price         float64 `json:"price"`
	EstimatedTime int     `json:"estimated_time"`
	CreatedAt     string  `json:"created_at"`
}

type ProductEditAddReq struct {
	Serial        int     `json:"serial"`
	Name          string  `json:"name"`
	ImagePath     string  `json:"image"`
	NameEn        string  `json:"name_en"`
	GroupCode     int     `json:"groupCode"`
	BarCode       string  `json:"barCode"`
	Price         float64 `json:"price"`
	EstimatedTime int     `json:"estimated_time"`
}
