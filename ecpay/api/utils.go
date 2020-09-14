package api

import (
	"crypto/sha256"
	"fmt"
	"net/url"
	"strings"
)

func GetCheckMacValue(req url.Values) string {
	str := "HashKey=5294y06JbISpM5x9&"
	if req["ChoosePayment"][0] != "" {
		str += "ChoosePayment=" + req["ChoosePayment"][0] + "&"
	}
	if req["EncryptType"][0] != "" {
		str += "EncryptType=" + req["EncryptType"][0] + "&"
	}
	if req["ItemName"][0] != "" {
		str += "ItemName=" + req["ItemName"][0] + "&"
	}
	if req["MerchantID"][0] != "" {
		str += "MerchantID=" + req["MerchantID"][0] + "&"
	}
	if req["MerchantTradeDate"][0] != "" {
		str += "MerchantTradeDate=" + req["MerchantTradeDate"][0] + "&"
	}
	if req["MerchantTradeNo"][0] != "" {
		str += "MerchantTradeNo=" + req["MerchantTradeNo"][0] + "&"
	}
	if req["PaymentType"][0] != "" {
		str += "PaymentType=" + req["PaymentType"][0] + "&"
	}
	if req["ReturnURL"][0] != "" {
		str += "ReturnURL=" + req["ReturnURL"][0] + "&"
	}
	if req["TotalAmount"][0] != "" {
		str += "TotalAmount=" + req["TotalAmount"][0] + "&"
	}
	if req["TradeDesc"][0] != "" {
		str += "TradeDesc=" + req["TradeDesc"][0] + "&"
	}
	str += "HashIV=v77hoKGq4kWxNNIS"
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
