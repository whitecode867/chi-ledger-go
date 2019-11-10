package ping

import (
	"chi-domain-go/models"
	"chi-domain-go/standard"
)

type Ping interface {
	Ping(req models.Request, varMap models.VarMap) standard.Response
}

type pingImpl struct {
}

func NewPingUseCase() Ping {
	return &pingImpl{}
}
