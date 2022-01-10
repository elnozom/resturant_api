package model

type Item struct {
	ItemSerial      int
	ItemPrice       float32
	ItemCode        int
	ItemName        string
	WithModifier    bool
	Screen          int
	OrderItemSerial int
	MainModSerial   int
}

type InsertItemReq struct {
	HeadSerial int
	ItemSerial int
	WithMod    bool
	IsMod      bool
	Qnt        int
}

type OrderChangeTableReq struct {
	NewSerial int
	OldSerial int
}

type InsertItemWithModifiersReq struct {
	ItemsSerials    string
	HeadSerial      int
	OrderItemSerial int
}
