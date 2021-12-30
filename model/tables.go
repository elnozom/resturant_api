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
	Serial       int
	TableNo      int
	TableName    string
	Pause        bool
	State        string
	PrintTimes   int
	Status       string
	OpenDateTime string
}
