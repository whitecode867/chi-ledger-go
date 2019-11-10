package rules

import (
	"chi-domain-go/helpers"
	"strings"
)

type Fail struct {
	Field   string   `json:"field"`
	Reasons []string `json:"reasons"`
}

func (fail Fail) ToMap() FailMap {
	return FailMap(map[string][]string{fail.Field: fail.Reasons})

}

type FailMap map[string][]string

func (fm FailMap) ToFail() Fail {
	fail := new(Fail)
	for field, reasons := range fm {
		fail.Field = field
		fail.Reasons = reasons
	}

	return *fail
}

type rulesFunc func(string, map[string]interface{}) string

type validateProcess struct {
	Data  map[string]interface{}
	Key   string
	Rules []rulesFunc
}

func (vp validateProcess) Validate() []Fail {
	reasons := []string{}
	for _, ruleFn := range vp.Rules {
		if reason := ruleFn(vp.Key, vp.Data); len(reason) != 0 {
			reasons = append(reasons, reason)
		}
	}

	failList := []Fail{}
	if len(reasons) != 0 {
		failList = append(failList, Fail{Field: vp.Key, Reasons: reasons})
	}
	return failList
}

func NewValidator(data interface{}) *validator {
	v := new(validator)
	switch dat := data.(type) {
	case []byte:
		helpers.BytesToStruct(dat, &v.data)
	case string:
		helpers.BytesToStruct([]byte(dat), &v.data)
	default:
		helpers.MergeData(data, &v.data)
	}
	return v
}

type validator struct {
	data          map[string]interface{}
	validateQueue []validateProcess
}

func (v *validator) Add(key string, validRules ...rulesFunc) {
	process := validateProcess{
		Key:   strings.ToLower(key),
		Data:  v.data,
		Rules: validRules,
	}
	v.validateQueue = append(v.validateQueue, process)
}

func (v validator) Validate() []Fail {
	failList := []Fail{}
	for _, process := range v.validateQueue {
		fList := process.Validate()
		if len(fList) > 0 {
			failList = append(failList, fList...)
		}
	}

	return failList
}
