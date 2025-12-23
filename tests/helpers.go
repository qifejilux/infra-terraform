package helpers

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func GetEnvOrDefault(key, default string) string {
	value := os.Getenv(key)
	if value == "" {
		return default
	}
	return value
}

func GetEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatal(fmt.Sprintf("Environment variable %s not set", key))
	}
	return value
}

func IsNil(i interface{}) bool {
	return i == nil
}

func IsNotNil(i interface{}) bool {
	return!IsNil(i)
}

func StringsJoin(sep string, strs []string) string {
	return strings.Join(strs, sep)
}

func StructSliceToMap[T any](slice []T) map[string]interface{} {
	result := map[string]interface{}{}
	for _, item := range slice {
		if v, ok := item.(map[string]interface{}); ok {
			for k, vv := range v {
				result[k] = vv
			}
		}
	}
	return result
}