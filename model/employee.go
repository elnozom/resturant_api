package model

type EmpGetByCodeReq struct {
	EmpCode int
	BarCode int
}

type Emp struct {
	EmpName     string
	EmpCode     int
	EmpPassword string
	SecLevel    int
}
