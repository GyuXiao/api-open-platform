package sdk

import (
	"fmt"
	"github.com/duke-git/lancet/v2/cryptor"
)

func GenSign(params map[string]string, secretKey string) string {
	concatString := concatMapString(params)
	hms := cryptor.HmacSha256(concatString, secretKey)
	return hms
}

func concatMapString(paramsMap map[string]string) string {
	n := len(paramsMap)
	sortKeys := make([]string, n)
	for i := 0; i < n; i++ {
		sortKeys[i] = fmt.Sprintf("title%d", i)
	}
	concatString := ""
	for _, key := range sortKeys {
		concatString += paramsMap[key]
	}
	return concatString
}
