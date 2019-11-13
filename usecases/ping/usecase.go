package ping

import (
	"chi-ledger-go/helpers"
	"chi-ledger-go/models"
	"chi-ledger-go/standard"
)

func (uc pingImpl) Ping(req models.Request, varMap models.VarMap) standard.Response {
	resp := helpers.NewOKSuccessResponse("OK")
	return resp
}
