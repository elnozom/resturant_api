package model

type Order struct {
	TableSerial int
	TableNo     int
	Imei        string
	OrderType   int
	WaiterCode  int
}

type OrderResp struct {
	DocDate    string
	DocNo      string
	WaiterCode int
}
type OrderCreateResp struct {
	HeadSerial int
	DocNo      string
}
type OrderItemsResp struct {
	OrderItemSerial int
	MainModSerial   int
	Qnt             int
	ItemPrice       float64
	ItemSerial      int
	WithModifier    bool
	ItemName        string
	IsMod           bool
	Screen          int
}
type Discount struct {
	DiscCode  int
	DiscDesc  string
	DiscValue float64
	DelTax    bool
}
type ApplyDiscountReq struct {
	HeadSerial int
	Comment    string
	DiscCode   int
	DiscValue  float64
}
