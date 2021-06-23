package tg

type (
	PushPayRequest struct {
		ReferenceID string
		Amount      string
		MSISDN      string
		Remarks     string
	}

	DisburseRequest struct {
	}

	//Reader implements only one method that takes the file
	// in CSV format read it and return an array of structs
	// of either PushPayRequest or DisburseRequest
	Reader interface {
		ReadFile(filepath string) (interface{}, error)
	}
)
