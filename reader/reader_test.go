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
	"github.com/techcraftt/tg"
	"reflect"
	"testing"
)

func TestMake(t *testing.T) {
	type args struct {
		options []Option
	}
	tests := []struct {
		name string
		args args
		want Reader
	}{
		{
			name: "reader with default parameters",
			args: args{options: nil},
			want: &reader{
				path:         "",
				pushFile:     "push.csv",
				disburseFile: "disburse.csv",
			},
		},
		{
			name: "reader without any default settings",
			args: args{
				options: []Option{
					WithDisburseFileName(""), WithPushFileName(""), WithDefaultPath(""),
				},
			},
			want: &reader{
				path:         "",
				pushFile:     "push.csv",
				disburseFile: "disburse.csv",
			},
		},
		{
			name: "reader with some names",
			args: args{
				options: []Option{
					WithDisburseFileName("theonedisbursefile.csv"), WithPushFileName("thepushmaestro.csv"),
				},
			},
			want: &reader{
				path:         "",
				pushFile:     "thepushmaestro.csv",
				disburseFile: "theonedisbursefile.csv",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Make(tt.args.options...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Make() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithDefaultFileName(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name string
		args args
		want Option
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithPushFileName(tt.args.filename); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithPushFileName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithDefaultPath(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want Option
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithDefaultPath(tt.args.path); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithDefaultPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_reader_ReadFile(t *testing.T) {
	type fields struct {
		path         string
		pushFile     string
		disburseFile string
	}
	type args struct {
		filepath    string
		requestType tg.RequestType
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []tg.Request
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &reader{
				path:         tt.fields.path,
				pushFile:     tt.fields.pushFile,
				disburseFile: tt.fields.disburseFile,
			}
			got, err := p.ReadFile(tt.args.filepath, tt.args.requestType)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadFile() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithDisburseFileName(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name string
		args args
		want Option
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithDisburseFileName(tt.args.filename); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithDisburseFileName() = %v, want %v", got, tt.want)
			}
		})
	}
}
