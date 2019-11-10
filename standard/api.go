package standard

import "chi-domain-go/models"

// type BusinessLogicInterface interface {
// 	GetResponse() Response
// }

type BusinessLogic func(req models.Request, varMap models.VarMap) Response

// func ()
