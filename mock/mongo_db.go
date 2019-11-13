package mock

import (
	"chi-ledger-go/helpers"
)

type MockMongoDBRepository struct {
	SelectResult interface{}
	SelectError  error
	InsertError  error
	UpdateError  error
}

func (mock MockMongoDBRepository) Select(selector interface{}, output interface{}) error {
	helpers.MergeData(mock.SelectResult, output)
	return mock.SelectError
}

func (mock MockMongoDBRepository) Insert(payload interface{}) error {
	return mock.InsertError
}

func (mock MockMongoDBRepository) Update(selector interface{}, updater interface{}) error {
	return mock.UpdateError
}
