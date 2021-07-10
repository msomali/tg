package tg

import (
	"fmt"
	"github.com/techcraftt/tigosdk/aw"
	"github.com/techcraftt/tigosdk/push"
)

var (
	_ Printer = (*responsePrinter)(nil)
)

type (
	Printer interface {
		TextOut(reqType RequestType,payload interface{})
	}

	responsePrinter struct {}
)

func (r *responsePrinter) TextOut(reqType RequestType, payload interface{}) {
	switch reqType {
	case PushPay:
		response, ok := payload.(push.PayResponse)
		if !ok{
			fmt.Printf("unkown push pay response format cannot log")
		}

		fmt.Printf("PUSH PAY RESPONSE\n-----------------")
		fmt.Printf(response.ResponseCode)
		fmt.Printf(response.ResponseDescription)
		fmt.Printf("%t",response.ResponseStatus)
		fmt.Printf(response.ReferenceID)
		fmt.Printf(response.Message)
		return

	case Disburse:
		response, ok := payload.(aw.DisburseResponse)
		if !ok{
			fmt.Printf("unkown disbursement response format cannot log")
		}

		fmt.Printf("DISBURSEMENT RESPONSE\n-----------------")
		fmt.Printf(response.Type)
		fmt.Printf(response.TxnID)
		fmt.Printf("%s",response.TxnStatus)
		fmt.Printf(response.ReferenceID)
		fmt.Printf(response.Message)
		return
	}
}




