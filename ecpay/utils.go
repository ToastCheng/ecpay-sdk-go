package ecpay

import "strings"

func MultipleItems(items []string) string {
	return strings.Join(items, "#")
}
