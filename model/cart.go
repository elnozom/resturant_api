package model

type CartItem struct {
	CartSerial     int
	CartItemSerial int
	Qnt            int
	Price          float64
	ItemSerial     int
	ItemName       string
	IsMod          bool
	MainModSerial  int
	AddItems       string
}
type CartCreateCallReq struct {
	Type        int
	CartSerial  int
	TableSerial int
	GuestName   string
}

type CartCallRespondReq struct {
	Serials    string
	WaiterCode int
}
type CartCall struct {
	Type           int
	CartSerial     int
	TableSerial    int
	GroupTableName string
	GuestName      string
	CreatedAt      string
	RespondedAt    string
	GroupTableNo   int
	TableNo        int
}

type CartCreateReq struct {
	CustomerSerial int
	DeviceId       string
	TableSerial    int
}

type CartCreateItemReq struct {
	ItemSerial int
	CartSerial int
	Price      float64
}

type Cart struct {
	CartSerial     int
	Amount         float64
	CustomerSerial int
	DeviceId       string
}
