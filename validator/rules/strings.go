package rules

func ShouldBeString(key string, data map[string]interface{}) string {
	if msg := Required(key, data); len(msg) != 0 {
		return EmptyMessage
	}
	if _, ok := data[key].(string); !ok {
		return IsInvalidTypeOfStringMessage
	}
	return EmptyMessage
}
