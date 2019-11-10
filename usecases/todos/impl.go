package todos

import (
	"chi-domain-go/models"
	"chi-domain-go/standard"
)

type Todos interface {
	GetList(req models.Request, varMap models.VarMap) standard.Response
	Add(req models.Request, varMap models.VarMap) standard.Response
	Update(req models.Request, varMap models.VarMap) standard.Response
	Delete(req models.Request, varMap models.VarMap) standard.Response
}

type todosImpl struct {
	TodosMongoDBRepository standard.DatabaseRepository
}

type TodosRepositories struct {
	TodosMongoDBRepository standard.DatabaseRepository
}

func NewTodosUseCase(repo TodosRepositories) Todos {
	return &todosImpl{
		TodosMongoDBRepository: repo.TodosMongoDBRepository,
	}
}
