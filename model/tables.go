package model

type TableGroup struct {
	GroupTableNo   int
	GroupTableName string
	TableCount     int
	UseMinimumBon  bool
	UseSellTax     bool
}

type TableGroupBo struct {
	GroupTableNo   int `json:"Id"`
	GroupTableName string
	StartNo        int
	TableCount     int
}
type TableGroupEditAddtReq struct {
	GroupTableNo   int
	GroupTableName string `json:"name"`
	StartNo        int    `json:"startTable"`
	TableCount     int    `json:"tableCount"`
}

type TablesListReq struct {
	GroupTableNo int
}

type TablesOpenOrderResp struct {
	IsOrderOpened bool
	Msg           string
}

type TablesOpenOrderReq struct {
	Imei       string
	Serial     int
	EmpCode    int
	HeadSerial int
}

type Table struct {
	Serial          int
	TableNo         int
	TableName       string
	Pause           bool
	State           string
	PrintTimes      int
	Status          string
	OpenDate        string
	OpenTime        string
	OrderNo         int
	BonNo           int
	DocNo           string
	HeadSerial      int
	Guests          int
	WaiterCode      int
	CustomerSerial  int
	Subtotal        float64
	DiscountPercent int
	DiscountValue   float64
	TotalCash       float64
	UseTax          bool
	MinimumBon      float64
	ComputerName    string
}
