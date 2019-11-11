package integration

import (
	"bytes"
	"chi-domain-go/helpers"
	"chi-domain-go/mock"
	"chi-domain-go/models"
	"chi-domain-go/router"
	"chi-domain-go/test/utils"
	"chi-domain-go/usecases/todos"
	"net/http"
	"net/http/httptest"
	"testing"
)

var emptyURLParams = []string{}
var emptyRequestBody = bytes.NewBuffer([]byte(""))

func TestTodosGetListSuccess(t *testing.T) {
	request, _ := http.NewRequest(http.MethodGet, "/bank/api/v1/todos", emptyRequestBody)
	recorder := httptest.NewRecorder()
	dbItem := models.TodoDBPayload{
		ID:          "todo:1234",
		UserID:      "user:1234",
		Title:       "Title Mock",
		Description: "Description Mock",
		IsArchive:   false,
		IsDone:      true,
	}

	uc := todos.NewTodosUseCase(todos.TodosRepositories{
		TodosMongoDBRepository: mock.MockMongoDBRepository{
			SelectResult: models.TodoDBPayloadList{dbItem},
		},
	})
	handler := router.UsecaseWrapper(uc.GetList, emptyURLParams)
	handler.ServeHTTP(recorder, request)

	response := helpers.NewOKSuccessResponse([]models.TodoResponse{dbItem.ToResponse()})
	expected := helpers.Stringify(response)
	result := recorder.Body.String()

	switch {
	case recorder.Code != http.StatusOK:
		t.Errorf("Expected no error but found %d", recorder.Code)
	case !utils.Equal(expected, result):
		t.Errorf("Expected result is %s but result shown %s", expected, result)
	}
}
