package validator

import (
	"chi-ledger-go/helpers"
	"chi-ledger-go/validator/rules"
	"net/http"
)

func AddTodoItemValidator(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		bodyBytes, _ := helpers.ExtractRequestBodyBytes(r)
		body := rules.NewValidator(bodyBytes)
		body.Add("title",
			rules.Required,
			rules.NotEmpty,
			rules.ShouldBeString,
		)

		body.Add("description",
			rules.ShouldBeString,
		)

		body.Add("is_done",
			rules.ShouldBeBoolean,
		)

		body.Add("is_archive",
			rules.ShouldBeBoolean,
		)

		fieldValidationHandler(fieldValidationHandlerParams{
			FailList:    body.Validate(),
			HTTPHandler: next,
			HTTPRequest: r,
			HTTPWriter:  w,
		})
	}
	return http.HandlerFunc(fn)
}

func UpdateTodoItemValidator(next http.Handler) http.Handler {
	return AddTodoItemValidator(next)
}
