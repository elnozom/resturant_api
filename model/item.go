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

type ProductEditAddReq struct {
	ItemCode       string
	GroupCode      string `json:"groupCode"`
	SupplierCode   string
	SupplierName   string
	BarCode        string `json:"barCode"`
	Name           string `json:"name"`
	MinorPerMajor  string
	AccountSerial  string
	ActiveItem     bool
	ItemTypeID     string
	ItemHaveSerial bool
	MasterItem     bool
	StoreCode      string
	LastBuyPrice   float64
	POSTP          float64
	POSPP          string `json:"price"`
	Ratio1         float64
	Ratio2         float64
	Percent1       float64
	Percen2        float64
	Disc1          float64
	Disc2          float64
	PriceBefore    float64
	Tax1           float64
}
