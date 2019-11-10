package validator

import (
	"chi-domain-go/constants"
	"chi-domain-go/helpers"
	"chi-domain-go/models"
	"chi-domain-go/validator/rules"
	"net/http"
)

type fieldValidationHandlerParams struct {
	HTTPHandler http.Handler
	HTTPWriter  http.ResponseWriter
	HTTPRequest *http.Request
	FailList    []rules.Fail
}

func fieldValidationHandler(p fieldValidationHandlerParams) {
	switch {
	case len(p.FailList) != 0:
		errors := []models.FieldValidationErrorItem{}
		for _, fail := range p.FailList {
			errors = append(errors, models.FieldValidationErrorItem{
				Field:   fail.Field,
				Reasons: fail.Reasons,
			})
		}

		response := models.FieldValidationErrorResponse{
			Code: http.StatusBadRequest,
			Data: models.FieldValidationError{
				Type:    constants.FieldValidationErrorType,
				Message: constants.FieldValidationErrorMessage,
				Info:    models.FieldValidationErrorInfo{Errors: errors},
			},
		}
		helpers.WriteJSONResponse(p.HTTPWriter, response)
	default:
		p.HTTPHandler.ServeHTTP(p.HTTPWriter, p.HTTPRequest)
	}
}
