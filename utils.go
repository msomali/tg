package tg

import (
	"fmt"
	"github.com/techcraftt/tigosdk/aw"
	"github.com/techcraftt/tigosdk/push"
	"os"
	"text/tabwriter"
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
	// initialize tabwriter
	w := new(tabwriter.Writer)

	// minwidth, tabwidth, padding, padchar, flags
	w.Init(os.Stdout, 8, 8, 0, '\t', 0)

	defer func(w *tabwriter.Writer) {
		err := w.Flush()
		if err != nil {
			fmt.Printf("error while closing tabwriter: %v\n",err)
		}
	}(w)

	_, _ = fmt.Fprintf(w, "\n %s\t", "RESPONSE")
	_, _ = fmt.Fprintf(w, "\n %s\t", "--------")

	switch reqType {
	case PushPay:
		response, ok := payload.(push.PayResponse)
		if !ok{
			fmt.Printf("unkown push pay response format cannot log")
		}

		_, _ = fmt.Fprintf(w, "\n %s\t%s\t", "ID:", response.ReferenceID)
		_, _ = fmt.Fprintf(w, "\n %s\t%s\t", "Code:", response.ResponseCode)
		_, _ = fmt.Fprintf(w, "\n %s\t%s\t", "Desc:", response.ResponseDescription)
		_, _ = fmt.Fprintf(w, "\n %s\t%t\t", "Status:", response.ResponseStatus)
		_, _ = fmt.Fprintf(w, "\n %s\t%s\t", "Message:", response.Message)

		return

	case Disburse:
		response, ok := payload.(aw.DisburseResponse)
		if !ok{
			fmt.Printf("unkown disbursement response format cannot log")
		}

		_, _ = fmt.Fprintf(w, "\n %s\t%s\t", "Type:", response.Type)
		_, _ = fmt.Fprintf(w, "\n %s\t%s\t", "Ref:", response.ReferenceID)
		_, _ = fmt.Fprintf(w, "\n %s\t%s\t", "ID:", response.TxnID)
		_, _ = fmt.Fprintf(w, "\n %s\t%s\t", "Status:", response.TxnStatus)
		_, _ = fmt.Fprintf(w, "\n %s\t%s\t", "Message:", response.Message)

		return
	}
}




