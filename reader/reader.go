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
	defaultPath             = ""
	defaultPushFileName     = "push.csv"
	defaultDisburseFileName = "disburse.csv"
)

var (
	_ Reader = (*reader)(nil)
)

type (
	reader struct {
		path         string
		pushFile     string
		disburseFile string
	}

	Option func(Reader)

	//Reader implements only one method that takes the file
	// in CSV format read it and return an array of structs
	// of either Request or Request
	Reader interface {
		ReadFile(filepath string, requestType tg.RequestType) ([]tg.Request, error)
	}
)


func WithDefaultPath(path string) Option {
	return func(r Reader) {
		rd, ok := r.(*reader)
		if ok {
			rd.path = path
			return
		}

	}
}

func WithDefaultFileName(filename string) Option {
	return func(r Reader) {
		rd, ok := r.(*reader)
		if ok {
			rd.path = filename
			return
		}

	}
}

func Make(options ...Option) Reader {
	rd := &reader{
		path:         defaultPath,
		pushFile:     defaultDisburseFileName,
		disburseFile: defaultPushFileName,
	}

	for _, option := range options {
		option(rd)
	}

	return rd
}

func (p *reader) ReadFile(filepath string, requestType tg.RequestType) ([]tg.Request, error) {

	if filepath == "" {
		//look for file in a working dir named push.csv
		filepath = defaultPushFileName
	}
	f, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	reader := csv.NewReader(bufio.NewReader(f))

	var requests []tg.Request
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}


		referenceID := line[0]
		amount := line[1]
		msisdn := line[2]

		var remarks string

		if requestType == tg.PushPay{
			remarks = line[3]
		}
		requests = append(requests, tg.Request{
			ReferenceID: referenceID,
			Amount:      amount,
			MSISDN:      msisdn,
			Remarks:     remarks,
		})
	}

	return requests, nil
}
