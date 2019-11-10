package database

import (
	"chi-domain-go/helpers"
)

type MockMongoDBRepository struct {
	SelectOutput interface{}
	SelectError  error
	InsertError  error
	UpdateError  error
}

func (mock MockMongoDBRepository) Select(selector interface{}, output interface{}) error {
	helpers.MergeData(mock.SelectOutput, output)
	return mock.SelectError
}

func (mock MockMongoDBRepository) Insert(payload interface{}) error {
	return mock.InsertError
}

func (mock MockMongoDBRepository) Update(selector interface{}, updater interface{}) error {
	return mock.UpdateError
}
