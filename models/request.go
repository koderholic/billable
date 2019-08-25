package models

type RequestModel struct {
	Employee  int    `json:"employeeId"`
	Rate      int    `json:"billableRate"`
	Project   string `json:"project"`
	Date      string `json:"date"`
	StartTime string `json:"startTime"`
	EndTime   string `json:"endTime"`
}
