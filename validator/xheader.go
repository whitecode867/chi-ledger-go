package validator

import (
	"chi-domain-go/constants"
	"chi-domain-go/models"
	"chi-domain-go/validator/rules"
	"fmt"
	"net/http"
)

// DomainLayerHeaderValidator - currently Support user_id ONLY
func DomainLayerHeaderValidator(additionalHeader ...string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			header := models.HTTPHeader(r.Header).ExtractHTTPHeader()

			hValidator := rules.NewValidator(header)
			hValidator.Add("x-app-id",
				rules.Required,
				rules.NotEmpty,
				rules.ShouldBeString,
			)

			hValidator.Add("x-correlation-id",
				rules.Required,
				rules.NotEmpty,
				rules.ShouldBeString,
			)

			hValidator.Add("x-activity",
				rules.Required,
				rules.NotEmpty,
				rules.ShouldBeString,
			)

			for _, aHeader := range additionalHeader {
				switch aHeader {
				case constants.UserIDHeaderName:
					hValidator.Add(constants.UserIDHeaderName,
						rules.Required,
						rules.NotEmpty,
						rules.ShouldBeString,
					)
				default:
					fmt.Println(aHeader, "is not supported. this will not be validated")
				}
			}

			fieldValidationHandler(fieldValidationHandlerParams{
				FailList:    hValidator.Validate(),
				HTTPHandler: next,
				HTTPRequest: r,
				HTTPWriter:  w,
			})
		}
		return http.HandlerFunc(fn)
	}
}

// func XHeaderValidator(next http.Handler) http.Handler {
// 	fn := func(w http.ResponseWriter, r *http.Request) {
// 		header := models.HTTPHeader(r.Header).ExtractHTTPHeader()

// 		hValidator := rules.NewValidator(header)
// 		hValidator.Add("x-app-id",
// 			rules.Required,
// 			rules.NotEmpty,
// 			rules.ShouldBeString,
// 		)

// 		hValidator.Add("x-correlation-id",
// 			rules.Required,
// 			rules.NotEmpty,
// 			rules.ShouldBeString,
// 		)

// 		hValidator.Add("x-activity",
// 			rules.Required,
// 			rules.NotEmpty,
// 			rules.ShouldBeString,
// 		)

// 		fieldValidationHandler(fieldValidationHandlerParams{
// 			FailList:    hValidator.Validate(),
// 			HTTPHandler: next,
// 			HTTPRequest: r,
// 			HTTPWriter:  w,
// 		})
// 	}
// 	return http.HandlerFunc(fn)
// }
