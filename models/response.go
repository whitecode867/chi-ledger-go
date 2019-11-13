package models

import "chi-ledger-go/models/utils"

// ======== SUCCESS RESPONSE ========

type SuccessResponse struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

func (resp SuccessResponse) GetCode() int {
	return resp.Code
}

func (resp SuccessResponse) GetPayload() interface{} {
	return resp.Data
}

func (resp SuccessResponse) IsError() bool {
	return false
}

func (resp SuccessResponse) GetPayloadAs(data interface{}) interface{} {
	utils.MergeData(resp.Data, data)
	return data
}

// ======== GENERAL ERROR RESPONSE ========

type GeneralError struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}

type GeneralErrorResponse struct {
	Code int          `json:"code"`
	Data GeneralError `json:"error"`
}

func (err GeneralErrorResponse) GetCode() int {
	return err.Code
}

func (err GeneralErrorResponse) GetPayload() interface{} {
	return err.Data
}

func (err GeneralErrorResponse) IsError() bool {
	return true
}

func (err GeneralErrorResponse) Error() string {
	return utils.Stringify(err)
}

// ======== FIELD VALIDATION ERROR RESPONSE ========

type FieldValidationErrorItem struct {
	Field   string   `json:"field"`
	Reasons []string `json:"reasons"`
}

type FieldValidationErrorInfo struct {
	Errors []FieldValidationErrorItem `json:"errors"`
}

type FieldValidationError struct {
	Type    string                   `json:"type"`
	Message string                   `json:"message"`
	Info    FieldValidationErrorInfo `json:"info"`
}

type FieldValidationErrorResponse struct {
	Code int                  `json:"code"`
	Data FieldValidationError `json:"error"`
}

func (err FieldValidationErrorResponse) GetCode() int {
	return err.Code
}

func (err FieldValidationErrorResponse) GetPayload() interface{} {
	return err.Data
}

func (err FieldValidationErrorResponse) IsError() bool {
	return true
}

func (err FieldValidationErrorResponse) Error() string {
	return utils.Stringify(err)
}

// ======== DATA VALIDATION ERROR RESPONSE ========

type DataValidationError struct {
	Type    string   `json:"type"`
	Message string   `json:"message"`
	Info    []string `json:"info"`
}

type DataValidationErrorResponse struct {
	Code int                 `json:"code"`
	Data DataValidationError `json:"error"`
}

func (err DataValidationErrorResponse) GetCode() int {
	return err.Code
}

func (err DataValidationErrorResponse) GetPayload() interface{} {
	return err.Data
}

func (err DataValidationErrorResponse) IsError() bool {
	return true
}

func (err DataValidationErrorResponse) Error() string {
	return utils.Stringify(err)
}
