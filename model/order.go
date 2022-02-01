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
	AddItems        string
}

type PrintItemResp struct {
	ItemName string
	Qnt      int
	Price    float64
	Total    float64
}
type PrintConfigResp struct {
	DocDate         string
	DocTime         string
	CashtryNo       int
	CustomerName    string
	OrderNo         int
	BonNo           int
	WaiterCode      int
	WaiterName      string
	TableNO         string
	GroupTableName  string
	GuestsNo        int
	DiscountPercent int
	DiscountValue   float64
	SaleTax         float64
	SubTotal        float64
	Total           float64
}
type PrintResp struct {
	Items  []PrintItemResp
	Config PrintConfigResp
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
type NoOfGuestsReq struct {
	HeadSerial int
	Guests     int
}

type TransferItemsReq struct {
	TableSerial  int
	ItemsSerials string
	Imei         string
	WaiterCode   int
	Split        bool
}
type AddonsReq struct {
	Serial int
	Addons string
}
