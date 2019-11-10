package ping

import (
	"chi-domain-go/helpers"
	"chi-domain-go/models"
	"chi-domain-go/standard"
)

func (uc pingImpl) Ping(req models.Request, varMap models.VarMap) standard.Response {
	resp := helpers.NewOKSuccessResponse("OK")
	return resp
}
