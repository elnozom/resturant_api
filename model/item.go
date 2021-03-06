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
