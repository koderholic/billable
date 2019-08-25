package api

import (
	"billable/models"
	"billable/utils"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	log "github.com/jeanphorn/log4go"
)

func (app App) Ping(w http.ResponseWriter, r *http.Request) {
	response := models.ResponseModel{
		Ok:      true,
		Code:    http.StatusOK,
		Message: "pong",
	}
	utils.SendResponse(w, http.StatusOK, response)
}

//GenerateInvoice gets request sent returns response gotten
func (app App) GenerateInvoice() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.LoadConfiguration(app.LogPath)

		// point to our verify finger print model
		response := &models.InvoiceModel{}

		file, handler, err := r.FormFile("csvReport")
		if err != nil {
			log.Error("Error retrieving file from request body, input file required >> %s", err)

			response.Ok = false
			response.Code = http.StatusBadRequest
			response.Message = "Error Retrieving File, input file required"

			log.Close()
			utils.SendResponse(w, http.StatusBadRequest, response)
			return
		}
		defer file.Close()
		log.Info("These are the request that came in FileName : %s, FileSize : %d", handler.Filename, handler.Size)

		if !strings.Contains(handler.Filename, ".csv") {
			log.Error("Invalid file extension for request >> FileName : %s. Only 'csv' files are allowed", handler.Filename)
			response.Ok = false
			response.Code = http.StatusBadRequest
			response.Message = "Only 'csv' files are allowed"
			utils.SendResponse(w, http.StatusBadRequest, response)
			return
		}

		err, csvContent := utils.ReadCSV(file)
		if err != nil {
			log.Error("could not read content of csv file. Reason >> %s", err)
			response.Ok = false
			response.Code = http.StatusBadRequest
			response.Message = "Error retrieving file content, ensure csv file contains proper formating and data properly structured."
			utils.SendResponse(w, http.StatusBadRequest, response)
			return
		}

		employeeInvoice := []models.Employee{}

		for _, employee := range csvContent[1:] {
			invoice := models.Employee{}
			startT, _ := time.Parse("2006-01-02T15:04", fmt.Sprintf("%sT%s", employee[3], employee[4]))
			endT, _ := time.Parse("2006-01-02T15:04", fmt.Sprintf("%sT%s", employee[3], employee[5]))
			workingHours := utils.GetDifferenceInSeconds(endT.Format("Mon, 02 Jan 2006 15:04:05 WAT"), startT.Format("Mon, 02 Jan 2006 15:04:05 WAT"))
			workingPrice, errPrice := strconv.Atoi(employee[1])
			if errPrice != nil {
				print(errPrice.Error())
			}
			EmployeeId, errEmployee := strconv.Atoi(employee[0])
			if errEmployee != nil {
				print(errEmployee.Error())
			}

			invoice.Employee = EmployeeId
			invoice.WorkHours = workingHours
			invoice.Price = workingPrice
			invoice.Cost = workingPrice * workingHours

			employeeInvoice = append(employeeInvoice, invoice)

		}

		response.Ok = true
		response.Code = http.StatusOK
		response.Message = "Employee invoice created successfully"
		response.Data = employeeInvoice

		log.Info("Response to request was successful %+v", response)
		log.Close()

		utils.SendResponse(w, http.StatusOK, response)
	}
}
