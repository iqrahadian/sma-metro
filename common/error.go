package common

import (
	"fmt"
)

type Error struct {
	Error error
	Code  int
}

type ErrorCode int

const (
	CardUnrecognized        ErrorCode = 100
	CardTypeUnrecognized              = 101
	CardInsufficientBalance           = 102

	TopupFailed = 200

	ChargeFailed = 300

	FaresUnknown = 400

	InternalUnknownError  = 500
	InternalParseTriptime = 501
)

var ErrorMap = map[ErrorCode]string{
	CardUnrecognized:        "Card unrecognized",
	CardTypeUnrecognized:    "Card type unrecognized",
	CardInsufficientBalance: "Insufficient card balance",

	TopupFailed: "Failed to process your topup",

	ChargeFailed: "Failed to charge your card",

	FaresUnknown: "Failed to calculate trip cost",

	InternalUnknownError:  "Currently we cannot process your request",
	InternalParseTriptime: "Failed to parse trip time",
}

func GetErrorMessage(code ErrorCode) string {

	if baseString, ok := ErrorMap[code]; ok {
		return fmt.Sprintf("%d | Sorry, %s,please call our customer service", code, baseString)
	} else {
		return GetErrorMessage(InternalUnknownError)
	}

}
