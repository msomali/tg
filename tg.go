package tg

const (
	PushPay RequestType = iota
	Disburse
)

type (
	RequestType int

	Request struct {
		ReferenceID string
		Amount      string
		MSISDN      string
		Remarks     string
	}

)
