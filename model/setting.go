package model

type Setting struct {
	ComName        string
	Capital        string
	CashtraySerial int
}

type Option struct {
	StoreCode        int
	MultiPOS         bool
	TransSerial      int
	UseWaiter        bool
	AccountSerialint int
	SaleTax          float64
}

type InsertDeviceReq struct {
	Imei    string
	ComName string
}
