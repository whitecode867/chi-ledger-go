package router

import (
	"chi-domain-go/constants"
	"chi-domain-go/database"
	"chi-domain-go/helpers"
	"chi-domain-go/models"
	"chi-domain-go/standard"
	"chi-domain-go/usecases/ping"
	"chi-domain-go/usecases/todos"
	"chi-domain-go/validator"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

var emptyURLParams = []string{}

func usecaseWrapper(businessLogicFunc standard.BusinessLogic, urlParams []string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var response standard.Response
		defer func() {
			helpers.WriteJSONResponse(w, response)
			r.Body.Close()
		}()

		bodyBytes, err := ioutil.ReadAll(r.Body)
		switch {
		case err != nil:
			response = helpers.CreateInternalServerErrorResponse()
		default:
			varMap := models.VarMap{}
			varMap.ExtractURLParams(r, urlParams)
			header := models.HTTPHeader(r.Header).ExtractHTTPHeader()
			req := models.Request{Header: header, Body: bodyBytes}

			response = businessLogicFunc(req, varMap)
		}
	}
}

func initialiseMiddleware(r *chi.Mux) {
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.SetHeader(constants.ContentTypeHeaderName, constants.JSONContentType))
}

func apiV1(dbSession *database.Session) func(chi.Router) {
	return func(router chi.Router) {
		todosUseCase := todos.NewTodosUseCase(todos.TodosRepositories{
			TodosMongoDBRepository: database.MongoDBRepository{
				Session:        dbSession.MongoDBSession,
				DatabaseName:   "poc",
				CollectionName: "todos",
			},
		})

		router.With(
			validator.DomainLayerHeaderValidator(constants.UserIDHeaderName),
		).Get("/todos", usecaseWrapper(todosUseCase.GetList, emptyURLParams))

		router.With(
			validator.DomainLayerHeaderValidator(constants.UserIDHeaderName),
			validator.AddTodoItemValidator,
		).Post("/todos", usecaseWrapper(todosUseCase.Add, emptyURLParams))

		router.With(
			validator.DomainLayerHeaderValidator(constants.UserIDHeaderName),
			validator.UpdateTodoItemValidator,
		).Put(fmt.Sprintf("/todos/{%s}", constants.TodoIDParamName), usecaseWrapper(todosUseCase.Update, []string{constants.TodoIDParamName}))

		router.With(
			validator.DomainLayerHeaderValidator(constants.UserIDHeaderName),
		).Delete(fmt.Sprintf("/todos/{%s}", constants.TodoIDParamName), usecaseWrapper(todosUseCase.Delete, []string{constants.TodoIDParamName}))
	}
}

func mappingUseCaseToAPI(router *chi.Mux, dbSession *database.Session) *chi.Mux {
	pingUseCase := ping.NewPingUseCase()
	router.Get("/ping", usecaseWrapper(pingUseCase.Ping, emptyURLParams))

	router.Route("/bank/api/v1", apiV1(dbSession))
	return router
}

func Initialise(dbSession *database.Session) *chi.Mux {
	router := chi.NewRouter()

	initialiseMiddleware(router)
	mappingUseCaseToAPI(router, dbSession)

	return router
}