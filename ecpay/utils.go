package ecpay

import "strings"

// MultipleItems join an array of item name with '#'.
func MultipleItems(items []string) string {
	return strings.Join(items, "#")
}
