package router

import (
	"chi-ledger-go/conf"
	"chi-ledger-go/constants"
	"chi-ledger-go/database"
	"chi-ledger-go/helpers"
	"chi-ledger-go/models"
	"chi-ledger-go/standard"
	"chi-ledger-go/usecases/ping"
	"chi-ledger-go/usecases/todos"
	"chi-ledger-go/validator"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

var emptyURLParams = []string{}

func UsecaseWrapper(businessLogicFunc standard.BusinessLogic, urlParams []string) http.HandlerFunc {
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
		configs := conf.Configs
		getTodosConfigPath := helpers.MakeGetConfigPathFunc("business.todos")
		todosUseCase := todos.NewTodosUseCase(todos.TodosRepositories{
			TodosMongoDBRepository: database.MongoDBRepository{
				Session:        dbSession.MongoDBSession,
				DatabaseName:   configs.GetString(getTodosConfigPath("database.name")),
				CollectionName: configs.GetString(getTodosConfigPath("database.collection")),
			},
		})

		router.With(
			validator.DomainLayerHeaderValidator(constants.UserIDHeaderName),
		).Get("/todos", UsecaseWrapper(todosUseCase.GetList, emptyURLParams))

		router.With(
			validator.DomainLayerHeaderValidator(constants.UserIDHeaderName),
			validator.AddTodoItemValidator,
		).Post("/todos", UsecaseWrapper(todosUseCase.Add, emptyURLParams))

		router.With(
			validator.DomainLayerHeaderValidator(constants.UserIDHeaderName),
			validator.UpdateTodoItemValidator,
		).Put(fmt.Sprintf("/todos/{%s}", constants.TodoIDParamName), UsecaseWrapper(todosUseCase.Update, []string{constants.TodoIDParamName}))

		router.With(
			validator.DomainLayerHeaderValidator(constants.UserIDHeaderName),
		).Delete(fmt.Sprintf("/todos/{%s}", constants.TodoIDParamName), UsecaseWrapper(todosUseCase.Delete, []string{constants.TodoIDParamName}))
	}
}

func mappingUseCaseToAPI(router *chi.Mux, dbSession *database.Session) *chi.Mux {
	pingUseCase := ping.NewPingUseCase()
	router.Get("/ping", UsecaseWrapper(pingUseCase.Ping, emptyURLParams))

	router.Route("/bank/api/v1", apiV1(dbSession))
	return router
}

func Initialise(dbSession *database.Session) *chi.Mux {
	router := chi.NewRouter()

	initialiseMiddleware(router)
	mappingUseCaseToAPI(router, dbSession)

	return router
}
