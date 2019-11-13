package helpers

import "fmt"

func MakeGetConfigPathFunc(prefix string) func(string) string {
	return func(key string) string {
		return fmt.Sprintf("%s.%s", prefix, key)
	}
}
