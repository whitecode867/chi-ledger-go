package rules

import "strings"

func Required(key string, data map[string]interface{}) string {
	for dKey := range data {
		if dKey == key {
			return EmptyMessage
		}
	}
	return IsRequireMessage
}

func NotEmpty(key string, data map[string]interface{}) string {
	if v, ok := data[key].(string); ok && len(strings.TrimSpace(v)) == 0 {
		return IsEmptyMessage
	}
	return EmptyMessage
}
