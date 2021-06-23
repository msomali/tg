package tg

const (
	PushPay RequestType = iota
	Disburse
)

type (

	Client struct {
		Push interface{}

	}
	RequestType int

	Request struct {
		ReferenceID string
		Amount      string
		MSISDN      string
		Remarks     string
	}

	Config struct {
		DisburseAccountName   string
		DisburseAccountMSISDN string
		DisburseBrandID       string
		DisbursePIN           string
		DisburseRequestURL    string

		PushUsername              string
		PushPassword              string
		PushPasswordGrantType     string
		PushApiBaseURL            string
		PushGetTokenURL           string
		PushBillerMSISDN          string
		PushBillerCode            string
		PushPayURL                string
		PushReverseTransactionURL string
		PushHealthCheckURL        string
	}

)
