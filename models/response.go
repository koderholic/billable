package models

// Response hold respose returned on request
type ResponseModel struct {
	Ok      bool   `json:"ok"`
	Code int `json:"code"`
	Message string `json:"message"`
}

// SpendRsp hold respose returned for customer spending
type InvoiceModel struct {
	ResponseModel
	Data []Employee `json:"data"`
}

type Employee struct {
	Employee  int `json:"employeeId"`
	WorkHours int `json:"workHours"`
	Price     int `json:"unitPrice"`
	Cost      int `json:"cost"`
}