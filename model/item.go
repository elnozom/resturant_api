package model

type Item struct {
	ItemSerial   int
	ItemPrice    float32
	ItemCode     int
	ItemName     string
	WithModifier bool
	Screen       int
}

type InsertItemReq struct {
	HeadSerial int
	ItemSerial int
	WithMod    bool
	IsMod      bool
	Qnt        int
}
