package ecpay

import (
	"crypto/sha256"
	"fmt"
	"net/url"
	"sort"
	"strings"
)

// GetCheckMacValue gets the check mac.
func GetCheckMacValue(req url.Values, hashKey, hashIV string) string {
	keys := []string{}
	for k := range req {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	str := "HashKey=" + hashKey + "&"
	for _, k := range keys {
		if req[k][0] != "" {
			str += k + "=" + req[k][0] + "&"
		}
	}

	str += "HashIV=" + hashIV
	str = url.QueryEscape(str)
	str = strings.ReplaceAll(str, "%2A", "*")
	str = strings.ReplaceAll(str, "%28", "(")
	str = strings.ReplaceAll(str, "%29", ")")
	str = strings.ReplaceAll(str, "%21", "!")
	str = strings.ToLower(str)
	str = fmt.Sprintf("%x", sha256.Sum256([]byte(str)))
	str = strings.ToUpper(str)

	return str
}
