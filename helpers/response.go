package helpers

import (
	"chi-domain-go/constants"
	"chi-domain-go/models"
	"chi-domain-go/standard"
	"net/http"
)

func WriteJSONResponse(w http.ResponseWriter, response standard.Response) {
	w.WriteHeader(response.GetCode())
	w.Write(ToByteArray(response))
}

func NewOKSuccessResponse(data interface{}) *models.SuccessResponse {
	return &models.SuccessResponse{Code: http.StatusOK, Data: data}
}

func CreateInternalServerErrorResponse() *models.GeneralErrorResponse {
	return &models.GeneralErrorResponse{
		Code: http.StatusInternalServerError,
		Data: models.GeneralError{
			Type:    constants.GeneralErrorType,
			Message: constants.InternalServerErrorErrorMessage,
		},
	}
}

func CreateDataValidationErrorResponse(code int, info ...string) *models.DataValidationErrorResponse {
	return &models.DataValidationErrorResponse{
		Code: code,
		Data: models.DataValidationError{
			Type:    constants.DataValidationErrorType,
			Message: constants.DataValidationErrorMessage,
			Info:    info,
		},
	}
}
