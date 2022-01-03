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
type OrderItemsResp struct {
	OrderItemSerial int     `json:"orderItemSerial"`
	Qnt             int     `json:"qnt"`
	ItemPrice       float64 `json:"itemPrice"`
	ItemSerial      int     `json:"itemSerial"`
	ItemName        string  `json:"itemName"`
	IsMod           bool    `json:"isMod"`
}
