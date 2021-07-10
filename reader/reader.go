/*
 * MIT License
 *
 * Copyright (c) 2021 TECHCRAFT TECHNOLOGIES CO LTD.
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

package reader

import (
	"bufio"
	"encoding/csv"
	"fmt"
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
		if path == "" {
			return
		}
		rd, ok := r.(*reader)
		if ok {
			rd.path = path
			return
		}

	}
}

func WithPushFileName(filename string) Option {
	return func(r Reader) {
		if filename == "" {
			return
		}
		rd, ok := r.(*reader)
		if ok {
			rd.pushFile = filename
			return
		}

	}
}

func WithDisburseFileName(filename string) Option {
	return func(r Reader) {
		if filename == "" {
			return
		}
		rd, ok := r.(*reader)
		if ok {
			rd.disburseFile = filename
			return
		}

	}
}

func Make(options ...Option) Reader {
	rd := &reader{
		path:         defaultPath,
		pushFile:     defaultPushFileName,
		disburseFile: defaultDisburseFileName,
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

		fmt.Printf("%v\n", amount)

		var remarks string

		if requestType == tg.PushPay {
			remarks = line[3]
		}
		requests = append(requests, tg.Request{
			ReferenceID: referenceID,
			//Amount:      int(amount),
			MSISDN:  msisdn,
			Remarks: remarks,
		})
	}

	return requests, nil
}
