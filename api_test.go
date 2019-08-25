package main

import (
	"billable/api"
	"billable/config"
	"billable/models"
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gorilla/mux"
)

// GenerateInvoice function should accept only CSV files and should return an array of response struct
// For any wrong input it should  return an error
// for empty CSV, it should return an empty response
// test ReadCSV and readFile

var appp = startUp()

func startUp() *api.App {

	app := &api.App{}

	Config := config.Data{}
	Config.Init("")
	logPath := Config.LogPath

	app.Router = mux.NewRouter()
	app.LogPath = logPath
	app.RegisterRoutes()

	return app
}

func TestRegisterRoutesWorksCorrectly(t *testing.T) {

	pingRequest, _ := http.NewRequest("GET", "/api/ping", bytes.NewBuffer([]byte("")))
	invoiceRequest, _ := http.NewRequest("POST", "/api/invoice", bytes.NewBuffer([]byte("")))

	pingResponse := fireRequest(pingRequest)
	invoiceResponse := fireRequest(invoiceRequest)

	if pingResponse.Code == http.StatusNotFound {
		t.Errorf("Expected response code to not be %d. Got %d\n", http.StatusNotFound, pingResponse.Code)
	}
	if invoiceResponse.Code == http.StatusNotFound {
		t.Errorf("Expected response code to not be %d. Got %d\n", http.StatusNotFound, invoiceResponse.Code)
	}

}

func TestPingReturnsStatusOk(t *testing.T) {

	request, _ := http.NewRequest("GET", "/api/ping", bytes.NewBuffer([]byte("")))

	response := fireRequest(request)
	checkResponseCode(t, http.StatusOK, response.Code)
}

func TestGenerateInvoiceReturnsErrorMsgForEmptyInput(t *testing.T) {
	err, response := SendEmptyRequestBody()
	if err != nil {
		t.Errorf("Test terminated. Reason >> %s\n", err)
	}

	responseBody := models.ResponseModel{}
	json.Unmarshal(response.Body.Bytes(), &responseBody)

	if responseBody.Ok {
		t.Errorf("Expected response body Ok to be %t. Got %t\n", false, responseBody.Ok)
	} else if responseBody.Message == "" {
		t.Errorf("Expected response body Message field to not be empty. Got %s\n", responseBody.Message)
	}
}

func TestGenerateInvoiceReturnsErrorMsgForNonCSVFiles(t *testing.T) {

	err, response := SendRequestBody("test/wrongRequestFile.txt")
	if err != nil {
		t.Errorf("Test terminated. Reason >> %s\n", err)
	}

	responseBody := models.ResponseModel{}
	json.Unmarshal(response.Body.Bytes(), &responseBody)

	if responseBody.Ok {
		t.Errorf("Expected response body Ok to be %t. Got %t\n", false, responseBody.Ok)
	} else if responseBody.Message == "" {
		t.Errorf("Expected response body Message field to not be empty. Got %s\n", responseBody.Message)
	}
}
func TestGenerateInvoiceReturnsNoValueForInValidCSVInput(t *testing.T) {

	err, response := SendRequestBody("test/invalidRequest.csv")
	if err != nil {
		t.Errorf("Test terminated. Reason >> %s\n", err)
	}

	responseBody := models.InvoiceModel{}
	json.Unmarshal(response.Body.Bytes(), &responseBody)

	if responseBody.Code != http.StatusBadRequest {
		t.Errorf("Expected response code %d. Got %d\n", http.StatusBadRequest, responseBody.Code)
	}
	if responseBody.Ok {
		t.Errorf("Expected response body Ok to be %t. Got %t\n", false, responseBody.Ok)
	}
	if len(responseBody.Data) > 0 {
		t.Errorf("Expected length of data to be %d. Got %d\n", 0, len(responseBody.Data))
	}

}

func TestGenerateInvoiceReturnsStatusOkForValidInput(t *testing.T) {

	err, response := SendRequestBody("test/validRequest.csv")
	if err != nil {
		t.Errorf("Test terminated. Reason >> %s\n", err)
	}

	checkResponseCode(t, http.StatusOK, response.Code)
}

func TestGenerateInvoiceReturnsValueForValidCSVInput(t *testing.T) {

	err, response := SendRequestBody("test/validRequest.csv")
	if err != nil {
		t.Errorf("Test terminated. Reason >> %s\n", err)
	}

	responseBody := models.InvoiceModel{}
	json.Unmarshal(response.Body.Bytes(), &responseBody)

	if responseBody.Code != http.StatusOK {
		t.Errorf("Expected response code %d. Got %d\n", http.StatusOK, responseBody.Code)
	}
	if !responseBody.Ok {
		t.Errorf("Expected response body Ok to be %t. Got %t\n", true, responseBody.Ok)
	}
	if len(responseBody.Data) <= 0 {
		t.Errorf("Expected length of data returned to not be %d. Got %d\n", 0, len(responseBody.Data))
	}

}
func fireRequest(request *http.Request) *httptest.ResponseRecorder {
	// app := startUp()
	rr := httptest.NewRecorder()
	appp.Router.ServeHTTP(rr, request)

	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

func SendEmptyRequestBody() (error, *httptest.ResponseRecorder) {
	var requestBuffer bytes.Buffer
	multipartWriter := multipart.NewWriter(&requestBuffer)

	request, _ := http.NewRequest("POST", "/api/invoice", bytes.NewBuffer([]byte(requestBuffer.String())))
	request.Header.Set("Content-Type", multipartWriter.FormDataContentType())

	response := fireRequest(request)

	return nil, response

}

func SendRequestBody(fileInput string) (error, *httptest.ResponseRecorder) {

	file, err := os.Open(fileInput)
	if err != nil {
		return err, httptest.NewRecorder()
	}

	fileInfo, _ := file.Stat()

	var requestBuffer bytes.Buffer
	multipartWriter := multipart.NewWriter(&requestBuffer)

	fw, multipartErr := multipartWriter.CreateFormFile("csvReport", fileInfo.Name())
	if multipartErr != nil {
		return err, httptest.NewRecorder()
	}

	if _, err = io.Copy(fw, file); err != nil {
		return err, httptest.NewRecorder()
	}
	multipartWriter.Close()

	request, _ := http.NewRequest("POST", "/api/invoice", bytes.NewBuffer([]byte(requestBuffer.String())))
	request.Header.Set("Content-Type", multipartWriter.FormDataContentType())

	response := fireRequest(request)

	return nil, response

}
