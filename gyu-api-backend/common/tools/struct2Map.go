package tools

import "fmt"

func MapConvertAnyToString(source map[string]any) map[string]string {
	result := make(map[string]string, len(source))
	for k, v := range source {
		result[k] = fmt.Sprintf("%v", v)
	}
	return result
}

func MapConvertStringToAny(source map[string]string) map[string]any {
	result := make(map[string]any, len(source))
	for k, v := range source {
		result[k] = v
	}
	return result
}
