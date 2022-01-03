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
