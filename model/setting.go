package model

type Setting struct {
	ComName        string
	Capital        string
	CashtraySerial int
}

type InsertDeviceReq struct {
	Imei    string
	ComName string
}
