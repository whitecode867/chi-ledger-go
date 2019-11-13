package todos

import (
	"chi-ledger-go/constants"
	"chi-ledger-go/helpers"
	"chi-ledger-go/models"
	"chi-ledger-go/standard"
	"log"
	"net/http"

	"github.com/globalsign/mgo"
)

func checkTodoDBPayloadList(todoDBList models.TodoDBPayloadList) standard.Response {
	if todoDBList.IsEmpty() {
		return helpers.CreateDataValidationErrorResponse(http.StatusBadRequest, constants.DataDoesNotExistErrorInfo)
	}

	// size more than one means "bug" has occurred
	if size := todoDBList.Size(); size > 1 {
		log.Printf("result has more than 1 item (%d)\n", size)
	}

	return nil
}

func (uc todosImpl) GetList(req models.Request, varMap models.VarMap) standard.Response {
	selector := new(models.TodoSelector).SetUserID(req.Header[constants.UserIDHeaderName]).SetIsArchive(false)
	todoDBList := models.TodoDBPayloadList{}

	if err := uc.TodosMongoDBRepository.Select(selector, &todoDBList); err != nil && err != mgo.ErrNotFound {
		log.Println(err)
		return helpers.CreateInternalServerErrorResponse()
	}

	response := helpers.NewOKSuccessResponse(todoDBList.ToResponseList())
	return response
}

func (uc todosImpl) Add(req models.Request, varMap models.VarMap) standard.Response {
	body := new(models.TodoRequestBody)
	helpers.BytesToStruct(req.Body, body)

	payload := new(models.TodoDBPayload)
	helpers.MergeData(body, payload)

	payload.ID = helpers.GenerateUUID("todo")
	payload.UserID = req.Header[constants.UserIDHeaderName]

	nowMillsec := helpers.GetCurrentMillisecond()
	payload.CreatedAt = nowMillsec
	payload.UpdatedAt = nowMillsec

	if err := uc.TodosMongoDBRepository.Insert(payload); err != nil {
		log.Println(err)
		return helpers.CreateInternalServerErrorResponse()
	}

	response := helpers.NewOKSuccessResponse(payload.ToResponse())
	return response
}

func (uc todosImpl) Update(req models.Request, varMap models.VarMap) standard.Response {
	todoID, err := varMap.GetString(constants.TodoIDParamName)
	if err != nil {
		log.Println(err)
		return helpers.CreateInternalServerErrorResponse()
	}

	todoDBList := models.TodoDBPayloadList{}
	selector := new(models.TodoSelector).SetID(todoID).SetUserID(req.Header[constants.UserIDHeaderName])
	if err := uc.TodosMongoDBRepository.Select(selector, &todoDBList); err != nil {
		log.Println(err)
		return helpers.CreateInternalServerErrorResponse()
	}

	if errResponse := checkTodoDBPayloadList(todoDBList); errResponse != nil {
		return errResponse
	}

	body := new(models.TodoRequestBody)
	helpers.BytesToStruct(req.Body, body)

	oldData := todoDBList[0]
	updater := new(models.TodoDBPayload)
	helpers.MergeData(oldData, updater)
	helpers.MergeData(body, updater)

	updater.UpdatedAt = helpers.GetCurrentMillisecond()
	if err := uc.TodosMongoDBRepository.Update(selector, updater); err != nil {
		log.Println(err)
		return helpers.CreateInternalServerErrorResponse()
	}

	response := helpers.NewOKSuccessResponse(updater.ToResponse())
	return response
}

func (uc todosImpl) Delete(req models.Request, varMap models.VarMap) standard.Response {
	todoID, err := varMap.GetString(constants.TodoIDParamName)
	if err != nil {
		log.Println(err)
		return helpers.CreateInternalServerErrorResponse()
	}

	todoDBList := models.TodoDBPayloadList{}
	selector := new(models.TodoSelector).SetID(todoID).SetUserID(req.Header[constants.UserIDHeaderName])
	if err := uc.TodosMongoDBRepository.Select(selector, &todoDBList); err != nil {
		log.Println(err)
		return helpers.CreateInternalServerErrorResponse()
	}

	if errResponse := checkTodoDBPayloadList(todoDBList); errResponse != nil {
		return errResponse
	}

	oldData := todoDBList[0]
	updater := new(models.TodoDBPayload)
	helpers.MergeData(oldData, updater)

	updater.IsArchive = true
	if err := uc.TodosMongoDBRepository.Update(selector, updater); err != nil {
		log.Println(err)
		return helpers.CreateInternalServerErrorResponse()
	}

	response := helpers.NewOKSuccessResponse(updater.ToResponse())
	return response
}
