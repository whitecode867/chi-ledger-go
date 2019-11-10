package rules

func ShouldBeBoolean(key string, data map[string]interface{}) string {
	if msg := Required(key, data); len(msg) != 0 {
		return EmptyMessage
	}

	if _, ok := data[key].(bool); !ok {
		return IsInvalidTypeOfBooleanMessage
	}
	return EmptyMessage
}
