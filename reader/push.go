package reader

import (
	"bufio"
	"encoding/csv"
	"github.com/techcraftt/tg"
	"io"
	"log"
	"os"
)

const (
	defaultPushPath = ""
	defaultDisbursePath = ""
	defaultPushFileName     = "push.csv"
	defaultDisburseFileName = "disburse.csv"
)

var (
	_ tg.Reader = (*PayFileReader)(nil)
	_ tg.Reader = (*DisburseFileReader)(nil)
)

type (
	PayFileReader struct {
		DefaultPath string
		DefaultFileName string
	}

	DisburseFileReader struct {
		DefaultPath string
		DefaultFileName string
	}

	OpErr struct {
		Err error
		Desc string
	}

	Option func(reader tg.Reader)

)

func (e *OpErr) Error() string{
	return ""
}

func DefaultPath(path string) Option {
	return func(reader tg.Reader) {
		push, pushOk := reader.(*PayFileReader)
		if pushOk{
			push.DefaultPath = path
			return
		}
		disburse, disburseOK := reader.(*DisburseFileReader)
		if disburseOK {
			disburse.DefaultPath = path
		}

	}
}

func DefaultFileName(filename string) Option {
	return func(reader tg.Reader) {
		push, pushOk := reader.(*PayFileReader)
		if pushOk{
			push.DefaultFileName = filename
			return
		}
		disburse, disburseOK := reader.(*DisburseFileReader)
		if disburseOK {
			disburse.DefaultFileName = filename
		}

	}
}

func ForPushOps(opts... Option) *PayFileReader{

	reader := &PayFileReader{
		DefaultPath:     defaultPushPath,
		DefaultFileName: defaultPushFileName,
	}

	for _, opt := range opts {
		opt(reader)
	}

	return reader
}

func ForDisburseOps(opts ... Option) *DisburseFileReader{
	reader := &DisburseFileReader{
		DefaultPath:     defaultDisbursePath,
		DefaultFileName: defaultDisburseFileName,
	}

	for _, opt := range opts {
		opt(reader)
	}

	return reader
}

func (d *DisburseFileReader) ReadFile(filepath string) (interface{}, error) {
	panic("implement me")
}

func (p *PayFileReader) ReadFile(filepath string) (interface{}, error) {

	if filepath == "" {
		//look for file in a working dir named push.csv
		filepath = defaultPushFileName
	}
	f, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	reader := csv.NewReader(bufio.NewReader(f))

	var requests []tg.PushPayRequest
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		requests = append(requests, tg.PushPayRequest{
			ReferenceID: line[0],
			Amount:      line[1],
			MSISDN:      line[2],
			Remarks:     line[3],
		})
	}

	return requests, nil
}
