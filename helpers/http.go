package helpers

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

func ExtractRequestBodyBytes(r *http.Request) ([]byte, error) {
	bodyBytes := []byte("")
	bodyBytes, err := ioutil.ReadAll(r.Body)
	r.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
	return bodyBytes, err
}
