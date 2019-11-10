package rules

func IsInt(key string, data map[string]interface{}) string {
	if failMessage := Required(key, data); len(failMessage) != 0 {
		return ""
	}

	if failMessage := NotEmpty(key, data); len(failMessage) != 0 {
		return ""
	}

	switch data[key].(type) {
	case int:
		return ""
	default:
		return "is invalid type of string"
	}
}
