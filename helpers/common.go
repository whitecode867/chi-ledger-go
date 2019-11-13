package helpers

import (
	"chi-ledger-go/models/utils"
	"encoding/json"
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/google/uuid"
)

func MergeData(source interface{}, output interface{}) {
	utils.MergeData(source, output)
}

func Stringify(data interface{}) string {
	return utils.Stringify(data)
}

func ToMapStringInterface(data interface{}) map[string]interface{} {
	m := map[string]interface{}{}
	utils.MergeData(data, &m)
	return m
}

func ToBSONQuery(data interface{}) bson.M {
	return bson.M(ToMapStringInterface(data))
}

func ToByteArray(data interface{}) []byte {
	if bytes, err := json.Marshal(data); err == nil {
		return bytes
	}
	return []byte("")
}

func BytesToStruct(bytes []byte, data interface{}) {
	json.Unmarshal(bytes, &data)
}

func PrintStruct(data interface{}) string {
	result := ""
	bytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		result = err.Error()
	} else {
		result = string(bytes)
	}

	fmt.Printf("\n%s\n", result)
	return result
}

func GenerateUUID(prefix string) string {
	id := uuid.New().String()
	if len(prefix) == 0 {
		return id
	}
	return fmt.Sprintf("%s:%s", prefix, id)
}
