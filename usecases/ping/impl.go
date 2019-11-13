package ping

import (
	"chi-ledger-go/models"
	"chi-ledger-go/standard"
)

type Ping interface {
	Ping(req models.Request, varMap models.VarMap) standard.Response
}

type pingImpl struct {
}

func NewPingUseCase() Ping {
	return &pingImpl{}
}
