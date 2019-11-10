package models

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-chi/chi"
	"gopkg.in/gookit/color.v1"
)

type HTTPHeader http.Header

func (h HTTPHeader) ExtractHTTPHeader() map[string]string {
	headerMap := map[string]string{}
	if len(h) != 0 {
		for key, value := range h {
			headerMap[strings.ToLower(key)] = strings.Join(value, ",")
		}
	}

	return headerMap
}

type Request struct {
	Header map[string]string `json:"request_header"`
	Body   []byte            `json:"request_body"`
}

type VarMap map[string]interface{}

func (vm VarMap) createCannotGetValueError(key string, t string) error {
	colorErrText := color.FgLightRed.Render(fmt.Sprintf("\"%s\"", key))
	typeColorErrText := color.FgLightYellow.Render(t)
	return errors.New(fmt.Sprintf("Cannot get %s as type %s", colorErrText, typeColorErrText))
}

func (vm VarMap) ExtractURLParams(r *http.Request, urlParams []string) {
	if len(urlParams) == 0 {
		return
	}

	for _, urlParamName := range urlParams {
		value := chi.URLParam(r, urlParamName)
		vm[urlParamName] = value
	}
}

func (vm VarMap) GetString(key string) (string, error) {
	value, ok := vm[key].(string)
	if !ok {
		return "", vm.createCannotGetValueError(key, "string")
	}
	return value, nil
}

func (vm VarMap) GetInt(key string) (int, error) {
	value, ok := vm[key].(int)
	if !ok {
		return 0, vm.createCannotGetValueError(key, "int")
	}
	return value, nil
}

func (vm VarMap) GetFloat32(key string) (float32, error) {
	value, ok := vm[key].(float32)
	if !ok {
		return float32(0.0), vm.createCannotGetValueError(key, "float32")
	}
	return value, nil
}

func (vm VarMap) GetFloat64(key string) (float64, error) {
	value, ok := vm[key].(float64)
	if !ok {
		return 0.0, vm.createCannotGetValueError(key, "float64")
	}
	return value, nil
}
