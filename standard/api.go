package standard

import "chi-ledger-go/models"

type BusinessLogic func(req models.Request, varMap models.VarMap) Response
