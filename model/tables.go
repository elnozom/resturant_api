package model

type TableGroup struct {
	GroupTableNo   int
	GroupTableName string
	TableCount     int
}

type TablesListReq struct {
	GroupTableNo int
}

type TablesOpenOrderResp struct {
	IsOrderOpened bool
	Msg           string
}

type TablesOpenOrderReq struct {
	Imei    string
	Serial  int
	EmpCode int
}

type Table struct {
	Serial         int
	TableNo        int
	TableName      string
	Pause          bool
	State          string
	PrintTimes     int
	Status         string
	OpenDate       string
	OpenTime       string
	OrderNo        int
	BonNo          int
	DocNo          string
	HeadSerial     int
	Guests         int
	WaiterCode     int
	CustomerSerial int
	TotalCash      float64
}
